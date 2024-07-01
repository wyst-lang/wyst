use variable::Variables;

mod parser;
mod transpiler;
mod utils;
mod variable;

fn main() {
    let mut vars = Variables::new();
    let mut trsp = transpiler::Transpiler::default();
    let out = trsp.transpile("void main() {int x}".to_string(), 0, &mut vars);
    for p in trsp.problems {
        println!("{:?}: {}", p.problem_type, p.problem_msg)
    }
    println!("{out}\n\n\n{:?}", vars);
}
