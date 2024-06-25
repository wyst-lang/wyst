#[derive(Debug, PartialEq, Clone)]
enum Token {
    Plus,
    Minus,
    Star,
    Slash,
    Equals,
    Semicolon,
    LParen,
    RParen,
    Identifier(String),
    Number(i64),
    EOF,
}

fn tokenize(input: &str) -> Vec<Token> {
    let mut tokens = Vec::new();
    let chars: Vec<char> = input.chars().collect();
    let mut i = 0;

    while i < chars.len() {
        match chars[i] {
            ' ' | '\t' | '\n' | '\r' => i += 1,
            '+' => {
                tokens.push(Token::Plus);
                i += 1;
            }
            '-' => {
                tokens.push(Token::Minus);
                i += 1;
            }
            '*' => {
                tokens.push(Token::Star);
                i += 1;
            }
            '/' => {
                tokens.push(Token::Slash);
                i += 1;
            }
            '=' => {
                tokens.push(Token::Equals);
                i += 1;
            }
            ';' => {
                tokens.push(Token::Semicolon);
                i += 1;
            }
            '(' => {
                tokens.push(Token::LParen);
                i += 1;
            }
            ')' => {
                tokens.push(Token::RParen);
                i += 1;
            }
            '0'..='9' => {
                let mut num = 0;
                while i < chars.len() && chars[i].is_digit(10) {
                    num = num * 10 + chars[i].to_digit(10).unwrap() as i64;
                    i += 1;
                }
                tokens.push(Token::Number(num));
            }
            'a'..='z' | 'A'..='Z' => {
                let mut id = String::new();
                while i < chars.len() && (chars[i].is_alphanumeric() || chars[i] == '_') {
                    id.push(chars[i]);
                    i += 1;
                }
                match id.as_str() {
                    _ => tokens.push(Token::Identifier(id)),
                }
            }
            _ => panic!("Unexpected character: {}", chars[i]),
        }
    }

    tokens.push(Token::EOF);
    tokens
}

#[derive(Debug)]
enum ASTNode {
    Number(i64),
    Identifier(String),
    BinaryOp(Box<ASTNode>, Token, Box<ASTNode>),
    Assignment(String, Box<ASTNode>),
    Print(Box<ASTNode>),
    Let(String, Box<ASTNode>),
    Sequence(Vec<ASTNode>),
}

struct Parser {
    tokens: Vec<Token>,
    current_token: Token,
    pos: usize,
}

impl Parser {
    fn new(tokens: Vec<Token>) -> Self {
        Self {
            tokens: tokens.clone(),
            current_token: tokens[0].clone(),
            pos: 0,
        }
    }

    fn eat(&mut self, token: Token) {
        if self.current_token == token {
            self.pos += 1;
            self.current_token = self.tokens[self.pos].clone();
        } else {
            panic!("Unexpected token: {:?}", self.current_token);
        }
    }

    fn factor(&mut self) -> ASTNode {
        match self.current_token.clone() {
            Token::Number(value) => {
                self.eat(Token::Number(value));
                ASTNode::Number(value)
            }
            Token::Identifier(name) => {
                self.eat(Token::Identifier(name.clone()));
                ASTNode::Identifier(name)
            }
            Token::LParen => {
                self.eat(Token::LParen);
                let expr = self.expression();
                self.eat(Token::RParen);
                expr
            }
            _ => panic!("Unexpected token: {:?}", self.current_token),
        }
    }

    fn term(&mut self) -> ASTNode {
        let mut node = self.factor();

        while self.current_token == Token::Star || self.current_token == Token::Slash {
            let token = self.current_token.clone();
            self.eat(token.clone());
            node = ASTNode::BinaryOp(Box::new(node), token, Box::new(self.factor()));
        }

        node
    }

    fn expression(&mut self) -> ASTNode {
        let mut node = self.term();

        while self.current_token == Token::Plus || self.current_token == Token::Minus {
            let token = self.current_token.clone();
            self.eat(token.clone());
            node = ASTNode::BinaryOp(Box::new(node), token, Box::new(self.term()));
        }

        node
    }

    fn statement(&mut self) -> ASTNode {
        match self.current_token.clone() {
            Token::Identifier(name) => {
                self.eat(Token::Identifier(name.clone()));
                self.eat(Token::Equals);
                let expr = self.expression();
                self.eat(Token::Semicolon);
                ASTNode::Assignment(name, Box::new(expr))
            }
            _ => panic!("Unexpected token: {:?}", self.current_token),
        }
    }

    fn parse(&mut self) -> ASTNode {
        let mut statements = Vec::new();

        while self.current_token != Token::EOF {
            statements.push(self.statement());
        }

        ASTNode::Sequence(statements)
    }
}

use std::collections::HashMap;

struct Interpreter {
    symbol_table: HashMap<String, i64>,
}

impl Interpreter {
    fn new() -> Self {
        Self {
            symbol_table: HashMap::new(),
        }
    }

    fn evaluate(&mut self, node: &ASTNode) -> i64 {
        match node {
            ASTNode::Number(value) => *value,
            ASTNode::Identifier(name) => *self.symbol_table.get(name).expect("Undefined variable"),
            ASTNode::BinaryOp(left, token, right) => {
                let left_val = self.evaluate(left);
                let right_val = self.evaluate(right);
                match token {
                    Token::Plus => left_val + right_val,
                    Token::Minus => left_val - right_val,
                    Token::Star => left_val * right_val,
                    Token::Slash => left_val / right_val,
                    _ => panic!("Unexpected binary operator"),
                }
            }
            ASTNode::Assignment(name, expr) => {
                let value = self.evaluate(expr);
                self.symbol_table.insert(name.clone(), value);
                value
            }
            ASTNode::Let(name, expr) => {
                let value = self.evaluate(expr);
                self.symbol_table.insert(name.clone(), value);
                value
            }
            ASTNode::Print(expr) => {
                let value = self.evaluate(expr);
                println!("{}", value);
                value
            }
            ASTNode::Sequence(statements) => {
                for statement in statements {
                    self.evaluate(statement);
                }
                0
            }
        }
    }
}

fn main() {
    let input = "
        let x = 5;
        let y = x + 3;
        print y;
        x = y * 2;
        print x;
    ";
    let tokens = tokenize(input);
    println!("Tokens: {:?}", tokens);

    let mut parser = Parser::new(tokens);
    let ast = parser.parse();
    println!("AST: {:?}", ast);

    let mut interpreter = Interpreter::new();
    interpreter.evaluate(&ast);
}
