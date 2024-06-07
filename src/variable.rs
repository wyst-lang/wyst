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
}
