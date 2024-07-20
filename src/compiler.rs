use std::str;
use std::{fs, path::Path, process::Command};

use crate::{
    transpiler::{State, Transpiler},
    utils::{Problem, ProblemCap, ProblemType},
};

pub fn transpile_file(input_file: String, output_file: String, is_main: bool) -> u8 {
    let mut ecode: u8 = 0;
    let mut trsp = Transpiler {
        ..Default::default()
    };
    if let Ok(file_contents) = fs::read_to_string(input_file) {
        let mut module_code = String::new();
        let mut output_code = String::new();
        match trsp.transpile_code(file_contents, 0) {
            Ok(rcode) => output_code += rcode.as_str(),
            Err(e) => {
                ecode = 1;
                println!("{e}")
            }
        }
        if is_main {
            for module in trsp.clone().mod_manager.modules {
                module_code += "mod ";
                module_code += module.rs_mod.as_str();
                module_code += ";\n";
            }
        }
        output_code = format!(
            "{module_code}\n{output_code}\nfn main() {}{}{}",
            "{_0x",
            hex::encode("main"),
            "();}"
        );
        if let Ok(_) = fs::write(Path::new("build").join(output_file), output_code) {
            trsp.problems.append(&mut trsp.mod_manager.write());
        }
    } else {
        trsp.problems.push(ProblemCap::Error(Problem {
            problem_msg: "File not found".to_string(),
            problem_type: ProblemType::FileNotFound,
            state: State {
                line: 0,
                column: 0,
                file: None,
            },
        }))
    }
    for p in trsp.problems {
        match p {
            ProblemCap::Error(problem) => {
                ecode = 1;
                if problem.state.file == None {
                    println!(
                        "\x1b[31;1m{:?}\x1b[0;1m: {}",
                        problem.problem_type, problem.problem_msg,
                    )
                } else {
                    println!(
                        "\x1b[31;1mError\x1b[0;1m: {}\x1b[0m\n{}:{}: {:?}",
                        problem.problem_msg,
                        problem.state.line,
                        problem.state.column,
                        problem.problem_type,
                    )
                }
            }
        }
    }
    ecode
}

pub fn compile_rust(exe_file: String) {
    let output = Command::new("rustc")
        .arg(Path::new("build").join("main.rs"))
        .arg("-o")
        .arg(exe_file.clone())
        .output()
        .expect("Error compiling rust: ");
    if &output.stderr.len() == &0 {
        println!("Compiled successfully: \x1b[34;4m{}", exe_file.clone());
    } else {
        println!(
            "{}",
            str::from_utf8(&output.stderr).expect("Error encoding stderr")
        );
    }
}
