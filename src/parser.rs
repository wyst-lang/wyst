// use std::collections::HashMap;
use  crate::lexer::{Token, TokenType};
use std::fmt;

#[derive(Debug, Clone)]
pub enum AstTypes {
    FunctionDecleration,
    Other
}

pub struct Ast {
    pub type_: AstTypes,
    pub values: Vec<Token>
}

impl<'a> fmt::Debug for Ast {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "type={:?} values={:?}", self.type_, self.values)
    }
}

#[derive(Debug, PartialEq, Clone)]
pub enum AstDef {
    Normal(TokenType),
    NormalValue(TokenType, Vec<String>),
    Repeated(Vec<TokenType>),
    Optional(TokenType),
    Else
}

pub struct PNode {
    ast_type: AstTypes,
    ast_match: Vec<AstDef>
}

pub fn parse(tokens: &mut Vec<Token>) -> Box<Vec<Ast>> {
    let ast_def: Vec<PNode> = vec![
        PNode {
            ast_type: AstTypes::FunctionDecleration,
            ast_match: vec![AstDef::NormalValue(TokenType::Identifier, vec!["if".to_string()]), AstDef::Repeated(vec![TokenType::Operator, TokenType::Identifier, ])]
        },
        PNode {
            ast_type: AstTypes::Other,
            ast_match: vec![AstDef::Else]
        },
    ];
    let tokens_len = tokens.len();
    let mut ast: Vec<Ast> = Vec::new(); 
    // let mut i = 0;
    while tokens.len()>0 {
        for ast_node in &ast_def {
            if tokens_len < ast_node.ast_match.len() {
                continue;
            }
            let mut is_match = false;
            let mut matched: Vec<Token> = Vec::new();
            for ast_match in &ast_node.ast_match {
                match ast_match {
                    AstDef::Normal(token_type) => {
                        if &tokens[0].token_type == token_type {
                            matched.push(tokens[0].clone());
                            tokens.drain(0..1);
                            is_match = true;
                        } else {
                            is_match = false;
                            break;
                        }
                    }
                    AstDef::NormalValue(token_type, token_value) => {
                        if &tokens[0].token_type == token_type && &tokens[0].token_values == token_value {
                            matched.push(tokens[0].clone());
                            tokens.drain(0..1);
                            is_match = true;
                        } else {
                            is_match = false;
                            break;
                        }
                    }
                    AstDef::Repeated(nodes) => {
                        loop {
                            let mut node_match = false;
                            for &node_ in nodes {
                                if node_ == tokens[0].token_type {
                                    matched.push(tokens[0].clone());
                                    node_match = true;
                                    tokens.drain(0..1);
                                }
                            }
                            if node_match {
                                is_match = true;
                            } else {
                                break;
                            }
                        }
                    }
                    AstDef::Optional(token_type) => {
                        if &tokens[0].token_type == token_type {
                            matched.push(tokens[0].clone());
                            tokens.drain(0..1);
                            is_match = true;
                            break;
                        }
                    }
                    AstDef::Else => {
                        matched.push(tokens[0].clone());
                        tokens.drain(0..1);
                        is_match = true;
                    }
                    _ => {}
                }
            }
            if is_match {
                ast.push(Ast {values: matched, type_: ast_node.ast_type.clone()});
                matched = Vec::new();
                break;
            }
        }
    }
    Box::new(ast)
}
