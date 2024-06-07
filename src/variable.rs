use crate::lexer::LexerState;

#[derive(Clone, PartialEq, Debug)]
pub enum VariableType {
    Var,
    Func,
    Keyword,
}

#[derive(Clone, Debug)]
pub struct Variable {
    pub vtype: VariableType,
    pub state: LexerState,
    pub params: Vec<[String; 2]>,
}

impl Variable {
    pub fn new_var(vtype: VariableType, state: LexerState) -> Variable {
        Variable {
            vtype,
            state,
            params: vec![],
        }
    }

    pub fn new_func(vtype: VariableType, state: LexerState) -> Variable {
        Variable {
            vtype,
            state,
            params: vec![],
        }
    }
}
