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
    Include,
    String,
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
        token_regex: Lazy::new(|| Regex::new(r"^mut|try|catch|return|fn\b").unwrap())
    },
    Node {
        token_type: TokenType::Identifier,
        token_regex: Lazy::new(|| Regex::new(r"^[._a-zA-Z][a-zA-Z0-9_]*(:?\s*<[._a-zA-Z][a-zA-Z0-9_<>]*>)?").unwrap())
    },
    Node {
        token_type: TokenType::Number,
        token_regex: Lazy::new(|| Regex::new(r"^(:?.)?(:?0[x|X])?\d+(:?.\d+)?\b").unwrap())
    },
    Node {
        token_type: TokenType::Ptr,
        token_regex: Lazy::new(|| Regex::new(r"^->").unwrap())
    },
    Node {
        token_type: TokenType::Operator,
        token_regex: Lazy::new(|| Regex::new(r"^[|\-|\+|\*|\=|\!]").unwrap())
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

pub fn lex(mut code: &str, use_whitespace: bool) -> Result<Vec<Token>, (LexerState, Vec<Token>)> {
    let mut state = LexerState { line: 1, column: 1 };
    let mut tokens: Vec<Token> = Vec::new();
    let mut brstr: String = String::new();
    let mut brvct: Vec<u8> = Vec::new();
    let mut inside_str: u8 = 0;
    while !code.is_empty() {
        let mut is_match = false;
        let fch = get_first_char(code);
        match fch.as_str() {
            "\"" => {
                brstr += "\"";
                code = code.strip_prefix(fch.as_str()).expect("");
                if inside_str == 1 {
                    inside_str = 0;
                    if brvct.len() == 0 {
                        tokens.push(Token {
                            token_type: TokenType::String,
                            value: brstr.clone(),
                            column: state.line,
                            line: state.column
                        });
                        brstr = String::new();
                    }
                } else if inside_str == 0 {
                    inside_str = 1;
                }
            }
            "'" => {
                brstr += "'";
                code = code.strip_prefix(fch.as_str()).expect("");
                if inside_str == 2 {
                    inside_str = 0;
                    if brvct.len() == 0 {
                        tokens.push(Token {
                            token_type: TokenType::String,
                            value: brstr.clone(),
                            column: state.line,
                            line: state.column
                        });
                        brstr = String::new();
                    }
                } else if inside_str == 0 {
                    inside_str = 2;
                }
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
            "{" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brvct.len() > 0 {
                    brstr += "{";
                }
                brvct.push(0);
            }
            "}" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                brvct.pop();
                if brvct.len() == 0 {
                    tokens.push(Token {
                        token_type: TokenType::Curly,
                        value: brstr.clone(),
                        column: state.line,
                        line: state.column
                    });
                } else {
                    brstr += "}";
                }
            }
            "(" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brvct.len() > 0 {
                    brstr += "(";
                }
                brvct.push(1);
            }
            ")" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                brvct.pop();
                if brvct.len() == 0 {
                    tokens.push(Token {
                        token_type: TokenType::Round,
                        value: brstr.clone(),
                        column: state.line,
                        line: state.column
                    });
                } else {
                    brstr += ")";
                }
            }
            "[" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                if brvct.len() > 0 {
                    brstr += "[";
                }
                brvct.push(2);
            }
            "]" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                brvct.pop();
                if brvct.len() == 0 {
                    tokens.push(Token {
                        token_type: TokenType::Square,
                        value: brstr.clone(),
                        column: state.line,
                        line: state.column
                    });
                } else {
                    brstr += "]";
                }
            }
            "\n" => {
                code = code.strip_prefix(fch.as_str()).expect("");
                state.line += 1;
            }
            _ => {
                if inside_str > 0 || brvct.len() > 0 {
                    code = code.strip_prefix(fch.as_str()).expect("");
                    brstr += fch.as_str();
                } else {
                    for s in &SYNTAX {
                        if let Some(caps) = s.token_regex.captures(code) {
                            is_match = true;
                            code = code.strip_prefix(&caps[0]).unwrap_or(code);
                            if (!use_whitespace && s.token_type!=TokenType::Whitespace) || use_whitespace {
                                tokens.push(Token {
                                    token_type: s.token_type,
                                    value: caps[0].to_string(),
                                    line: state.line,
                                    column: state.column
                                });
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
