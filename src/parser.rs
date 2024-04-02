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

#[derive(Debug, PartialEq, Clone, Copy)]
pub enum AstDef {
    Normal(TokenType),
    Repeated(TokenType),
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
            ast_match: vec![AstDef::Normal(TokenType::Identifier), AstDef::Normal(TokenType::Identifier), AstDef::Normal(TokenType::Round), AstDef::Normal(TokenType::Curly)]
        },
        PNode {
            ast_type: AstTypes::FunctionDecleration,
            ast_match: vec![AstDef::Normal(TokenType::Keyword)]
        },
        PNode {
            ast_type: AstTypes::Other,
            ast_match: vec![AstDef::Else]
        },
    ];
    let tokens_len = tokens.len();
    let mut ast: Vec<Ast> = Vec::new(); 
    let mut current = 0;
    let mut i = 0;
    while tokens.len()>0 {
        for ast_node in &ast_def {
            if tokens_len < ast_node.ast_match.len() {
                continue;
            }
            let mut ast_types: Vec<AstTypes> = vec![];
            let mut is_match = false;
            let mut matched: Vec<Token> = Vec::new();
            for ast_match in &ast_node.ast_match {
                match ast_match {
                    AstDef::Normal(token_type) => {
                        if &tokens[i].token_type == token_type {
                            matched.push(tokens[i].clone());
                            tokens.drain(0..1);
                            is_match = true;
                        } else {
                            is_match = false;
                            break;
                        }
                    }
                    AstDef::Else => {
                        matched.push(tokens[i].clone());
                        tokens.drain(0..1);
                        is_match = true;
                    }
                    _ => {}
                }
            }
            if is_match {
                i = 0;
                ast.push(Ast {values: matched, type_: ast_node.ast_type.clone()});
                matched = Vec::new();
                break;
            }
        }
    }
    Box::new(ast)
}
