mod transpiler;
mod parser;
mod lexer;
use std::fs;


fn main() {
    let contents = fs::read_to_string("code.ws").expect("cannot read for some reason");
    let result = transpiler::transpile(contents, 0);
    println!("{result}")
}