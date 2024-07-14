use crate::transpiler::State;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum ProblemType {
    VariableNotFound,
    FileNotFound,
    HeaderSyntaxError,
    SyntaxError,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum ProblemCap {
    Error(Problem),
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Problem {
    pub problem_type: ProblemType,
    pub problem_msg: String,
    pub state: State,
}