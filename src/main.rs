mod compile;
mod dllmgr;
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
use variable::{Variable, VariableType, Variables};

use crate::lsp::run_lsp_server;

// Arguments for Wyst (short and long version)
#[derive(Parser)]
struct Args {
    //#[clap(short, long)]
    //rust: Option<String>,
    #[clap(short, long)]
    compile: Option<String>,

    #[clap(short, long)]
    dll: Option<String>,

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
            match args.compile {
                Some(ref exe_name) => {
                    let file_content = fs::read_to_string("main.wt").expect("Error reading file");
                    if Path::new("build").exists() {
                        fs::remove_dir_all("build").expect("err rm build");
                    }
                    fs::create_dir("build").expect("error making build");
                    let mut trsp = Transpiler::default();
                    let mut vars = Variables::new();
                    let mut transpiled_code = trsp.transpile(file_content, 0, &mut vars);
                    transpiled_code += "\nfn main() {";
                    transpiled_code += "std::process::exit(";
                    transpiled_code += vars.get_var("main".to_string(), &mut trsp).as_str();
                    transpiled_code += "())}";
                    for problem in &trsp.problems {
                        println!("{:?}: {}", problem.problem_type, problem.problem_msg)
                    }
                    if trsp.problems.len() > 0 {
                        return;
                    }
                    trsp.writer.write();

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
            match args.dll {
                Some(ref dll_path) => {
                    let file_content = fs::read_to_string("lib.wt").expect("Error reading file");
                    if Path::new("build").exists() {
                        fs::remove_dir_all("build").expect("err rm build");
                    }
                    fs::create_dir("build").expect("error making build");
                    let mut trsp = Transpiler::default();
                    let mut vars = Variables::new();
                    let transpiled_code = trsp.transpile(file_content, 0, &mut vars);
                    for problem in trsp.problems {
                        println!("{}", problem.problem_msg)
                    }
                    trsp.writer.write();
                    let mut dll_main = String::from(
                        "mod wslib;use wslib::*;\nfn call_fn(fn_name: &str, params: Vec<Param>)->i32{match fn_name {",
                    );
                    for (name, var) in vars.vars.clone() {
                        if var.vtype != VariableType::Func {
                            continue;
                        }
                        let mut dparams = String::new();
                        println!("{:?}", var.params.vars);
                        for i in 0..var.params.vars.len() {
                            dparams += format!("params.get({}).expect(\"Err_prms\"),", i).as_str();
                        }
                        dll_main += format!(
                            "\"{}\" => {}return {}({dparams});{}",
                            name, "{", var.rname, "}"
                        )
                        .as_str();
                    }
                    dll_main += "}}\nfn main(){}";
                    //dll_main = "fn main(){}".to_string();
                    compile::write_to_rust_file(&transpiled_code, "build/wslib.rs")
                        .expect("Error writing to temporary Rust file");
                    compile::write_to_rust_file(&dll_main, "build/main.rs")
                        .expect("Error writing to main dll");
                    std::env::set_current_dir("build").expect("setDir err: ");
                    compile::compile_to_executable("run").expect("Error compiling to executable");
                    dllmgr::write_dll(vars, "run".to_string(), dll_path.to_string());
                    // fs::remove_dir_all("build").expect("err rm build");
                }
                None => {}
            }
        }
    }
}
