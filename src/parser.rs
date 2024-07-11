use std::usize;

use once_cell::sync::Lazy;
use regex::Regex;

use crate::{
    transpiler::{State, Transpiler},
    utils::{Problem, ProblemCap, ProblemType, Variables},
};

pub struct Node {
    pub regex_pattern: Lazy<Regex>,
    pub token: u32,
}

const NODES: [Node; 3] = [
    Node {
        token: 1,
        regex_pattern: Lazy::new(|| Regex::new(r"^(0(x|X)\d+|\d+)").unwrap()),
    },
    Node {
        token: 2,
        regex_pattern: Lazy::new(|| Regex::new(r"^[.:_a-zA-Z][a-zA-Z0-9_.:]*").unwrap()),
    },
    Node {
        token: 5,
        regex_pattern: Lazy::new(|| Regex::new(r"^(\;|\!|\=)").unwrap()),
    },
];

#[derive(Clone, Debug)]
#[allow(dead_code)]
pub enum Token {
    Number(String, State),
    Keyword(String, State),
    SKeyword(String, State),
    DKeyword(String, State),
    Identifier(String, State),
    Curly(String, State),
    Round(String, State),
    Operator(String, State),
    String(String, State),
    // Square(String),
}

pub fn extract_values(token: Token) -> (u8, String, State) {
    match token {
        Token::Number(x, state) => (0, x, state),
        Token::Identifier(x, state) => (1, x, state),
        Token::Curly(x, state) => (2, x, state),
        Token::Round(x, state) => (3, x, state),
        Token::DKeyword(x, state) => (4, x, state),
        Token::SKeyword(x, state) => (5, x, state),
        Token::Keyword(x, state) => (6, x, state),
        Token::Operator(x, state) => (7, x, state),
        Token::String(x, state) => (8, x, state),
    }
}

pub fn extract_ast(ast: Ast) -> Vec<String> {
    let mut vstr: Vec<String> = Vec::new();
    match ast {
        Ast::Struct(a, _) => {
            vstr.push(a);
        }
        Ast::Single(a) => match a {
            Token::Identifier(b, _) => vstr.push(b),
            _ => {}
        },
        Ast::Variable(a, b) => {
            vstr.push(a);
            vstr.push(b);
        }
        Ast::Function(a, b, _, _) => {
            vstr.push(a);
            vstr.push(b);
        }
        _ => {}
    }
    vstr
}

#[derive(Clone, Debug)]
pub enum Ast {
    Variable(String, String),
    Function(String, String, String, String),
    Struct(String, String),
    Rust(String),
    Single(Token),
}

pub fn get_token(v: String, t: u32, state: State) -> Token {
    match t {
        1 => Token::Number(v, state),
        2 => match v.as_str() {
            "enum" | "struct" | "namespace" => Token::DKeyword(v, state),
            "rust" => Token::SKeyword(v, state),
            _ => Token::Identifier(v, state),
        },
        3 => Token::Round(v, state),
        4 => Token::Curly(v, state),
        5 => Token::Operator(v, state),
        6 => Token::String(v, state),
        _ => Token::Number(String::from(""), state),
    }
}

pub fn tokenzie_regex(code: String, state: &mut State) -> (u32, String) {
    let mut token = (0, String::new());
    if !code.is_empty() {
        for node in &NODES {
            if let Some(caps) = node.regex_pattern.captures(&code) {
                token.0 = node.token;
                token.1 = caps[0].to_string();
                state.column += token.1.len() as u32;
                if token.1.contains("\n") {
                    state.column = 0;
                    state.line += token.1.chars().filter(|c| *c == '\n').count() as u32;
                }
            }
        }
    }
    token
}

pub fn tokenize(code: String, state: &mut State, root: &mut Transpiler) -> Vec<Token> {
    let mut code = code;
    code += " ";
    let mut tokens: Vec<Token> = Vec::new();
    let mut token_value: String = String::new();
    let mut token_type: u32 = 0;
    let mut stoken: (u8, u8) = (0, 0);
    let mut token_state = State {
        line: 0,
        column: 0,
        file: None,
    };
    while !code.is_empty() {
        if let Some(c) = code.chars().next() {
            code.remove(0);
            let tc = c.to_string();
            let mut tval = tc.as_str();
            let mut t: u32 = 0;
            let mut ignore = false;
            state.column += 1;
            if root.inject.inject
                && root.inject.state.line == state.line
                && root.inject.state.column == state.column
            {
                token_value += "list_vx";
            }
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
                        if stoken.0 == 2 {
                            stoken.1 += 1;
                            token_value += c.to_string().as_str();
                        }
                    }
                    '}' => {
                        if stoken.0 == 2 {
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
                    '"' => {
                        token_value += c.to_string().as_str();
                        if stoken.0 == 3 {
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
                            }
                        }
                    }
                    '\\' => {
                        token_value += c.to_string().as_str();
                        token_value += code
                            .chars()
                            .next()
                            .expect("Error expected a char for \\")
                            .to_string()
                            .as_str();
                    }
                    _ => token_value += c.to_string().as_str(),
                }
            } else {
                match c {
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
                        tval = " ";
                        t = 3;
                    }
                    '{' => {
                        stoken.0 = 2;
                        stoken.1 = 1;
                        tval = " ";
                        t = 4;
                    }
                    '"' => {
                        stoken.0 = 3;
                        stoken.1 = 1;
                        t = 6;
                        ignore = true;
                    }
                    x => {
                        ignore = true;
                        let vals = tokenzie_regex(format!("{}{}", c, code.clone()), state);
                        if vals.0 != 0 {
                            tokens.push(get_token(vals.1.clone(), vals.0, state.clone()));
                            code.drain(0..vals.1.len() - 1);
                        } else {
                            root.problems.push(ProblemCap::Error(Problem {
                                problem_msg: format!("Invalid syntax '{}'", x,),
                                problem_type: ProblemType::SyntaxError,
                                state: state.clone(),
                            }));
                        }
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
    while tokens.len() > 0 {
        let mut drain: usize = 1;
        let vals = extract_values(tokens[0].clone());
        if vals.1.contains("list_vx") {
            ast.push(Ast::Single(Token::Identifier(vals.1, vals.2)));
            tokens.drain(0..1);
            continue;
        }
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
            [Token::SKeyword(keyword, _), Token::Curly(curly, _), ..] => {
                match keyword.as_str() {
                    "rust" => ast.push(Ast::Rust(curly.clone())),
                    _ => {}
                }
                drain += 1;
            }
            x => {
                ast.push(Ast::Single(x[0].clone()));
                //let estate = extract_values(x[0].clone()).1;
                //if !pcon.contains(&(estate.line, estate.column)) {
                //    root.problems.push(ProblemCap::Error(Problem {
                //        problem_msg: format!("Invalid placement"),
                //        problem_type: ProblemType::SyntaxError,
                //        state: estate,
                //    }));
                //    pcon.push((root.state.line, root.state.column));
                //}
            }
        }
        tokens.drain(0..drain);
    }
    ast
}
