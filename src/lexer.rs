use regex::Regex;
use once_cell::sync::Lazy;

// Define token types
#[derive(Debug, PartialEq)]
pub enum Token {
    Number,
    Dot,
    Plus,
    Minus,
    Multiply,
    Divide,
    LParen,
    RParen,
    EOF,
}


pub struct Node<'a> {
    token_type: Token,
    token_regex: Lazy<Regex>,
    token_value: &'a str
}

const SYNTAX: [Node; 1] = [
    Node {
        token_type: Token::Dot,
        token_regex: Lazy::new(|| Regex::new(r"^test").unwrap()),
        token_value: ""
    }
];

pub fn lex(code: &str) -> Vec<Node> {
    let mut tokens: Vec<Node> = Vec::new();
    while (code!="") {
        for s in SYNTAX {
            if (s.token_regex.is_match(code)) {
                println!("It works!");
                tokens.push(Node {
                    token_type: s.token_type,
                    token_regex: s.token_regex,
                    token_value: s.token_regex.captures(wisp)
                });
            }
        }
    } {
        
    }
    tokens
}

// Define Lexer struct
// pub struct Lexer<'a> {
//     input: &'a str,
//     position: usize,
// }

// impl<'a> Lexer<'a> {
//     // Constructor
//     pub fn new(input: &'a str) -> Self {
//         // Lexer { input, position: 0 }
//     }

//     // Advance position in input
//     fn advance(&mut self) {
//         self.position += 1;
//     }

//     // Get current character without advancing position
//     fn current_char(&self) -> Option<char> {
//         self.input.chars().nth(self.position)
//     }

//     // Lexical analysis
//     pub fn next_token(&mut self) -> Token {
//         // Skip whitespace
//         while let Some(c) = self.current_char() {
//             if c.is_whitespace() {
//                 self.advance();
//             } else {
//                 break;
//             }
//         }

//         // Check for end of input
//         if let None = self.current_char() {
//             return Token::EOF;
//         }

//         // Match characters to tokens
//         match self.current_char().unwrap() {
//             '+' => {
//                 self.advance();
//                 Token::Plus
//             }
//             '-' => {
//                 self.advance();
//                 Token::Minus
//             }
//             '*' => {
//                 self.advance();
//                 Token::Multiply
//             }
//             '/' => {
//                 self.advance();
//                 Token::Divide
//             }
//             '(' => {
//                 self.advance();
//                 Token::LParen
//             }
//             ')' => {
//                 self.advance();
//                 Token::RParen
//             }
//             '.' => {
//                 self.advance();
//                 Token::Dot
//             }
//             // Match numbers
//             digit if digit.is_digit(10) => {
//                 let mut num_str = String::new();
//                 while let Some(digit) = self.current_char() {
//                     if digit.is_digit(10) {
//                         num_str.push(digit);
//                         self.advance();
//                     } else {
//                         break;
//                     }
//                 }
//                 Token::Number(num_str.parse().unwrap())
//             }
//             // Unknown character
//             _ => {
//                 panic!("Invalid character: {}", self.current_char().unwrap());
//             }
//         }
//     }
// }
