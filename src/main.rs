mod transpiler;
mod parser;
mod lexer;
mod compile; 
use std::fs;
use clap::Parser;
use lexer::LexerState;

// Arguments for Wyst (short and long version)
#[derive(Parser)]
struct Args {
    #[clap()]
    file: std::path::PathBuf,

    #[clap(short, long)]
    rust: Option<String>,

    #[clap(short, long)]
    compile: Option<String>,
}

fn main() {
    let args = Args::parse();
    let file_content = fs::read_to_string(&args.file)
        .expect("Error reading file");

    let transpiled_code = transpiler::transpile(file_content, 0, LexerState { line: 1, column: 0 });

    match args.rust {
        Some(ref rust_file_name) => {
            // Write transpiled code to Rust file
            compile::write_to_rust_file(&transpiled_code, rust_file_name)
                .expect("Error writing to Rust file");
            println!("Transpiled code written to {}", rust_file_name);
        }
        None => {
            // No Rust file specified, print transpiled code
            println!("{}", transpiled_code);
        }
    }

    match args.compile {
        Some(ref exe_name) => {
            // Compile transpiled Rust code into executable
            compile::write_to_rust_file(&transpiled_code, "temp.rs")
                .expect("Error writing to temporary Rust file");
            compile::compile_to_executable("temp.rs", exe_name)
                .expect("Error compiling to executable");
            fs::remove_file("temp.rs")
                .expect("Error removing temporary Rust file");
            println!("Compiled executable: {}", exe_name);
        }
        None => {}
    }
}
