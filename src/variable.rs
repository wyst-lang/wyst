use std::collections::HashMap;

use rand::Rng;
use serde::{Deserialize, Serialize};

use crate::{
    lexer::LexerState,
    lspcom::{Problem, ProblemType},
    transpiler::Transpiler,
};

#[derive(Clone, PartialEq, Debug, Serialize, Deserialize)]
pub enum VariableType {
    Var,
    Func,
    Keyword,
    Struct,
    Namespace,
    Enum,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Variable {
    pub vtype: VariableType,
    pub desc: String,
    pub state: LexerState,
    pub params: Variables,
    pub rname: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Variables {
    pub vars: HashMap<String, Variable>,
}

impl Variables {
    pub fn new() -> Variables {
        Variables {
            vars: HashMap::from([(
                "void".to_string(),
                Variable {
                    vtype: VariableType::Keyword,
                    desc: "".to_string(),
                    state: LexerState { line: 0, column: 0 },
                    params: Variables {
                        vars: HashMap::new(),
                    },
                    rname: "".to_string(),
                },
            )]),
        }
    }
    pub fn new_var(&mut self, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype: VariableType::Var,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }

    pub fn new_namespace(&mut self, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype: VariableType::Namespace,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }

    pub fn new_struct(&mut self, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype: VariableType::Struct,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }
    pub fn new_enum(&mut self, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype: VariableType::Enum,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }

    // pub fn new_keyword(&mut self, name: String, state: LexerState, desc: String) {
    //     self.vars.insert(
    //         name,
    //         Variable {
    //             vtype: VariableType::Keyword,
    //             desc,
    //             state,
    //             params: Variables {
    //                 vars: HashMap::new(),
    //             },
    //             rname: generate_varname(),
    //         },
    //     );
    // }

    pub fn new_func(&mut self, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype: VariableType::Func,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }
    pub fn add(&mut self, vtype: VariableType, name: String, state: LexerState, desc: String) {
        self.vars.insert(
            name,
            Variable {
                vtype,
                desc,
                state,
                params: Variables {
                    vars: HashMap::new(),
                },
                rname: generate_varname(),
            },
        );
    }
    pub fn get_var(&mut self, name: String, root: &mut Transpiler) -> String {
        if let Some(x) = self.get_mut(name.clone()) {
            return x.rname.clone();
        } else {
            root.problems.push(Problem {
                problem_type: ProblemType::VariableNotFound,
                problem_msg: format!("Variable '{}' doesn't exist", &name),
            });
            return name;
        }
    }
    pub fn iter_mut(
        &mut self,
    ) -> std::collections::hash_map::IterMut<'_, std::string::String, Variable> {
        self.vars.iter_mut()
    }
    pub fn get_mut(&mut self, name: String) -> Option<&mut Variable> {
        self.vars.get_mut(&name)
    }
    pub fn expand(&mut self, vars: Variables) {
        for (x, y) in vars.vars {
            self.vars.insert(x, y);
        }
    }
}

pub fn generate_varname() -> String {
    let mut rng = rand::thread_rng();
    let charset: &[u8] = b"ABCDEFGHIJKLMNOPQRSTUVWXYZ\
                          abcdefghijklmnopqrstuvwxyz\
                          0123456789";
    let mut result = String::with_capacity(19);
    result.push('_');
    result.push('0');
    result.push('x');
    for _ in 0..14 {
        let random_char = charset[rng.gen_range(0..charset.len())] as char;
        result.push(random_char);
    }
    result
}
