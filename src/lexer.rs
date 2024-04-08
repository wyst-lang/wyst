use regex::Regex;
use once_cell::sync::Lazy;
use std::fmt;

pub struct LexerState {
    line: usize,
    column: usize
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
    Round,
    Curly,
    Square,
    Angle,
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
        write!(f, "(TokenType: {:?}, TokenValue: {}, Line: {}, Column {})", self.token_type, self.value, self.line, self.column)
    }
}

impl fmt::Display for Token {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "(\x1b[33mTokenType:\x1b[0m \x1b[35m{:?}\x1b[0m, \x1b[33mTokenValue:\x1b[0m \x1b[32m\"{}\"\x1b[0m, \x1b[33mLine:\x1b[0m \x1b[36m{}\x1b[0m, \x1b[33mColumn:\x1b[0m \x1b[36m{}\x1b[0m)", self.token_type, self.value, self.line, self.column)
    }
}

pub struct Node {
    token_type: TokenType,
    token_regex: Lazy<Regex>
}

const SYNTAX: [Node; 12] = [
    Node {
        token_type: TokenType::Newline,
        token_regex: Lazy::new(|| Regex::new(r"^\n+").unwrap()),
    },
    Node {
        token_type: TokenType::Whitespace,
        token_regex: Lazy::new(|| Regex::new(r"^\s+").unwrap())
    },
    Node {
        token_type: TokenType::Number,
        token_regex: Lazy::new(|| Regex::new(r"^\b(:?.)?(:?0[x|X])?\d+(:?.\d+)?\b").unwrap())
    },
    Node {
        token_type: TokenType::Keyword,
        token_regex: Lazy::new(|| Regex::new(r"^mut|try|catch|return|fn\b").unwrap())
    },
    Node {
        token_type: TokenType::Identifier,
        token_regex: Lazy::new(|| Regex::new(r"^[._a-zA-Z][a-zA-Z0-9_]*(:?<[._a-zA-Z][a-zA-Z0-9_<>]*>)?").unwrap())
    },
    Node {
        token_type: TokenType::Ptr,
        token_regex: Lazy::new(|| Regex::new(r"^->").unwrap())
    },
    Node {
        token_type: TokenType::Operator,
        token_regex: Lazy::new(|| Regex::new(r"^[\-|\+|\*]").unwrap())
    },
    Node {
        token_type: TokenType::Round,
        token_regex: Lazy::new(|| Regex::new(r"\((?:[^()]|(?R))*\)").unwrap())
    },
    Node {
        token_type: TokenType::Curly,
        token_regex: Lazy::new(|| Regex::new(r"\{(?:[^{}]|(?R))*}").unwrap())
    },
    Node {
        token_type: TokenType::Square,
        token_regex: Lazy::new(|| Regex::new(r"\[(?:[^\[\]]|(?R))*]").unwrap())
    },
    Node {
        token_type: TokenType::Angle,
        token_regex: Lazy::new(|| Regex::new(r"<(?:[^<>]|(?R))*>").unwrap())
    },
    Node {
        token_type: TokenType::SecondOperator,
        token_regex: Lazy::new(|| Regex::new(r"^,|;").unwrap())
    },
];

pub fn lex(mut code: &str, use_whitespace: bool) -> Vec<Token> {
    let mut state = LexerState { line: 1, column: 1 };
    let mut tokens: Vec<Token> = Vec::new();

    while !code.is_empty() {
        let mut is_match = false;
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
                    TokenType::Newline => {
                        state.line += caps[0].len();
                        state.column = 1;
                    },
                    _ => { state.column += caps[0].len(); }
                }
            } else {
                continue;
            };
            break;
        }
        if !is_match {
            println!("Error: syntax -> {code}");
            break;
        }
    }
    tokens
}
