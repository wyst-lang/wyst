use utils::{ProblemCap, Variables};

mod parser;
mod transpiler;
mod utils;

fn main() {
    let mut vars = Variables::new();
    let mut trsp = transpiler::Transpiler::default();
    let out = trsp.transpile("rust{let mut x;}".to_string(), 0, &mut vars);
    for p in trsp.problems {
        match p {
            ProblemCap::Error(problem) => {
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
    println!("{out}\n\n\n");
}
