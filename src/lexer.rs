use regex::Regex;
use once_cell::sync::Lazy;
use std::fmt;

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
    Round,
    Curly,
    // EOF,
}

#[derive(Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub token_values: Vec<String>
}

impl<'a> fmt::Debug for Token {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "(TokenType: {:?}, TokenValue: {})", self.token_type, self.token_values.join(", "))
    }
}

pub struct Node {
    token_type: TokenType,
    token_regex: Lazy<Regex>
}

const SYNTAX: [Node; 9] = [
    Node {
        token_type: TokenType::Newline,
        token_regex: Lazy::new(|| Regex::new(r"^\n+").unwrap())
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
        token_regex: Lazy::new(|| Regex::new(r"^if|mut|try|catch|return|fn").unwrap())
    },
    Node {
        token_type: TokenType::Identifier,
        token_regex: Lazy::new(|| Regex::new(r"^[._a-zA-Z][a-zA-Z0-9_]*").unwrap())
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
        token_regex: Lazy::new(|| Regex::new(r"^[\{|\}]").unwrap())
    }
];

pub fn lex(mut code: &str, use_whitespace: bool) -> Vec<Token> {
    let mut tokens: Vec<Token> = Vec::new();
    while !code.is_empty() {
        let mut is_match = false;
        for s in &SYNTAX {
            if let Some(caps) = s.token_regex.captures(code) {
                is_match = true;
                code = code.strip_prefix(&caps[0]).unwrap_or(code);
                let mut vcaps: Vec<String> = Vec::new();
                let capsl = caps.len();
                let mut x = 0;
                while x < capsl {
                    vcaps.push(caps[x].to_string());
                    x+=1;
                }
                if (!use_whitespace && s.token_type!=TokenType::Whitespace) || use_whitespace {
                    tokens.push(Token {
                        token_type: s.token_type,
                        token_values: vcaps,
                    });
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