use  crate::lexer::{Token, TokenType};
use std::fmt;
use regex::Regex;
use once_cell::sync::Lazy;

#[derive(Clone, Debug, PartialEq, Eq)]
pub enum AstType {
    FunctionDeceleration,
    VoidFunctionDeceleration,
    VariableDeceleration,
    MutVariableDeceleration,
    Include,
    IncludeLocal,
    CodeBlock,
    Other
}

pub struct Ast {
    pub tokens: Vec<Token>,
    pub ast_type: AstType
}

impl fmt::Display for Ast {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "\x1b[36m{:?}:\x1b[0m [\n", self.ast_type)?;
        for (i, token) in self.tokens.iter().enumerate() {
            if i < self.tokens.len() - 1 {
                write!(f, "    {},\n", token)?;
            } else {
                write!(f, "    {}\n", token)?;
            }
        }
        write!(f, "]")
    }
}

impl fmt::Debug for Ast {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}: [\n", self.ast_type)?;
        for (i, token) in self.tokens.iter().enumerate() {
            if i < self.tokens.len() - 1 {
                write!(f, "    {:?},\n", token)?;
            } else {
                write!(f, "    {:?}\n", token)?;
            }
        }
        write!(f, "]")
    }
}

pub struct Parser {
    pub tokens: Vec<Token>,
    pub index: u32,
    pub include_regex: Lazy<Regex>,
    pub include_regex_local: Lazy<Regex>
}

impl Parser {
    pub fn new(tokens: Vec<Token>) -> Parser {
        Parser {
            tokens: tokens,
            index: 0,
            include_regex: Lazy::new(|| Regex::new(r"^(#include *)<(.*?)>").unwrap()),
            include_regex_local: Lazy::new(|| Regex::new(r#"^(#include *)"(.*?)""#).unwrap())
        }
    }
    pub fn next(&mut self) -> Ast {
        let mut ast_res: Ast = Ast {tokens: vec![], ast_type: AstType::Other};
        let index = self.index as usize;
        if index == self.tokens.len() {panic!("Reached the end of tokens")}
        let token = &self.tokens[index];
        match token.token_type {
            TokenType::Identifier => {
                ast_res.tokens.push(self.tokens[index].clone());
                if self.tokens.len()-index > 3 && self.tokens[index+1].token_type==TokenType::Identifier && self.tokens[index+2].token_type==TokenType::Round && self.tokens[index+3].token_type==TokenType::Curly {
                    ast_res.tokens.push(self.tokens[index+1].clone());
                    ast_res.tokens.push(self.tokens[index+2].clone());
                    ast_res.tokens.push(self.tokens[index+3].clone());
                    if token.value == "void" {
                        ast_res.ast_type = AstType::VoidFunctionDeceleration;
                    } else {
                        ast_res.ast_type = AstType::FunctionDeceleration;
                    }
                    self.index += 3;
                } else if self.tokens.len()-index > 1 {
                    if token.value == "cb" && self.tokens[index+1].token_type == TokenType::Curly {
                        ast_res.tokens.push(self.tokens[index+1].clone());
                        ast_res.ast_type = AstType::CodeBlock;
                        self.index += 1;
                    } else if self.tokens[index+1].token_type==TokenType::Identifier {
                        ast_res.tokens.push(self.tokens[index+1].clone());
                        ast_res.ast_type = AstType::VariableDeceleration;
                        self.index += 1;
                    } else if self.tokens[index+1].value=="mut" && self.tokens[index+2].token_type==TokenType::Identifier {
                        ast_res.tokens.push(self.tokens[index+2].clone());
                        ast_res.ast_type = AstType::MutVariableDeceleration;
                        self.index += 2;
                    }
                }
            }
            TokenType::Include => {
                if let Some(caps) = self.include_regex.captures(&token.value) {
                    ast_res.tokens.push(Token {token_type: TokenType::String, value: caps[2].to_owned().to_string(), line: 0, column: 0});
                    ast_res.ast_type = AstType::Include;
                } else if let Some(caps) = self.include_regex_local.captures(&token.value) {
                    ast_res.tokens.push(Token {token_type: TokenType::String, value: caps[2].to_owned().to_string(), line: 0, column: 0});
                    ast_res.ast_type = AstType::IncludeLocal;
                } else {
                    ast_res.tokens.push(token.clone());
                    ast_res.ast_type = AstType::Include;
                }
            }
            _ => {
                ast_res.tokens.push(token.clone());
            }
        }
        self.index += 1;
        ast_res
    }
}