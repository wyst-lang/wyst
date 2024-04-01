mod lexer;
use lexer::lex;

fn main() {
    let input = "int test() {}";
    // let input = "test->xy";
    let tokens = lex(input, true);
    for token in tokens {
        println!("{:?} {}", token.token_type, token.token_value);
    }
}
