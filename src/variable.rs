use std::collections::HashMap;

use crate::lexer::LexerState;

#[derive(Clone, PartialEq, Debug)]
pub enum VariableType {
    Var,
    Func,
    Keyword,
    Struct,
    Namespace,
}

#[derive(Clone, Debug)]
pub struct Variable {
    pub vtype: VariableType,
    pub desc: String,
    pub state: LexerState,
    pub params: HashMap<String, Variable>,
}

impl Variable {
    pub fn new_var(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Var,
            desc,
            state,
            params: HashMap::new(),
        }
    }

    pub fn new_namespace(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Namespace,
            desc,
            state,
            params: HashMap::new(),
        }
    }

    pub fn new_struct(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Struct,
            desc,
            state,
            params: HashMap::new(),
        }
    }

    pub fn new_keyword(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Keyword,
            desc,
            state,
            params: HashMap::new(),
        }
    }

    pub fn new_func(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Func,
            desc,
            state,
            params: HashMap::new(),
        }
    }
}
