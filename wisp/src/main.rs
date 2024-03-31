mod lexer;  // Import the lexer module

use lexer::lexer::Lexer;  // Import the Lexer struct
use lexer::lexer::Token;  // Import the Token enum

fn main() {
    let input = "3 + 4 * (10 - 2)";
    let mut lexer = Lexer::new(input);  

    loop {
        let token = lexer.next_token();
        println!("{:?}", token);
        if token == Token::EOF {
            break;
        }
    }
}
