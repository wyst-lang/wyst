use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum ProblemType {
    VariableNotFound,
    FileNotFound,
    HeaderSyntaxError,
    SyntaxError,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Problem {
    //pub code: u32,
    pub problem_type: ProblemType,
    pub problem_msg: String,
}
