use std::usize;

use crate::{
    transpiler::{State, Transpiler},
    utils::{Problem, ProblemType},
    variable::Variables,
};

#[derive(Clone, Debug)]
pub enum Token {
    Number(String, State),
    Identifier(String, State),
    // Curly(String),
    // Round(String),
    // Square(String),
}
pub enum Ast {
    Variable(String, String),
}

pub fn get_token(v: String, t: u32, state: State) -> Token {
    match t {
        1 => Token::Number(v, state),
        2 => Token::Identifier(v, state),
        _ => Token::Number(String::from(""), state),
    }
}

pub fn tokenize(code: String, root: &mut Transpiler) -> Vec<Token> {
    let mut code = code;
    let mut tokens: Vec<Token> = Vec::new();
    let mut token_value: String = String::new();
    let mut token_type: u32 = 0;
    //let mut state = State { line:  }
    while !code.is_empty() {
        if let Some(c) = code.chars().next() {
            code.remove(0);
            let mut tval: char = c;
            let mut t: u32 = 0;
            root.state.column += 1;
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
                '\n' | '\r' => {
                    tval = '\n';
                    root.state.line += 1;
                    root.state.column = 0;
                }
                x => {
                    root.problems.push(Problem {
                        problem_msg: format!(
                            "Invalid character '{}' at {}:{}, LexingError",
                            x, root.state.line, root.state.column
                        ),
                        problem_type: ProblemType::SyntaxError,
                    });
                }
            }
            if t != token_type && token_type != 0 || tval == ' ' {
                tokens.push(get_token(
                    token_value.clone(),
                    token_type,
                    root.state.clone(),
                ));
                token_type = 0;
                token_value = String::new();
            }
            if tval != ' ' && tval != '\n' {
                token_type = t;
                token_value += tval.to_string().as_str();
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
            [Token::Identifier(var_type, state0), Token::Identifier(var_name, state1), ..] => {
                ast.push(Ast::Variable(var_type.clone(), var_name.clone()));
                drain += 1;
                vars.new_var(var_name.clone(), state1.clone(), "".to_string());
            }
            x => {
                if !pcon.contains(&(root.state.line, root.state.column)) {
                    root.problems.push(Problem {
                        problem_msg: format!(
                            "Invalid placement '{:?}' at {}:{}, ParserError",
                            x, root.state.line, root.state.column
                        ),
                        problem_type: ProblemType::SyntaxError,
                    });
                    pcon.push((root.state.line, root.state.column));
                }
            }
        }
        tokens.drain(0..drain);
    }
    ast
}
