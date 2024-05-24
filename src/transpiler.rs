use std::fs;
use crate::lexer::{lex, LexerState, TokenType};
use crate::parser::{Ast, AstType, Parser};

#[derive(Debug, PartialEq, Clone)]
pub struct Options {
    auto_mut: bool,
    auto_macro: bool,
    macros: Vec<String>,
    modnum: u32
}

impl Default for Options {
    fn default() -> Options {
        Options { auto_mut: true, auto_macro: true, macros: vec![String::from("println")], modnum: 0 }
    }
}

fn clean_incl(input: &str) -> String {
    input
        .chars()
        .map(|c| if c.is_alphanumeric() || c == '_' { c } else { '_' })
        .collect()
}

pub fn transpile_mod(ast: Ast, options: &mut Options) -> String {
    let modfile = ast.tokens[0].value.as_str();
    let modname = format!("{}_{}", clean_incl(modfile.split(".").collect::<Vec<_>>()[0]), options.clone().modnum);
    let file_content = fs::read_to_string("lib/".to_string()+modfile)
        .expect("Error reading file");
    let transpiled_code = transpile(file_content, 0, LexerState { line: 1, column: 0 }, &mut Options::default());
    fs::write(("wyst_tmp/".to_string()+modname.as_str())+".rs",
        transpiled_code)
        .expect("Error writing file");
    modname
}

