use regex::Regex;
use once_cell::sync::Lazy;
use std::fmt;
pub struct LexerState {
    pub line: usize,
    pub column: usize
}

// Define token types
#[derive(Debug, PartialEq, Clone, Copy)]
pub enum TokenType {
    Keyword,
    Newline,
    Whitespace,
    Number,
    Identifier,
    Ptr,
    Operator,
    SecondOperator,
    Semicolon,
    Round,
    Curly,
    Square,
    Angle,
    Include,
    String,
    Comment,
    // EOF,
}

#[derive(Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub value: String,
    pub line: usize,
    pub column: usize
}

impl fmt::Debug for Token {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "(TokenType: {:?}, TokenValue: {}, Line: {}, Column {})", self.token_type, self.value.replace("\n", "\\n"), self.line, self.column)
    }
}

impl fmt::Display for Token {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "(\x1b[33mTokenType:\x1b[0m \x1b[35m{:?}\x1b[0m, \x1b[33mTokenValue:\x1b[0m \x1b[32m\"{}\"\x1b[0m, \x1b[33mLine:\x1b[0m \x1b[36m{}\x1b[0m, \x1b[33mColumn:\x1b[0m \x1b[36m{}\x1b[0m)", self.token_type, self.value.replace("\n", "\\n"), self.line, self.column)
    }
}

pub struct Node {
    token_type: TokenType,
    token_regex: Lazy<Regex>
}

