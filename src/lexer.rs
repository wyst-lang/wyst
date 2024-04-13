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

fn get_second_char(value: &str) -> String {
    let mut sv2 = value.chars();
    sv2.next();
    sv2.next().unwrap().to_string()
}

pub fn lex(mut code: &str, use_whitespace: bool) -> Vec<Token> {
    let mut state = LexerState { line: 1, column: 1 };
    let mut tokens: Vec<Token> = Vec::new();
    let mut bracket_vec: Vec<char> = Vec::new();
    let mut inner_string = String::new();
    // let mut bracket_state: LexerState = LexerState {line: 0, column: 0};
    while !code.is_empty() {
        let mut is_match = false;
        let svl = bracket_vec.len();
        match code.chars().next().unwrap() {
            '\"' => {
                is_match = true;
                code = code.strip_prefix("\"").unwrap_or(code);
                inner_string += "\"";
                if svl == 0 {
                    bracket_vec.push('\"');
                } else if bracket_vec[svl-1] == '\"' {
                    bracket_vec.pop();
                    if svl == 1 {
                        tokens.push(Token {
                            token_type:TokenType::String,
                            value: inner_string.clone(),
                            line: state.line,
                            column: state.column
                        });
                        inner_string = String::new();
                    }
                }
            }
            '\'' => {
                is_match = true;
                code = code.strip_prefix("\'").unwrap_or(code);
                inner_string += "\'";
                if svl == 0 {
                    bracket_vec.push('\'');
                } else if bracket_vec[svl-1] == '\'' {
                    bracket_vec.pop();
                    if svl == 1 {
                        tokens.push(Token {
                            token_type:TokenType::String,
                            value: inner_string.clone(),
                            line: state.line,
                            column: state.column
                        });
                        inner_string = String::new();
                    }
                }
            }
            '{' => {
                is_match = true;
                code = code.strip_prefix("{").unwrap_or(code);
                if svl == 0 {
                    bracket_vec.push('{');
                }
            }
            '}' => {
                is_match = true;
                code = code.strip_prefix("}").unwrap_or(code);
                if bracket_vec[svl-1] == '{' {
                    bracket_vec.pop();
                    tokens.push(Token {
                        token_type: TokenType::Curly,
                        value: inner_string.clone(),
                        line: state.line,
                        column: state.column
                    });
                    inner_string = String::new();
                    state.column += inner_string.len();
                }
            }
            '(' => {
                is_match = true;
                code = code.strip_prefix("(").unwrap_or(code);
                if svl == 0 {
                    bracket_vec.push('(');
                }
            }
            ')' => {
                is_match = true;
                code = code.strip_prefix(")").unwrap_or(code);
                if bracket_vec[svl-1] == '(' {
                    bracket_vec.pop();
                    tokens.push(Token {
                        token_type:TokenType::Round,
                        value: inner_string.clone(),
                        line: state.line,
                        column: state.column
                    });
                    inner_string = String::new();
                    state.column += inner_string.len();
                }
            }
            '[' => {
                is_match = true;
                code = code.strip_prefix("[").unwrap_or(code);
                
                if svl == 0 {
                    bracket_vec.push('[');
                }
            }
            ']' => {
                is_match = true;
                code = code.strip_prefix("]").unwrap_or(code);
                if bracket_vec[svl-1] == '[' {
                    bracket_vec.pop();
                    tokens.push(Token {
                        token_type:TokenType::Square,
                        value: inner_string.clone(),
                        line: state.line,
                        column: state.column
                    });
                    inner_string = String::new();
                    state.column += inner_string.len();
                }
            }
            '\n' => {
                is_match = true;
                code = code.strip_prefix("\n").unwrap_or(code);
                state.line += 1;
                state.column = 1;
                if svl > 0 {
                    inner_string+="\n";
                } else {
                    tokens.push(Token {
                        token_type: TokenType::Newline,
                        value: "\n".to_string(),
                        line: state.line,
                        column: state.column
                    });
                }
            }
            _ => {
                if svl > 0 {
                    let cfc = get_first_char(&code);
                    let cfc1 = get_second_char(&code);
                    inner_string += cfc.as_str();
                    code = code.strip_prefix(&cfc).unwrap_or(code);
                    is_match = true;
                    if cfc == "\\" && (cfc1 == "\"" || cfc1 == "'") {
                        inner_string += cfc1.as_str();
                        code = code.strip_prefix(&cfc1).unwrap_or(code);
                        is_match = true;
                    }
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
                            match s.token_type {
                                _ => {
                                    state.column += caps[0].len();
                                }
                            }
                        } else {
                            continue;
                        };
                        break;
                    }
                }
                
            }
        }
        if !is_match {
            println!("Error: syntax ->{code}");
            break;
        }
    }
    tokens
}
