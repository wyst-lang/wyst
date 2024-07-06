use crate::{
    parser::{self, extract_ast, Ast},
    utils::{ProblemCap, VariableType, Variables},
};
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct State {
    pub line: u32,
    pub column: u32,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Transpiler {
    pub matched_vars: Variables,
    pub problems: Vec<ProblemCap>,
    pub state: State,
    pub inject: Inject,
}
#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Inject {
    pub state: State,
    pub inject: bool,
}

impl Default for Inject {
    fn default() -> Self {
        Inject {
            inject: false,
            state: State { line: 0, column: 0 },
        }
    }
}

impl Default for Transpiler {
    fn default() -> Self {
        Transpiler {
            matched_vars: Variables::new(),
            inject: Inject::default(),
            problems: Vec::new(),
            state: State { line: 1, column: 0 },
        }
    }
}

impl Transpiler {
    pub fn transpile(&mut self, code: String, indent: usize, vars: &mut Variables) -> String {
        let ast = parser::parse(code, vars, &mut self.state.clone(), self);
        let mut res = String::new();
        if indent > 0 {
            res += " ".repeat(indent * 2).as_str();
        }
        for a in ast {
            let astval = extract_ast(a.clone());
            if self.inject.inject {
                for val in astval {
                    if val.contains("list_vx") {
                        for (name, var) in vars.clone().iter_mut() {
                            if self.inject.state.line > var.state.line
                                || var.vtype != VariableType::Var
                            {
                                self.matched_vars.vars.insert(name.clone(), var.clone());
                            }
                        }
                    }
                }
            }
            match a {
                Ast::Variable(var_type, var_name) => {
                    res += format!("let mut {}: {}", var_name, var_type).as_str();
                }
                Ast::Function(var_type, var_name, round, curly) => {
                    res += format!(
                        "fn {}({}) -> {} {}\n{}{}\n",
                        var_name,
                        round,
                        var_type,
                        "{",
                        self.transpile(curly, indent + 1, &mut vars.clone()),
                        "}"
                    )
                    .as_str();
                }
                Ast::Struct(name, curly) => {
                    res += format!(
                        "struct {} {}\n{}{}\n",
                        name,
                        "{",
                        self.transpile(curly, indent + 1, &mut vars.clone()),
                        "}"
                    )
                    .as_str()
                }
                Ast::Rust(rs) => {
                    res += "{";
                    res += rs.as_str();
                    res += "}";
                }
                Ast::Single(_) => {}
            }
        }
        if indent > 0 {
            res += "\n";
            res += " ".repeat((indent - 1) * 2).as_str();
        }
        res
    }
}
