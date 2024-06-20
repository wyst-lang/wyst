mod compile;
mod file_writer;
mod lexer;
mod lsp;
mod lspcom;
mod parser;
mod transpiler;
mod variable;
use clap::Parser;
use std::{fs, path::Path};
use transpiler::Transpiler;
use variable::Variables;

use crate::lsp::run_lsp_server;

// Arguments for Wyst (short and long version)
#[derive(Parser)]
struct Args {
    #[clap(short, long)]
    rust: Option<String>,

    #[clap(short, long)]
    compile: Option<String>,

    #[clap(long)]
    stdio: bool,
}

fn main() {
    let args = Args::parse();
    match args.stdio {
        true => {
            run_lsp_server();
        }
        false => {
            let file_content = fs::read_to_string("main.wt").expect("Error reading file");
            if Path::new("build").exists() {
                fs::remove_dir_all("build").expect("err rm build");
            }
            fs::create_dir("build").expect("error making build");
            let mut trsp = Transpiler::default();
            let mut vars = Variables::new();
            let mut transpiled_code = trsp.transpile(file_content, 0, &mut vars);
            transpiled_code += "\nfn main() {std::process::exit(";
            transpiled_code += vars.get_var("main".to_string(), &mut trsp).as_str();
            transpiled_code += "())}";
            for problem in trsp.problems {
                println!("{}", problem.problem_msg)
            }

            trsp.writer.write();

            match args.rust {
                Some(ref rust_file_name) => {
                    // Write transpiled code to Rust file
                    compile::write_to_rust_file(
                        &transpiled_code,
                        ("build/".to_string() + "main.rs").as_str(),
                    )
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
                    compile::write_to_rust_file(&transpiled_code, "build/main.rs")
                        .expect("Error writing to temporary Rust file");
                    std::env::set_current_dir("build").expect("setDir err: ");
                    compile::compile_to_executable(exe_name)
                        .expect("Error compiling to executable");
                    std::env::set_current_dir("..").expect("setDir0 err: ");
                    fs::rename(Path::new("build").join(exe_name).as_path(), exe_name)
                        .expect("RenameErrBuld: ");
                    fs::remove_dir_all("build").expect("err rm build");
                }
                None => {}
            }
        }
    }
}
