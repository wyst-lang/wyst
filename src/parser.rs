use  crate::lexer::{Token, TokenType};
use std::fmt;

#[derive(Clone, Debug)]
pub enum AstType {
    FunctionDecleration,
    VariableDecleration,
    Other
}

pub struct Ast {
    pub tokens: Vec<Token>,
    pub ast_type: AstType
}

impl fmt::Display for Ast {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}: [\n", self.ast_type)?;
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



pub struct Parser {
    pub tokens: Vec<Token>,
    pub index: u32
}


impl Parser {
    pub fn new(tokens: Vec<Token>) -> Parser {
        Parser {tokens: tokens, index: 0}
    }
    pub fn next(&mut self) -> Ast {
        let mut ast_res: Ast = Ast {tokens: vec![], ast_type: AstType::Other};
        let mut index = self.index as usize;
        if index == self.tokens.len() {panic!("Reached the end of tokens")}
        let token = &self.tokens[index];
        match token.token_type {
            TokenType::Identifier => {
                if self.tokens.len()-index > 2 && self.tokens[index+1].token_type==TokenType::Identifier && self.tokens[index+2].token_type==TokenType::Round && self.tokens[index+3].token_type==TokenType::Curly {
                    ast_res.tokens.push(self.tokens[index].clone());
                    ast_res.tokens.push(self.tokens[index+1].clone());
                    ast_res.tokens.push(self.tokens[index+2].clone());
                    ast_res.tokens.push(self.tokens[index+3].clone());
                    ast_res.ast_type = AstType::FunctionDecleration;
                    self.index += 3;
                } else {
                    loop {
                        if index+1 >= self.tokens.len() {break;}
                        index = self.index as usize; // Update the index
                        let ntk = &self.tokens[index]; // Stands for next token
                        if ntk.value=="," {
                            ast_res.tokens.push(ntk.clone());
                            self.index += 1;
                        } else if ntk.token_type==TokenType::Identifier {
                            ast_res.tokens.push(ntk.clone());
                            ast_res.ast_type = AstType::VariableDecleration;
                            self.index += 1;
                        } else if ntk.token_type==TokenType::Keyword {
                            ast_res.tokens.push(ntk.clone());
                            ast_res.ast_type = AstType::VariableDecleration;
                            self.index += 1;
                        } else if ntk.token_type==TokenType::Angle {
                            ast_res.tokens.push(ntk.clone());
                            self.index += 1;
                        } else {
                            self.index -= 1;
                            break;
                        } 
                    }
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