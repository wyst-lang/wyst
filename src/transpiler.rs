use crate::{
    parser::{self, Ast},
    utils::Problem,
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
    pub problems: Vec<Problem>,
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
    pub fn transpile(&mut self, code: String, indent: u32, vars: &mut Variables) -> String {
        let ast = parser::parse(code, vars, self);
        let mut res = String::new();
        for a in ast {
            match a {
                Ast::Variable(var_type, var_name) => {
                    res += format!("let mut {}: {}", var_name, var_type).as_str();
                }
            }
        }
        res
    }
}
