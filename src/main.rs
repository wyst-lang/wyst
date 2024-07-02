use utils::ProblemCap;
use variable::Variables;

mod parser;
mod transpiler;
mod utils;
mod variable;

fn main() {
    let mut vars = Variables::new();
    let mut trsp = transpiler::Transpiler::default();
    let out = trsp.transpile("void main() {x int x}".to_string(), 0, &mut vars);
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
    println!("{out}\n\n\n{:?}", vars);
}
