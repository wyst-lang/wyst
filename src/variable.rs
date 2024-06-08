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
    pub params: Vec<[String; 2]>,
}

impl Variable {
    pub fn new_var(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Var,
            desc,
            state,
            params: vec![],
        }
    }

    pub fn new_namespace(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Namespace,
            desc,
            state,
            params: vec![],
        }
    }

    pub fn new_struct(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Struct,
            desc,
            state,
            params: vec![],
        }
    }

    pub fn new_keyword(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Keyword,
            desc,
            state,
            params: vec![],
        }
    }

    pub fn new_func(state: LexerState, desc: String) -> Variable {
        Variable {
            vtype: VariableType::Func,
            desc,
            state,
            params: vec![],
        }
    }
}
