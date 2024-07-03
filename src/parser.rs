use std::usize;

use crate::{
    transpiler::{State, Transpiler},
    utils::{Problem, ProblemCap, ProblemType},
    variable::Variables,
};

#[derive(Clone, Debug)]
#[allow(dead_code)]
pub enum Token {
    Number(String, State),
    DKeyword(String, State),
    Identifier(String, State),
    Curly(String, State),
    Round(String, State),
    // Square(String),
}

pub fn extract_values(token: Token) -> (String, State) {
    match token {
        Token::Number(x, state) => (x, state),
        Token::Identifier(x, state) => (x, state),
        Token::Curly(x, state) => (x, state),
        Token::Round(x, state) => (x, state),
        Token::DKeyword(x, state) => (x, state),
    }
}

pub enum Ast {
    Variable(String, String),
    Function(String, String, String, String),
    Struct(String, String),
}

pub fn get_token(v: String, t: u32, state: State) -> Token {
    match t {
        1 => Token::Number(v, state),
        2 => match v.as_str() {
            "enum" | "struct" | "namespace" => Token::DKeyword(v, state),
            _ => Token::Identifier(v, state),
        },
        3 => Token::Round(v, state),
        4 => Token::Curly(v, state),
        _ => Token::Number(String::from(""), state),
    }
}

pub fn tokenize(code: String, state: &mut State, root: &mut Transpiler) -> Vec<Token> {
    let mut code = code;
    let mut tokens: Vec<Token> = Vec::new();
    let mut token_value: String = String::new();
    let mut token_type: u32 = 0;
    let mut stoken: (u8, u8) = (0, 0);
    let mut token_state = State { line: 0, column: 0 };
    while !code.is_empty() {
        if let Some(c) = code.chars().next() {
            code.remove(0);
            let mut tval = c;
            let mut t: u32 = 0;
            let mut ignore = false;
            state.column += 1;
            if stoken.1 > 0 {
                match c {
                    '(' => {
                        if stoken.0 == 1 {
                            stoken.1 += 1;
                            token_value += c.to_string().as_str();
                        }
                    }
                    ')' => {
                        if stoken.0 == 1 {
                            stoken.1 -= 1;
                            if stoken.1 == 0 {
                                stoken.0 = 0;
                                tokens.push(get_token(
                                    token_value.clone(),
                                    token_type,
                                    token_state.clone(),
                                ));
                                token_type = 0;
                                token_value = String::new();
                            } else {
                                token_value += c.to_string().as_str();
                            }
                        }
                    }
                    '{' => {
                        if stoken.0 == 1 {
                            stoken.1 += 1;
                            token_value += c.to_string().as_str();
                        }
                    }
                    '}' => {
                        if stoken.0 == 1 {
                            stoken.1 -= 1;
                            if stoken.1 == 0 {
                                stoken.0 = 0;
                                tokens.push(get_token(
                                    token_value.clone(),
                                    token_type,
                                    token_state.clone(),
                                ));
                                token_type = 0;
                                token_value = String::new();
                            } else {
                                token_value += c.to_string().as_str();
                            }
                        }
                    }
                    _ => token_value += c.to_string().as_str(),
                }
            } else {
                match c {
                    '0'..='9' => {
                        if token_type == 2 {
                            t = 2;
                        } else {
                            t = 1;
                        }
                    }
                    'a'..='z' | 'A'..='Z' => {
                        t = 2;
                    }
                    '\t' | ' ' => {
                        ignore = true;
                    }
                    '\n' | '\r' => {
                        ignore = true;
                        state.line += 1;
                        state.column = 0;
                    }
                    '(' => {
                        stoken.0 = 1;
                        stoken.1 = 1;
                        tval = ' ';
                        t = 3;
                    }
                    '{' => {
                        stoken.0 = 2;
                        stoken.1 = 1;
                        tval = ' ';
                        t = 4;
                    }
                    x => {
                        root.problems.push(ProblemCap::Error(Problem {
                            problem_msg: format!("Invalid character '{}'", x,),
                            problem_type: ProblemType::SyntaxError,
                            state: state.clone(),
                        }));
                        ignore = true;
                    }
                }
                if token_type != t && token_type != 0 {
                    tokens.push(get_token(
                        token_value.clone(),
                        token_type,
                        token_state.clone(),
                    ));
                    token_type = 0;
                    token_value = String::new();
                }
                if token_type != t {
                    token_state = state.clone();
                    token_type = t;
                    token_value += tval.to_string().as_str();
                } else if !ignore {
                    token_value += tval.to_string().as_str();
                }
            }
        }
    }
    if token_type != 0 {
        tokens.push(get_token(token_value.clone(), token_type, state.clone()));
    }
    tokens
}

pub fn parse(
    code: String,
    vars: &mut Variables,
    state: &mut State,
    root: &mut Transpiler,
) -> Vec<Ast> {
    let mut ast: Vec<Ast> = Vec::new();
    let mut tokens = tokenize(code, state, root);
    let mut pcon: Vec<(u32, u32)> = Vec::new();
    while tokens.len() > 0 {
        let mut drain: usize = 1;
        match tokens.as_slice() {
            [Token::Identifier(var_type, state0), Token::Identifier(var_name, _state1), Token::Round(round, _state2), Token::Curly(curly, _state3), ..] =>
            {
                ast.push(Ast::Function(
                    if var_type == "void" {
                        "()".to_string()
                    } else {
                        var_type.clone()
                    },
                    vars.new_func(var_name.clone(), state0.clone(), "".to_string()),
                    round.clone(),
                    curly.clone(),
                ));
                drain += 3;
            }
            [Token::Identifier(var_type, _state0), Token::Identifier(var_name, state1), ..] => {
                ast.push(Ast::Variable(
                    var_type.clone(),
                    vars.new_var(var_name.clone(), state1.clone(), "".to_string()),
                ));
                drain += 1;
            }
            [Token::DKeyword(keyword, _), Token::Identifier(name, state), Token::Curly(curly, _), ..] =>
            {
                match keyword.as_str() {
                    "struct" => ast.push(Ast::Struct(
                        vars.new_struct(name.clone(), state.clone(), String::from("")),
                        curly.clone(),
                    )),
                    _ => {}
                }
                drain += 2;
            }
            x => {
                let estate = extract_values(x[0].clone()).1;
                if !pcon.contains(&(estate.line, estate.column)) {
                    root.problems.push(ProblemCap::Error(Problem {
                        problem_msg: format!("Invalid placement"),
                        problem_type: ProblemType::SyntaxError,
                        state: estate,
                    }));
                    pcon.push((root.state.line, root.state.column));
                }
            }
        }
        tokens.drain(0..drain);
    }
    ast
}
