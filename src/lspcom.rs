use std::collections::HashMap;
use crate::transpiler::Transpiler;
use lsp_types::{CompletionResponse, InitializeResult};
use rand::{thread_rng, Rng};
use crate::parser::{new_vars, Variable};

pub fn place_at(input: String, in2: String, line_goal: usize, column_goal: usize) -> String {
    let mut line: usize = 1;
    let mut column: usize = 0;
    let mut out = String::new();
    for c in input.chars() {
        out += c.to_string().as_str();
        column += 1;
        if c == '\n' {
            line += 1;
            column = 0;
        } else if line == line_goal && column == column_goal {
            out += in2.as_str();
            column += 1
        }
    }
    out
}

pub fn get_completion(input: String, line: usize, column: usize) -> HashMap<String, Variable> {
    let rand_id: u32 = thread_rng().gen();
    let rand_id: String = rand_id.to_string();
    let mut transpiler = Transpiler {
        peek: rand_id.clone(),
        ..Default::default()
    };
    transpiler.transpile(place_at(input, rand_id, line, column), 0, new_vars());
    transpiler.matched_vars
}

pub mod request_methods {
    pub const INITIALIZE: &str = "initialize";
    pub const COMPLETION: &str = "textDocument/completion";
    pub const INITIALIZED: &str = "initialized";
    pub const SHUTDOWN: &str = "shutdown";
}

pub trait LspServer {
    fn completion(&mut self) -> CompletionResponse { CompletionResponse::Array(vec![]) }
    fn initialize(&mut self) -> InitializeResult { InitializeResult::default() }
}