pub fn transpile(input: String, indent: u32, state: LexerState, options: &mut Options) -> String {
    let mut result = String::new();
    
    if indent == 0 {
        // result += "type int = i32;\n";
    } else {
        result += " ".repeat((indent as usize) * 2).as_str();
    }
    
    let lexer_out = lex(input.as_str(), false, state);
    
    match lexer_out {
        Ok(tokens) => {
            let mut full_ast = Parser::new(tokens.clone());
            let mut last_ast = Ast {ast_type: AstType::Other, tokens: vec![]};
            while full_ast.tokens.len() > full_ast.index as usize {
                let ast = full_ast.next();
                println!("{ast}");
                if last_ast.tokens.len() > 0 {
                    let mut fl = 0;
                    for t in &last_ast.tokens { fl+=t.value.len() }
                    // if ast.tokens[ast.tokens.len()-1].line > last_ast.tokens[last_ast.tokens.len()-1].line {
                    //     // result += ("\n".to_string() + " ".repeat(((indent+1) as usize)*2).as_str()).repeat(ast.tokens[ast.tokens.len()-1].line - last_ast.tokens[last_ast.tokens.len()-1].line).as_str();
                    //     result += "\n".repeat(ast.tokens[ast.tokens.len()-1].line - last_ast.tokens[last_ast.tokens.len()-1].line).as_str();
                    // }
                    if ast.tokens[ast.tokens.len()-1].column > last_ast.tokens[last_ast.tokens.len()-1].column+fl {
                        result += " ".repeat(ast.tokens[ast.tokens.len()-1].column - (last_ast.tokens[last_ast.tokens.len()-1].column+fl)).as_str();
                    }
                }
                last_ast = Ast {ast_type: ast.ast_type.clone(), tokens: ast.tokens.clone()};

                if ast.ast_type == AstType::FunctionDeceleration {
                    result += format!(
                        "fn {}({}) -> {} {}",
                        ast.tokens[1].value,
                        transpile_round(ast.tokens[2].value.clone(), LexerState {
                            line: ast.tokens[2].line,
                            column: ast.tokens[2].column
                        }),
                        ast.tokens[0].value,
                        transpile(
                            ast.tokens[3].value.clone(),
                            indent + 1,
                            LexerState {
                                line: ast.tokens[3].line,
                                column: ast.tokens[3].column
                            },
                            &mut options.clone()
                        )
                    )
                    .as_str();
                } else if ast.ast_type == AstType::VoidFunctionDeceleration {
                    result += format!(
                        "fn {}({}) {}",
                        ast.tokens[1].value,
                        transpile_round(ast.tokens[2].value.clone(), LexerState {
                            line: ast.tokens[2].line,
                            column: ast.tokens[2].column
                        }),
                        transpile(
                            ast.tokens[3].value.clone(),
                            indent + 1,
                            LexerState {
                                line: ast.tokens[3].line,
                                column: ast.tokens[3].column
                            },
                            &mut options.clone()
                        )
                    )
                    .as_str();
                } else if ast.ast_type == AstType::StructDeceleration {
                    result += format!(
                        "struct {} {} {}",
                        ast.tokens[0].value,
                        "{\n",
                        transpile_round(
                            ast.tokens[1].value.clone(),
                            LexerState {
                                line: ast.tokens[1].line,
                                column: ast.tokens[1].column
                            }
                        ).trim_end()
                    ).replace("\n", ("\n".to_string() + " ".repeat(((indent+1) as usize)*2).as_str()).as_str())
                    .as_str();
                    result += "\n}\n";
                } else if ast.ast_type == AstType::VariableDeceleration {
                    if options.clone().auto_mut {
                        result += format!("let mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else {
                        result += format!("let {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    }
                } else if ast.ast_type == AstType::MutVariableDeceleration {
                    result += format!("let mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.ast_type == AstType::Other && ast.tokens[0].token_type == TokenType::Round {
                    result += format!(
                        "({})",
                        transpile_round(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                    result += format!(
                        "[{}]",
                        transpile_square(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.ast_type == AstType::CodeBlock {
                    result+="{";
                    result+=ast.tokens[0].value.as_str();
                    result+="}";
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                    // if ast.tokens[0].token_type == TokenType::Newline {
                    //     result += (ast.tokens[0].value.as_str().to_owned()
                    //         + (" ".repeat((indent as usize) * 2).as_str()))
                    //     .as_str();
                    // } else {
                        // result += ast.tokens[0].value.as_str();
                        result += transpile_json(ast.tokens[0].value.as_str().to_string(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        }).as_str();
                    // }
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                    result+=".";
                } else if ast.ast_type == AstType::StructCall {
                    result += ast.tokens[0].value.as_str();
                    result += " {";
                    result += ast.tokens[1].value.as_str();
                    result += "}";
                } else if ast.ast_type == AstType::StructVar {
                    result += format!("let mut {}: {} = {} {}{}{}", ast.tokens[1].value.as_str(), ast.tokens[0].value.as_str(), ast.tokens[0].value.as_str(), "{", ast.tokens[2].value.as_str(), "}").as_str();
                } else if ast.ast_type == AstType::Include {
                    let modname = transpile_mod(ast, options);
                    result += "mod ";
                    result += modname.as_str();
                    result += ";\n";
                    result += "use ";
                    result += modname.as_str();
                    result += "::*;\n";
                    options.modnum += 1;
                } // flp
                else {
                    if ast.tokens[0].token_type == TokenType::Newline {
                        result += (ast.tokens[0].value.as_str().to_owned()
                            + (" ".repeat((indent as usize) * 2).as_str()))
                        .as_str();
                    } else if ast.tokens[0].token_type == TokenType::Semicolon {
                        result += ";\n";
                        result += " ".repeat((indent as usize) * 2).as_str();
                    } else {
                            // if last_ast.tokens.len() > 0 && (
                            // last_ast.tokens[last_ast.tokens.len()-1].token_type == TokenType::Identifier ||
                            // last_ast.tokens[last_ast.tokens.len()-1].token_type == TokenType::Keyword ||
                            // last_ast.tokens[last_ast.tokens.len()-1].token_type == TokenType::Number
                        // ) {
                            // let ltkn = last_ast.tokens[last_ast.tokens.len()-1].token_type;
                            // if ltkn == TokenType::Identifier ||
                                // ltkn == TokenType::Keyword ||
                                // ltkn == TokenType::Number {
                            // }
                        // }
                        result += ast.tokens[0].value.as_str();
                        if options.auto_macro && options.macros.contains(&ast.tokens[0].value.as_str().to_string()) {
                            result += "!";
                        }
                    }
                }
            }
            
            result = result.trim_end().to_string();
            
            if indent > 0 {
                result += "\n";
                result += " ".repeat((indent as usize - 1) * 2).as_str();
                return "{\n".to_owned() + result.as_str() + "}";
            } else {
                return result;
            }
        }
        Err((state, _tokens)) => {
            panic!("Invalid syntax at code.ws:{}:{}", state.line, state.column);
        }
    }
}

pub fn transpile_round(input: String, state: LexerState) -> String {
    let mut result = String::new();
    let lexer_out = lex(input.as_str(), false, state);
    
    match lexer_out {
        Ok(tokens) => {
            let mut full_ast = Parser::new(tokens.clone());
            while full_ast.tokens.len() > full_ast.index as usize {
                let ast = full_ast.next();
                if ast.ast_type == AstType::VariableDeceleration {
                    result += format!("{}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.ast_type == AstType::MutVariableDeceleration {
                    result += format!("mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.ast_type == AstType::PointerDeceleration {
                    result += format!("{}: &mut {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                    result += format!(
                        "({})",
                        transpile_round(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                    result += format!(
                        "[{}]",
                        transpile_square(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                    result+=".";
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                    result += transpile_json(ast.tokens[0].value.as_str().to_string(), LexerState {
                        line: ast.tokens[0].line,
                        column: ast.tokens[0].column
                    }).as_str();
                } else if ast.ast_type == AstType::StructCall {
                    result += ast.tokens[0].value.as_str();
                    result += " {";
                    result += ast.tokens[1].value.as_str();
                    result += "}";
                } // flp
                else {
                    result += ast.tokens[0].value.as_str();
                    result += " ";
                }
            }
            
            result = result.trim_end().to_string();
            result
        }
        Err((state, _tokens)) => {
            panic!("Invalid syntax at code.ws:{}:{}", state.line, state.column);
        }
    }
}

pub fn transpile_square(input: String, state: LexerState) -> String {
    let mut result = String::new();
    let lexer_out = lex(input.as_str(), false, state);
    
    match lexer_out {
        Ok(tokens) => {
            let mut full_ast = Parser::new(tokens.clone());
            
            while full_ast.tokens.len() > full_ast.index as usize {
                let ast = full_ast.next();
                println!("{ast}");
                
                if ast.ast_type == AstType::VariableDeceleration {
                    result += format!("{}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.ast_type == AstType::MutVariableDeceleration {
                    result += format!("mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.ast_type == AstType::PointerDeceleration {
                    result += format!("{}: &mut {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                    result += format!(
                        "({})",
                        transpile_round(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                    result += format!(
                        "[{}]",
                        transpile_square(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                    result+=".";
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                    result += transpile_json(ast.tokens[0].value.as_str().to_string(), LexerState {
                        line: ast.tokens[0].line,
                        column: ast.tokens[0].column
                    }).as_str();
                } else if ast.ast_type == AstType::StructCall {
                    result += ast.tokens[0].value.as_str();
                    result += " {";
                    result += ast.tokens[1].value.as_str();
                    result += "}";
                } // flp
                else {
                    result += ast.tokens[0].value.as_str();
                    result += " ";
                }
            }
            
            result = result.trim_end().to_string();
            result
        }
        Err((state, _tokens)) => {
            panic!("Invalid syntax at code.ws:{}:{}", state.line, state.column);
        }
    }
}

pub fn transpile_json(input: String, state: LexerState) -> String {
    let mut result = String::new();
    let lexer_out = lex(input.as_str(), false, state);
    
    match lexer_out {
        Ok(tokens) => {
            let mut full_ast = Parser::new(tokens.clone());
            full_ast.json = true;
            result += "HashMap::from([";
            while full_ast.tokens.len() > full_ast.index as usize {
                let ast = full_ast.next();
                println!("{ast}");
                if ast.ast_type == AstType::Json {
                    result += "(";
                    result += ast.tokens[0].value.as_str();
                    result += ",";
                    result += ast.tokens[1].value.as_str();
                    result += ")";
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                    result += format!(
                        "({})",
                        transpile_round(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                    result += format!(
                        "[{}]",
                        transpile_square(ast.tokens[0].value.clone(), LexerState {
                            line: ast.tokens[0].line,
                            column: ast.tokens[0].column
                        })
                    )
                    .as_str();
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                    result += transpile_json(ast.tokens[0].value.as_str().to_string(), LexerState {
                        line: ast.tokens[0].line,
                        column: ast.tokens[0].column
                    }).as_str();
                } else {
                    result += ast.tokens[0].value.as_str();
                    result += " ";
                }
            }
            result += "])";
            result = result.trim_end().to_string();
            result
        }
        Err((state, _tokens)) => {
            panic!("Invalid syntax at code.ws:{}:{}", state.line, state.column);
        }
    }
}