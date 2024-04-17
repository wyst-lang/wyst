use crate::lexer::{lex, TokenType};
use crate::parser::{Parser, AstType};

pub fn transpile(input: String, indent: u32) -> String {
    let mut result = String::new();
    let lexer_out = lex(input.as_str(), false);
    match lexer_out {
        Ok(tokens) => {
            println!("\n\n\n\n");
            let mut full_ast = Parser::new(tokens.clone());
            while full_ast.tokens.len() > full_ast.index as usize {
                let ast = full_ast.next();
                println!("{ast}");
                if ast.ast_type == AstType::FunctionDeceleration {
                    result += format!("fn {}{} -> {} {}", ast.tokens[1].value, ast.tokens[2].value,  ast.tokens[0].value, transpile(ast.tokens[3].value.clone(), indent+1)).as_str();
                } else if ast.ast_type == AstType::VoidFunctionDeceleration {
                    result += format!("fn {}{} {}", ast.tokens[1].value, ast.tokens[2].value, transpile(ast.tokens[3].value.clone(), indent+1)).as_str();
                } else if ast.ast_type == AstType::VariableDeceleration {
                    let _o = String::new();
                    let mut oy = ast.tokens.clone();
                    oy.remove(0);
                    if oy.len() > 1 {
                        result += "(";
                        for t in &oy {
                            result+=t.value.as_str()
                        }
                        result += "): (";
                        for t in oy {
                            if t.value == "," {
                                result+=", "
                            } else {
                                result+=ast.tokens[0].value.as_str()
                            }
                        }
                        result += ")";
                    } else {
                        result += format!("let {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    }
                } else if ast.ast_type == AstType::MutVariableDeceleration {
                    let _o = String::new();
                    let mut oy = ast.tokens.clone();
                    oy.remove(0);
                    if oy.len() > 1 {
                        result += "(";
                        for t in &oy {
                            result+=t.value.as_str()
                        }
                        result += "): (";
                        for t in oy {
                            if t.value == "," {
                                result+=", "
                            } else {
                                result+=ast.tokens[0].value.as_str()
                            }
                        }
                        result += ")";
                    } else {
                        result += format!("let mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    }
                } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                    result += format!("({})", ast.tokens[0].value).as_str();
                } else {
                    if ast.tokens[0].token_type==TokenType::Newline {
                        result += (ast.tokens[0].value.as_str().to_owned() + (" ".repeat((indent as usize)*2).as_str())).as_str()
                    }
                    else {
                        result += ast.tokens[0].value.as_str()
                    }
                }
                
            }
            result = result.trim_end().to_string();
            if indent > 0 {
                result += "\n";
                result += " ".repeat((indent as usize-1)*2).as_str();
                return "{".to_owned()+result.as_str()+"}"
            } else {
                return result
            }
        },
        Err((state, _tokens)) => {
            println!("Invalid character at code.ws:{}:{}", state.line, state.column);
            return "".to_string();
        }
    }
    
}
