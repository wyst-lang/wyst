use crate::{
    transpiler::Transpiler,
    variable::{VariableType, Variables},
};
use lsp_types::{
    CompletionItem, CompletionItemKind, CompletionParams, CompletionResponse, InitializeResult,
};
use rand::{thread_rng, Rng};
use serde::{Deserialize, Serialize};

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

pub fn get_completion(input: String, line: usize, column: usize) -> Variables {
    let rand_id: u32 = thread_rng().gen();
    let rand_id: String = rand_id.to_string();
    let mut transpiler = Transpiler {
        peek: rand_id.clone(),
        ..Default::default()
    };
    let mut vars = Variables::new();
    transpiler.transpile(place_at(input, rand_id, line, column), 0, &mut vars);
    transpiler.matched_vars
}

pub mod request_methods {
    pub const INITIALIZE: &str = "initialize";
    pub const COMPLETION: &str = "textDocument/completion";
    pub const INITIALIZED: &str = "initialized";
    pub const SHUTDOWN: &str = "shutdown";
    pub const DID_CHANGE: &str = "textDocument/didChange";
}

pub trait LspServer {
    fn did_change(&mut self, _params: TextDocumentChangeParams) {}
    fn completion(&mut self, _params: CompletionParams) -> CompletionResponse {
        CompletionResponse::Array(vec![])
    }
    fn initialize(&mut self) -> InitializeResult {
        InitializeResult::default()
    }
}

#[derive(Debug, Eq, PartialEq, Clone, Default, Deserialize, Serialize)]
pub struct TextDocumentChangeParams {
    pub uri: String,
    pub text: String,
}

pub fn get_items(mut items: Variables, lname: String) -> Vec<CompletionItem> {
    let mut completion_items: Vec<CompletionItem> = Vec::new();
    for (name, var) in items.iter_mut() {
        let mut item = CompletionItem::new_simple(lname.clone() + &name, var.desc.clone());
        if var.vtype != VariableType::Func {
            completion_items.extend(
                get_items(var.params.clone(), name.to_owned() + "::")
                    .iter()
                    .cloned(),
            )
        } else {
            // item.documentation = Some(Documentation::MarkupContent(MarkupContent {
            //     kind: MarkupKind::Markdown,
            //     value: "Markup".to_string(),
            // }));
        }
        match var.vtype {
            VariableType::Func => {
                item.kind = Some(CompletionItemKind::FUNCTION);
            }
            VariableType::Var => {
                item.kind = Some(CompletionItemKind::VARIABLE);
            }
            VariableType::Keyword => {
                item.kind = Some(CompletionItemKind::KEYWORD);
            }
            VariableType::Struct => {
                item.kind = Some(CompletionItemKind::STRUCT);
            }
            VariableType::Namespace => {
                item.kind = Some(CompletionItemKind::MODULE);
            }
        }
        completion_items.push(item);
    }
    completion_items
}

#[derive(Clone, Debug)]
pub enum ProblemType {
    VariableNotFound,
}

#[derive(Clone, Debug)]
pub struct Problem {
    pub code: u32,
    pub problem_type: ProblemType,
    pub problem_msg: String,
}