const SYNTAX: [Node; 14] = [
    Node {
        token_type: TokenType::Semicolon,
        token_regex: Lazy::new(|| Regex::new(r"^\;").unwrap())
    },
    Node {
        token_type: TokenType::SecondOperator,
        token_regex: Lazy::new(|| Regex::new(r"^,").unwrap())
    },
    Node {
        token_type: TokenType::String,
        token_regex: Lazy::new(|| Regex::new("^\"").unwrap())
    },
    Node {
        token_type: TokenType::Whitespace,
        token_regex: Lazy::new(|| Regex::new(r"^\s+").unwrap())
    },
    Node {
        token_type: TokenType::Keyword,
        token_regex: Lazy::new(|| Regex::new(r"^(mut|try|catch|return|fn|let|use|cb)\b").unwrap())
    },
    Node {
        token_type: TokenType::Identifier,
        token_regex: Lazy::new(|| Regex::new(r"^[._a-zA-Z][a-zA-Z0-9_]*").unwrap())
    },
    Node {
        token_type: TokenType::Number,
        token_regex: Lazy::new(|| Regex::new(r"^\d+").unwrap())
    },
    Node {
        token_type: TokenType::Ptr,
        token_regex: Lazy::new(|| Regex::new(r"^->").unwrap())
    },
    Node {
        token_type: TokenType::Operator,
        token_regex: Lazy::new(|| Regex::new(r"^[|\-|\+|\*|\=|\!|\&|\:]").unwrap())
    },
    Node {
        token_type: TokenType::Round,
        token_regex: Lazy::new(|| Regex::new(r"^[\(|\)]").unwrap())
    },
    Node {
        token_type: TokenType::Curly,
        token_regex: Lazy::new(|| Regex::new(r"^[\{|\}]").unwrap())
    },
    Node {
        token_type: TokenType::Square,
        token_regex: Lazy::new(|| Regex::new(r"^[\[|\]]").unwrap())
    },
    Node {
        token_type: TokenType::Include,
        token_regex: Lazy::new(|| Regex::new(r"^#include *<(.*?)>").unwrap())
    },
    Node {
        token_type: TokenType::Include,
        token_regex: Lazy::new(|| Regex::new(r#"^#include *"(.*?)""#).unwrap())
    }
];

fn get_first_char(value: &str) -> String {
    value.chars().next().unwrap().to_string()
}

pub fn lex(mut code: &str, use_whitespace: bool, state: LexerState) -> Result<Vec<Token>, (LexerState, Vec<Token>)> {
    let mut state = state;
    let mut tokens: Vec<Token> = Vec::new();
    let mut brstr: String = String::new();
    let mut br_state: LexerState = LexerState { line: 0, column: 0 };
    let mut brtp: Vec<u8> = Vec::new();
    while !code.is_empty() {
        let mut is_match = false;
        let fch = get_first_char(code);
        let brln = brtp.len();
        match fch.as_str() {
            "/" => {
                brstr += "/";
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln == 0 { 
                    if code.len() > 0 {
                        let sch = get_first_char(code);
                        if sch=="/" {
                            code = code.strip_prefix(sch.as_str()).expect("");
                            brstr += &sch;
                            brtp.push(4);
                            if brln > 0 {
                                brstr += fch.as_str();
                            } else {
                                br_state.line = state.line;
                                br_state.column = state.column;
                            }
                        } else if sch=="*" {
                            code = code.strip_prefix(sch.as_str()).expect("");
                            brstr += &sch;
                            brtp.push(5);
                            if brln > 0 {
                                brstr += fch.as_str();
                            } else {
                                br_state.line = state.line;
                                br_state.column = state.column;
                            }
                        } else if brln == 0 {
                            tokens.push(Token {
                                token_type: TokenType::Operator,
                                value: "/".to_string(),
                                column: state.column,
                                line: state.line
                            });
                        }
                    } else if brln == 0 {
                        tokens.push(Token {
                            token_type: TokenType::Operator,
                            value: "*".to_string(),
                            column: state.column,
                            line: state.line
                        });
                    }
                }
            }
            "*" => {
                brstr += "*";
                code = code.strip_prefix(fch.as_str()).expect("");
                if code.len() > 0 {
                    let sch = get_first_char(code);
                    if sch=="/" && brtp[brln-1]==5 {
                        code = code.strip_prefix(sch.as_str()).expect("");
                        brstr += &sch;
                        brtp.pop();
                        tokens.push(Token {
                            token_type: TokenType::Comment,
                            value: brstr.clone(),
                            column: br_state.column,
                            line: br_state.line
                        });
                        brstr = String::new();
                    } else if brln == 0 {
                        tokens.push(Token {
                            token_type: TokenType::Operator,
                            value: "*".to_string(),
                            column: state.column,
                            line: state.line
                        });
                    }
                } else if brln == 0 {
                    tokens.push(Token {
                        token_type: TokenType::Operator,
                        value: "*".to_string(),
                        column: state.column,
                        line: state.line
                    });
                }
            }
            "\"" => {
                brstr += "\"";
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 0 {
                    brtp.pop();
                    tokens.push(Token {
                        token_type: TokenType::String,
                        value: brstr.clone(),
                        column: br_state.column,
                        line: br_state.line
                    });
                    brstr = String::new();
                } else if brln == 0 {
                    brtp.push(0);
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
            }
            "'" => {
                brstr += "'";
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 1 {
                    brtp.pop();
                    tokens.push(Token {
                        token_type: TokenType::String,
                        value: brstr.clone(),
                        column: br_state.column,
                        line: br_state.line
                    });
                    brstr = String::new();
                } else if brln == 0 {
                    brtp.push(1);
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
            }
            "(" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 {
                    brstr += fch.as_str();
                } else {
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
                brtp.push(2);
            }
            ")" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 2 {
                    brtp.pop();
                    if brln == 1 {
                        tokens.push(Token {
                            token_type: TokenType::Round,
                            value: brstr.clone(),
                            column: br_state.column,
                            line: br_state.line
                        });
                        brstr = String::new();
                    } else {brstr += fch.as_str();}
                } else {brstr += fch.as_str();}
            }
            "{" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 {
                    brstr += fch.as_str();
                } else {
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
                brtp.push(3);
            }
            "}" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 3 {
                    brtp.pop();
                    if brln == 1 {
                        tokens.push(Token {
                            token_type: TokenType::Curly,
                            value: brstr.clone(),
                            column: br_state.column,
                            line: br_state.line
                        });
                        brstr = String::new();
                    } else {brstr += fch.as_str();}
                } else {brstr += fch.as_str();}
            }
            "[" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 {
                    brstr += fch.as_str();
                } else {
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
                brtp.push(3);
            }
            "]" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 3 {
                    brtp.pop();
                    if brln == 1 {
                        tokens.push(Token {
                            token_type: TokenType::Square,
                            value: brstr.clone(),
                            column: br_state.column,
                            line: br_state.line
                        });
                        brstr = String::new();
                    } else {brstr += fch.as_str();}
                } else {brstr += fch.as_str();}
            }
            "<" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 {
                    brstr += fch.as_str();
                } else {
                    br_state.line = state.line;
                    br_state.column = state.column;
                }
                brtp.push(6);
            }
            ">" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 && brtp[brln-1] == 6 {
                    brtp.pop();
                    if brln == 1 {
                        tokens.push(Token {
                            token_type: TokenType::Angle,
                            value: brstr.clone(),
                            column: br_state.column,
                            line: br_state.line
                        });
                        brstr = String::new();
                    } else {brstr += fch.as_str();}
                } else {brstr += fch.as_str();}
            }
            "\\" => {
                brstr += "\\";
                code = code.strip_prefix(fch.as_str()).expect("");
                if code.len() > 0 {
                    let sch = get_first_char(code);
                    code = code.strip_prefix(sch.as_str()).expect("");
                    brstr += &sch;
                }
            }
            "\n" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brln > 0 {
                    brstr += "\n";
                } if brln==1 && brtp[0]==4 {
                    brtp.pop();
                    tokens.push(Token {
                        token_type: TokenType::Comment,
                        value: brstr.clone(),
                        column: br_state.column,
                        line: br_state.line
                    });
                    brstr = String::new();
                }
                state.line += 1;
                state.column = 0;
            }
            _ => {
                if brln > 0 {
                    code = code.strip_prefix(fch.as_str()).expect("");
                    brstr += fch.as_str();
                } else {
                    for s in &SYNTAX {
                        if let Some(caps) = s.token_regex.captures(code) {
                            is_match = true;
                            code = code.strip_prefix(&caps[0]).unwrap_or(code);
                            if (!use_whitespace && s.token_type!=TokenType::Whitespace) || use_whitespace {
                                let cap = caps[0].to_string();
                                match cap.as_str() {
                                    "int" => {
                                        tokens.push(Token {
                                            token_type: s.token_type,
                                            value: "i32".to_string(),
                                            line: state.line,
                                            column: state.column
                                        });
                                    }
                                    _ => {
                                        tokens.push(Token {
                                            token_type: s.token_type,
                                            value: cap,
                                            line: state.line,
                                            column: state.column
                                        });
                                    }
                                }
                            }
                            state.column += caps[0].len();
                        } else {
                            continue;
                        };
                        break;
                    }
                    if !is_match {
                        return Err((state, tokens));
                    }
                }
            }
        }
    }
    Ok(tokens)
}
