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
    Keyword(String, State),
    Identifier(String, State),
    Curly(String, State),
    Round(String, State),
    // Square(String),
}

pub fn extract_state(token: Token) -> State {
    match token {
        Token::Number(_, state) => state,
        Token::Identifier(_, state) => state,
        Token::Curly(_, state) => state,
        Token::Round(_, state) => state,
        Token::Keyword(_, state) => state,
    }
}

pub enum Ast {
    Variable(String, String),
    Function(String, String, String, String),
}

pub fn get_token(v: String, t: u32, state: State) -> Token {
    match t {
        1 => Token::Number(v, state),
        2 => match v.as_str() {
            "enum" | "struct" | "namespace" => Token::Keyword(v, state),
            _ => Token::Identifier(v, state),
        },
        3 => Token::Round(v, state),
        4 => Token::Curly(v, state),
        _ => Token::Number(String::from(""), state),
    }
}

pub fn tokenize(code: String, root: &mut Transpiler) -> Vec<Token> {
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
            let mut t: u32 = token_type;
            root.state.column += 1;
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
                    '\t' | ' ' => tval = '\n',
                    '\n' | '\r' => {
                        tval = '\n';
                        root.state.line += 1;
                        root.state.column = 0;
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
                            state: root.state.clone(),
                        }));
                    }
                }
                if (t != token_type || tval == '\n') && token_type != 0 {
                    tokens.push(get_token(
                        token_value.clone(),
                        token_type,
                        token_state.clone(),
                    ));
                    token_type = 0;
                    token_value = String::new();
                }
                if tval != '\n' {
                    token_state = root.state.clone();
                    token_type = t;
                    token_value += tval.to_string().as_str().trim();
                }
            }
        }
    }
    if token_type != 0 {
        tokens.push(get_token(
            token_value.clone(),
            token_type,
            root.state.clone(),
        ));
    }
    tokens
}

pub fn parse(code: String, vars: &mut Variables, root: &mut Transpiler) -> Vec<Ast> {
    let mut ast: Vec<Ast> = Vec::new();
    let mut tokens = tokenize(code, root);
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
            x => {
                if !pcon.contains(&(root.state.line, root.state.column)) {
                    root.problems.push(ProblemCap::Error(Problem {
                        problem_msg: format!("Invalid placement"),
                        problem_type: ProblemType::SyntaxError,
                        state: extract_state(x[0].clone()),
                    }));
                    pcon.push((root.state.line, root.state.column));
                }
            }
        }
        tokens.drain(0..drain);
    }
    ast
}
