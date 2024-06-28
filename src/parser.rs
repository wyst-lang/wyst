use crate::{transpiler::State, variable::Variables};

#[derive(Clone, Debug)]
pub enum Token {
    Number(String),
    Identifier(String),
    // Curly(String),
    // Round(String),
    // Square(String),
}
pub enum Ast {
    Variable(String, String),
}

pub fn get_token(v: String, t: u32) -> Token {
    match t {
        1 => Token::Number(v),
        2 => Token::Identifier(v),
        _ => Token::Number(String::from("")),
    }
}

pub fn tokenize(code: String) -> Vec<Token> {
    let mut code = code;
    let mut tokens: Vec<Token> = Vec::new();
    let mut token_value: String = String::new();
    let mut token_type: u32 = 0;
    while !code.is_empty() {
        if let Some(c) = code.chars().next() {
            code.remove(0);
            let mut tval: char = c;
            let mut t: u32 = 0;
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
                '\t' | ' ' => tval = ' ',
                _ => {}
            }
            if t != token_type && token_type != 0 || tval == ' ' {
                tokens.push(get_token(token_value.clone(), token_type));
                token_type = 0;
                token_value = String::new();
            }
            if tval != ' ' {
                token_type = t;
                token_value += tval.to_string().as_str();
            }
        }
    }
    if token_type != 0 {
        tokens.push(get_token(token_value.clone(), token_type));
    }
    tokens
}

pub fn parse(code: String, vars: &mut Variables) -> Vec<Ast> {
    let mut ast: Vec<Ast> = Vec::new();
    let mut tokens = tokenize(code);
    while tokens.len() > 0 {
        let mut drain: usize = 0;
        match tokens.as_slice() {
            [Token::Identifier(var_type), Token::Identifier(var_name), ..] => {
                ast.push(Ast::Variable(var_type.clone(), var_name.clone()));
                drain += 2;
                // vars.new_var(var_name, State {}, desc)
            }
            _ => {}
        }
        tokens.drain(0..drain);
    }
    ast
}
