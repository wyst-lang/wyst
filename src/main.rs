mod lexer;
use lexer::lex;

fn main() {
    // let input = "-3.232313 + 4 * (10 - 2)";
    let input = "test";
    let mut tokens = lex(input);
}
