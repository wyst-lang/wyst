use std::{fs, path::Path, process::Command};
use std::str;

use crate::{
    transpiler::{State, Transpiler},
    utils::{Problem, ProblemCap, ProblemType},
};

pub fn transpile(input_file: String, output_file: String, is_main: bool) -> u8 {
    let mut ecode: u8 = 0;
    let mut trsp = Transpiler {
        ..Default::default()
    };
    if let Ok(file_contents) = fs::read_to_string(input_file) {
        let output_code = trsp.transpile_code(file_contents, 0);
        println!("{output_code}");
        if is_main {
            // output_code += "fn main() {";
            // output_code += vars
            //     .get_var(
            //         "main".to_string(),
            //         &mut trsp,
            //         State {
            //             line: 0,
            //             column: 0,
            //             file: None,
            //         },
            //     )
            //     .as_str();
            // output_code += "();}";
        }
        if let Ok(_) = fs::write(Path::new("build").join(output_file), output_code) {}
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
        .arg(exe_file)
        .output()
        .expect("Error compiling rust: ");
    if &output.stderr.len() == &0 {
        println!("Compiled successfully");
    } else {
        println!("{}", str::from_utf8(&output.stderr).expect("Error encoding stderr"));
    }
}
