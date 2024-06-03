use crate::lexer::{lex, LexerState, TokenType};
use crate::parser::{Ast, AstType, Parser};
use std::collections::HashMap;
use std::fs;

#[derive(Debug, PartialEq, Clone)]
pub enum VarTypes {
    Var
}

#[derive(Debug, PartialEq, Clone)]
pub struct Var {
    vtype: VarTypes,
    state: LexerState
}

#[derive(Debug, PartialEq, Clone)]
pub struct Transpiler {
    pub auto_mut: bool,
    pub auto_macro: bool,
    pub auto_pub: bool,
    pub macros: Vec<String>,
    pub modnum: u32,
    pub var_match: String,
    pub var_state: LexerState,
    pub matched_vars: HashMap<String, Var>,
}

impl Default for Transpiler {
    fn default() -> Transpiler {
        Transpiler {
            auto_mut: true,
            auto_macro: true,
            auto_pub: false,
            macros: vec![String::from("println")],
            modnum: 0,
            var_state: LexerState { line: 0, column: 0 },
            var_match: String::new(),
            matched_vars: HashMap::new()
        }
    }
}

impl Transpiler {
    pub fn transpile(&mut self, input: String, indent: u32, state: LexerState) -> String {
        let mut result = String::new();
        let mut variables: HashMap<String, Var> = HashMap::new();
    
        if indent == 0 {
            // result += "type int = i32;\n";
        } else {
            result += " ".repeat((indent as usize) * 2).as_str();
        }
    
        let lexer_out = lex(input.as_str(), false, state);
    
        match lexer_out {
            Ok(tokens) => {
                let mut full_ast = Parser::new(tokens.clone());
                let mut last_ast = Ast {
                    ast_type: AstType::Other,
                    tokens: vec![],
                };
                while full_ast.tokens.len() > full_ast.index as usize {
                    let ast = full_ast.next();
                    if last_ast.tokens.len() > 0 {
                        let mut fl = 0;
                        for t in &last_ast.tokens {
                            fl += t.value.len()
                        }
                        if ast.tokens[ast.tokens.len() - 1].column
                            > last_ast.tokens[last_ast.tokens.len() - 1].column + fl
                        {
                            result += " "
                                .repeat(
                                    ast.tokens[ast.tokens.len() - 1].column
                                        - (last_ast.tokens[last_ast.tokens.len() - 1].column + fl),
                                )
                                .as_str();
                        }
                    }
                    last_ast = Ast {
                        ast_type: ast.ast_type.clone(),
                        tokens: ast.tokens.clone(),
                    };
    
                    if ast.ast_type == AstType::FunctionDeceleration {
                        if self.auto_pub {
                            result += "pub ";
                        }
                        result += format!(
                            "fn {}({}) -> {} {}",
                            ast.tokens[1].value,
                            self.transpile_round(
                                ast.tokens[2].value.clone(),
                                LexerState {
                                    line: ast.tokens[2].line,
                                    column: ast.tokens[2].column
                                }
                            ),
                            ast.tokens[0].value,
                            self.transpile(
                                ast.tokens[3].value.clone(),
                                indent + 1,
                                LexerState {
                                    line: ast.tokens[3].line,
                                    column: ast.tokens[3].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::VoidFunctionDeceleration {
                        if self.auto_pub {
                            result += "pub ";
                        }
                        result += format!(
                            "fn {}({}) {}",
                            ast.tokens[1].value,
                            self.transpile_round(
                                ast.tokens[2].value.clone(),
                                LexerState {
                                    line: ast.tokens[2].line,
                                    column: ast.tokens[2].column
                                }
                            ),
                            self.transpile(
                                ast.tokens[3].value.clone(),
                                indent + 1,
                                LexerState {
                                    line: ast.tokens[3].line,
                                    column: ast.tokens[3].column
                                },
                            )
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::StructDeceleration {
                        if self.auto_pub {
                            result += "pub ";
                        }
                        result += format!(
                            "struct {} {} {}",
                            ast.tokens[0].value,
                            "{\n",
                            self.transpile_round(
                                ast.tokens[1].value.clone(),
                                LexerState {
                                    line: ast.tokens[1].line,
                                    column: ast.tokens[1].column
                                }
                            )
                            .trim_end()
                        )
                        .replace(
                            "\n",
                            ("\n".to_string() + " ".repeat(((indent + 1) as usize) * 2).as_str())
                                .as_str(),
                        )
                        .as_str();
                        result += "\n}\n";
                    } else if ast.ast_type == AstType::VariableDeceleration {
                        variables.insert(ast.tokens[1].value.clone(), Var {vtype: VarTypes::Var, state: LexerState {
                            line: ast.tokens[1].line,
                            column: ast.tokens[1].column,
                        }});
                        if self.clone().auto_mut {
                            result +=
                                format!("let mut {}: {}", ast.tokens[1].value, ast.tokens[0].value)
                                    .as_str();
                        } else {
                            result += format!("let {}: {}", ast.tokens[1].value, ast.tokens[0].value)
                                .as_str();
                        }
                    } else if ast.ast_type == AstType::MutVariableDeceleration {
                        result += format!("let mut {}: {}", ast.tokens[1].value, ast.tokens[0].value)
                            .as_str();
                    } else if ast.ast_type == AstType::Other
                        && ast.tokens[0].token_type == TokenType::Round
                    {
                        result += format!(
                            "({})",
                            self.transpile_round(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                        result += format!(
                            "[{}]",
                            self.transpile_square(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::CodeBlock {
                        result += "{";
                        result += ast.tokens[0].value.as_str();
                        result += "}";
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                        // if ast.tokens[0].token_type == TokenType::Newline {
                        //     result += (ast.tokens[0].value.as_str().to_owned()
                        //         + (" ".repeat((indent as usize) * 2).as_str()))
                        //     .as_str();
                        // } else {
                        // result += ast.tokens[0].value.as_str();
                        result += self.transpile_json(
                            ast.tokens[0].value.as_str().to_string(),
                            LexerState {
                                line: ast.tokens[0].line,
                                column: ast.tokens[0].column,
                            },
                        )
                        .as_str();
                        // }
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                        result += ".";
                    } else if ast.ast_type == AstType::StructCall {
                        result += ast.tokens[0].value.as_str();
                        result += " {";
                        result += ast.tokens[1].value.as_str();
                        result += "}";
                    } else if ast.ast_type == AstType::StructVar {
                        result += format!(
                            "let mut {}: {} = {} {}{}{}",
                            ast.tokens[1].value.as_str(),
                            ast.tokens[0].value.as_str(),
                            ast.tokens[0].value.as_str(),
                            "{",
                            ast.tokens[2].value.as_str(),
                            "}"
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::Include {
                        let modname = self.transpile_mod(ast, "lib/");
                        result += "mod ";
                        result += modname.as_str();
                        result += ";\n";
                        result += "use ";
                        result += modname.as_str();
                        result += "::*;\n";
                        self.modnum += 1;
                    } else if ast.ast_type == AstType::IncludeLocal {
                        let modname = self.transpile_mod(ast, "");
                        result += "mod ";
                        result += modname.as_str();
                        result += ";\n";
                        result += "use ";
                        result += modname.as_str();
                        result += "::*;\n";
                        self.modnum += 1;
                    } else if ast.ast_type == AstType::State3 {
                        result += format!(
                            "{} {} {}",
                            ast.tokens[0].value.clone(),
                            self.transpile_round(
                                ast.tokens[1].value.clone(),
                                LexerState {
                                    line: ast.tokens[1].line,
                                    column: ast.tokens[1].column
                                }
                            ),
                            self.transpile(
                                ast.tokens[2].value.clone(),
                                indent + 1,
                                LexerState {
                                    line: ast.tokens[2].line,
                                    column: ast.tokens[2].column
                                }
                            ),
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::State2 {
                        result += format!(
                            "{} {}",
                            ast.tokens[0].value.clone(),
                            self.transpile(
                                ast.tokens[1].value.clone(),
                                indent + 1,
                                LexerState {
                                    line: ast.tokens[1].line,
                                    column: ast.tokens[1].column
                                },
                            ),
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::Namespace {
                        result += format!(
                            "mod {} {}{}{}",
                            &ast.tokens[0].value.clone(),
                            "{",
                            self.transpile(
                                ast.tokens[1].value.clone(),
                                0,
                                state,
                            ),
                            "}"
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::Impl {
                        result += format!(
                            "impl {} {}{}{}",
                            &ast.tokens[0].value.clone(),
                            "{",
                            self.transpile(
                                ast.tokens[1].value.clone(),
                                0,
                                state,
                            ),
                            "}"
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::PointerDeceleration {
                        if self.auto_mut {
                            result +=
                                format!("let mut {}: &mut {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                        } else {
                            result +=
                               format!("let {}: &mut {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                        }
                    }
                    // flp
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
                            if self.auto_macro
                                && self
                                    .macros
                                    .contains(&ast.tokens[0].value.as_str().to_string())
                            {
                                result += "!";
                            }
                        }
                    }
                }
    
                result = result.trim_end().to_string();
    
                if self.var_match != "" && (state.line + input.len()) > self.var_state.line {
                    for (name, var) in variables {
                        if (self.var_state.line > var.state.line ||
                            (self.var_state.line == var.state.line && self.var_state.column > var.state.column)) &&
                            name.to_lowercase().starts_with(&self.var_match) {
                                self.matched_vars.insert(name, var);
                        }
                    }
                    self.var_match = String::new();
                }
    
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
    pub fn transpile_mod(&mut self, ast: Ast, s: &str) -> String {
        let modfile = ast.tokens[0].value.as_str();
        let modname = format!(
            "{}_{}",
            clean_incl(modfile.split(".").collect::<Vec<_>>()[0]),
            self.clone().modnum
        );
        let file_content = fs::read_to_string(s.to_string() + modfile).expect("Error reading file");
        let transpiled_code = self.transpile(
            file_content,
            0,
            LexerState { line: 1, column: 0 }
        );
        fs::write(
            ("build/".to_string() + modname.as_str()) + ".rs",
            transpiled_code,
        )
        .expect("Error writing file");
        modname
    }
    pub fn transpile_round(&mut self, input: String, state: LexerState) -> String {
        let mut result = String::new();
        let lexer_out = lex(input.as_str(), false, state);
    
        match lexer_out {
            Ok(tokens) => {
                let mut full_ast = Parser::new(tokens.clone());
                let mut last_ast = Ast {
                    ast_type: AstType::Other,
                    tokens: vec![],
                };
                while full_ast.tokens.len() > full_ast.index as usize {
                    let ast = full_ast.next();
                    if last_ast.tokens.len() > 0 {
                        let mut fl = 0;
                        for t in &last_ast.tokens {
                            fl += t.value.len()
                        }
                        if ast.tokens[ast.tokens.len() - 1].column
                            > last_ast.tokens[last_ast.tokens.len() - 1].column + fl
                        {
                            result += " "
                                .repeat(
                                    ast.tokens[ast.tokens.len() - 1].column
                                        - (last_ast.tokens[last_ast.tokens.len() - 1].column + fl),
                                )
                                .as_str();
                        }
                    }
                    last_ast = Ast {
                        ast_type: ast.ast_type.clone(),
                        tokens: ast.tokens.clone(),
                    };
                    if ast.ast_type == AstType::VariableDeceleration {
                        result += format!("{}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else if ast.ast_type == AstType::MutVariableDeceleration {
                        result +=
                            format!("mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else if ast.ast_type == AstType::PointerDeceleration {
                        result +=
                            format!("{}: &mut {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                        result += format!(
                            "({})",
                            self.transpile_round(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                        result += format!(
                            "[{}]",
                            self.transpile_square(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                        result += ".";
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                        result += self.transpile_json(
                            ast.tokens[0].value.as_str().to_string(),
                            LexerState {
                                line: ast.tokens[0].line,
                                column: ast.tokens[0].column,
                            },
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::StructCall {
                        result += ast.tokens[0].value.as_str();
                        result += " {";
                        result += ast.tokens[1].value.as_str();
                        result += "}";
                    } else if ast.ast_type == AstType::Ref {
                        result += "&mut ";
                        result += ast.tokens[0].value.as_str();
                    }
                    // flp
                    else {
                        result += ast.tokens[0].value.as_str();
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
    
    pub fn transpile_square(&mut self, input: String, state: LexerState) -> String {
        let mut result = String::new();
        let lexer_out = lex(input.as_str(), false, state);
    
        match lexer_out {
            Ok(tokens) => {
                let mut full_ast = Parser::new(tokens.clone());
                let mut last_ast = Ast {
                    ast_type: AstType::Other,
                    tokens: vec![],
                };
                while full_ast.tokens.len() > full_ast.index as usize {
                    let ast = full_ast.next();
                    if last_ast.tokens.len() > 0 {
                        let mut fl = 0;
                        for t in &last_ast.tokens {
                            fl += t.value.len()
                        }
                        if ast.tokens[ast.tokens.len() - 1].column
                            > last_ast.tokens[last_ast.tokens.len() - 1].column + fl
                        {
                            result += " "
                                .repeat(
                                    ast.tokens[ast.tokens.len() - 1].column
                                        - (last_ast.tokens[last_ast.tokens.len() - 1].column + fl),
                                )
                                .as_str();
                        }
                    }
                    last_ast = Ast {
                        ast_type: ast.ast_type.clone(),
                        tokens: ast.tokens.clone(),
                    };
    
                    if ast.ast_type == AstType::VariableDeceleration {
                        result += format!("{}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else if ast.ast_type == AstType::MutVariableDeceleration {
                        result +=
                            format!("mut {}: {}", ast.tokens[1].value, ast.tokens[0].value).as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                        result += format!(
                            "({})",
                            self.transpile_round(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                        result += format!(
                            "[{}]",
                            self.transpile_square(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Ptr {
                        result += ".";
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                        result += self.transpile_json(
                            ast.tokens[0].value.as_str().to_string(),
                            LexerState {
                                line: ast.tokens[0].line,
                                column: ast.tokens[0].column,
                            },
                        )
                        .as_str();
                    } else if ast.ast_type == AstType::StructCall {
                        result += ast.tokens[0].value.as_str();
                        result += " {";
                        result += ast.tokens[1].value.as_str();
                        result += "}";
                    }
                    // flp
                    else {
                        result += ast.tokens[0].value.as_str();
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
    
    pub fn transpile_json(&mut self, input: String, state: LexerState) -> String {
        let mut result = String::new();
        let lexer_out = lex(input.as_str(), false, state);
    
        match lexer_out {
            Ok(tokens) => {
                let mut full_ast = Parser::new(tokens.clone());
                full_ast.json = true;
                result += "HashMap::from([";
                let mut last_ast = Ast {
                    ast_type: AstType::Other,
                    tokens: vec![],
                };
                while full_ast.tokens.len() > full_ast.index as usize {
                    let ast = full_ast.next();
                    if last_ast.tokens.len() > 0 {
                        let mut fl = 0;
                        for t in &last_ast.tokens {
                            fl += t.value.len()
                        }
                        if ast.tokens[ast.tokens.len() - 1].column
                            > last_ast.tokens[last_ast.tokens.len() - 1].column + fl
                        {
                            result += " "
                                .repeat(
                                    ast.tokens[ast.tokens.len() - 1].column
                                        - (last_ast.tokens[last_ast.tokens.len() - 1].column + fl),
                                )
                                .as_str();
                        }
                    }
                    last_ast = Ast {
                        ast_type: ast.ast_type.clone(),
                        tokens: ast.tokens.clone(),
                    };
    
                    if ast.ast_type == AstType::Json {
                        result += "(";
                        result += ast.tokens[0].value.as_str();
                        result += ",";
                        result += ast.tokens[1].value.as_str();
                        result += ")";
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Round {
                        result += format!(
                            "({})",
                            self.transpile_round(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Square {
                        result += format!(
                            "[{}]",
                            self.transpile_square(
                                ast.tokens[0].value.clone(),
                                LexerState {
                                    line: ast.tokens[0].line,
                                    column: ast.tokens[0].column
                                }
                            )
                        )
                        .as_str();
                    } else if ast.tokens.len() == 1 && ast.tokens[0].token_type == TokenType::Curly {
                        result += self.transpile_json(
                            ast.tokens[0].value.as_str().to_string(),
                            LexerState {
                                line: ast.tokens[0].line,
                                column: ast.tokens[0].column,
                            },
                        )
                        .as_str();
                    } else {
                        result += ast.tokens[0].value.as_str();
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
}

fn clean_incl(input: &str) -> String {
    input
        .chars()
        .map(|c| {
            if c.is_alphanumeric() || c == '_' {
                c
            } else {
                '_'
            }
        })
        .collect()
}