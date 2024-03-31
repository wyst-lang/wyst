// Define token types
#[derive(Debug, PartialEq)]
pub enum Token {
    Number(i32),
    Plus,
    Minus,
    Multiply,
    Divide,
    LParen,
    RParen,
    EOF,
}

// Define Lexer struct
pub struct Lexer<'a> {
    input: &'a str,
    position: usize,
}

impl<'a> Lexer<'a> {
    // Constructor
    pub fn new(input: &'a str) -> Self {
        Lexer { input, position: 0 }
    }

    // Advance position in input
    fn advance(&mut self) {
        self.position += 1;
    }

    // Get current character without advancing position
    fn current_char(&self) -> Option<char> {
        self.input.chars().nth(self.position)
    }

    // Lexical analysis
    pub fn next_token(&mut self) -> Token {
        // Skip whitespace
        while let Some(c) = self.current_char() {
            if c.is_whitespace() {
                self.advance();
            } else {
                break;
            }
        }

        // Check for end of input
        if let None = self.current_char() {
            return Token::EOF;
        }

        // Match characters to tokens
        match self.current_char().unwrap() {
            '+' => {
                self.advance();
                Token::Plus
            }
            '-' => {
                self.advance();
                Token::Minus
            }
            '*' => {
                self.advance();
                Token::Multiply
            }
            '/' => {
                self.advance();
                Token::Divide
            }
            '(' => {
                self.advance();
                Token::LParen
            }
            ')' => {
                self.advance();
                Token::RParen
            }
            // Match numbers
            digit if digit.is_digit(10) => {
                let mut num_str = String::new();
                while let Some(digit) = self.current_char() {
                    if digit.is_digit(10) {
                        num_str.push(digit);
                        self.advance();
                    } else {
                        break;
                    }
                }
                Token::Number(num_str.parse().unwrap())
            }
            // Unknown character
            _ => {
                panic!("Invalid character: {}", self.current_char().unwrap());
            }
        }
    }
}
