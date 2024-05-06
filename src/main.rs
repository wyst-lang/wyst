mod transpiler;
mod parser;
mod lexer;
use std::fs;
use clap::Parser;


//Arguments for Wyst (short and long version)
#[derive(Parser)]
struct Args {
    #[clap(short, long)]
    file: String,
}

fn main() {
    let args = Args::parse();
    let file_content = fs::read_to_string(&args.file)
        .expect("Error reading file");

    let transpiled_code = transpiler::transpile(file_content, 0);
    println!("{}", transpiled_code);
}