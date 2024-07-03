use crate::{
    parser::{self, Ast},
    utils::ProblemCap,
    variable::Variables,
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
    pub peek: String,
    // pub libmgr: LibManager,
    pub problems: Vec<ProblemCap>,
    pub state: State,
}

impl Default for Transpiler {
    fn default() -> Self {
        Transpiler {
            matched_vars: Variables::new(),
            peek: String::new(),
            // libmgr: LibManager::new(".".to_string()),
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
            }
        }
        if indent > 0 {
            res += "\n";
            res += " ".repeat((indent - 1) * 2).as_str();
        }
        res
    }
}
