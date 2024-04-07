mod lexer;
use lexer::lex;
mod parser;
use parser::Parser;

fn main() {
    let input = "int main() {}";
    println!("\n\n\n\n");
    let tokens = lex(input, false);
    let mut ast = Parser::new(tokens.clone());
    while ast.tokens.len() > ast.index as usize {
        println!("{}\n\n", ast.next());
    }
}
