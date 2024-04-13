mod transpiler;
mod parser;
mod lexer;
use std::fs;


fn main() {
    let code_file = "code.ws";
    let contents = fs::read_to_string(code_file).expect("cannot read for some reason");
    let result = transpiler::transpile(contents, 0);
    println!("{result}")
}