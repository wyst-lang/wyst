mod transpiler;
mod parser;
mod lexer;
mod compile; 
use std::{fs, path::Path};
use clap::Parser;
use lexer::LexerState;
use transpiler::Options;

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
    if Path::new("build").exists() {
        fs::remove_dir_all("build").expect("err rm build");
    }
    fs::create_dir("build").expect("error making build");
    let transpiled_code = transpiler::transpile(file_content, 0, LexerState { line: 1, column: 0 }, &mut Options::default());

    match args.rust {
        Some(ref rust_file_name) => {
            // Write transpiled code to Rust file
            compile::write_to_rust_file(&transpiled_code, ("build/".to_string()+rust_file_name).as_str())
                .expect("Error writing to Rust file");
            println!("Transpiled code written to {}", rust_file_name);
        }
        None => {
            // No Rust file specified, print transpiled code
            if !args.compile.is_some() {
                println!("{}", transpiled_code);
            }
        }
    }

    match args.compile {
        Some(ref exe_name) => {
            compile::write_to_rust_file(&transpiled_code, "build/temp.rs")
                .expect("Error writing to temporary Rust file");
            compile::compile_to_executable("build/temp.rs", exe_name)
                .expect("Error compiling to executable");
            fs::remove_dir_all("build").expect("err rm build");
        }
        None => {}
    }
}
