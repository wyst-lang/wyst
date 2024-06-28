use variable::Variables;

mod parser;
mod transpiler;
mod utils;
mod variable;

fn main() {
    let mut vars = Variables::new();
    let out = transpiler::Transpiler::default().transpile(r#"int x"#.to_string(), 0, &mut vars);
    println!("{out}");
}
