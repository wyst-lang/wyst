mod lexer;
use lexer::lex;
mod parser;
use parser::parse;

fn main() {
    let input = "int main() {}";
    println!("\n\n\n\n");
    // let input = "test->xy";
    let mut tokens = lex(input, false);
    let ast = *parse(&mut tokens);
    // for t in tokens {
    //     println!("{:?}", t);
    // }
    let mut _x = 1;
    for a in ast {
        println!("{:?}: {:?}\n", a.type_, a.values);
        _x+=1;
    }
}
