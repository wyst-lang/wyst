// use std::collections::HashMap;
// use std::fs;

use crate::parser::{new_vars, Variable};
use crate::transpiler::Transpiler;
use lspower::jsonrpc::Result;
use lspower::lsp::{self, *};
use lspower::{self, LspService, Server};
use lspower::{Client, LanguageServer};
use once_cell::sync::Lazy;
use rand::{thread_rng, Rng};
use std::collections::HashMap;
use std::vec;

#[derive(Debug)]
struct Backend {
    client: Client,
}

static DOCUMENTS: Lazy<HashMap<Url, String>> = Lazy::new(|| HashMap::default());

#[lspower::async_trait]
impl LanguageServer for Backend {
    async fn initialize(&self, _: InitializeParams) -> Result<InitializeResult> {
        // Ok(InitializeResult::default())
        Ok(InitializeResult {
            capabilities: ServerCapabilities {
                completion_provider: Some(CompletionOptions {
                    trigger_characters: vec!["::".to_string()].into(),
                    ..CompletionOptions::default()
                }),
                ..ServerCapabilities::default()
            },
            ..InitializeResult::default()
        })
    }

    async fn initialized(&self, _: InitializedParams) {
        self.client
            .log_message(MessageType::INFO, "Wyst LSP initialized!")
            .await;
    }

    async fn did_change(&self, params: lsp::DidChangeTextDocumentParams) {
        // DOCUMENTS.insert(params.text_document.uri, d);

    }

    async fn completion(
        &self,
        params: lsp::CompletionParams,
    ) -> lspower::jsonrpc::Result<Option<lsp::CompletionResponse>> {
        // let document_manager = &self.document;
        // let uri = params.text_document_position.text_document.uri.to_string();
        // let document_content = document_manager.get_document(&uri).unwrap_or_default(); // Handle missing document

        // let variables = get_completion(
        //     document_content.to_string(),
        //     params.text_document_position.position.line as usize + 1,
        //     params.text_document_position.position.character as usize,
        // );

        let mut completion_items: Vec<CompletionItem> = Vec::new();
        // for (name, _) in variables {
        // completion_items.push(CompletionItem::new_simple(name, "".to_string()));
        // }

        Ok(Some(CompletionResponse::Array(completion_items)))
    }

    async fn shutdown(&self) -> Result<()> {
        Ok(())
    }
}

#[tokio::main]
pub async fn run_lsp_server() {
    let stdin = tokio::io::stdin();
    let stdout = tokio::io::stdout();

    let (service, messages) = LspService::new(|client| Backend { client });
    Server::new(stdin, stdout)
        .interleave(messages)
        .serve(service)
        .await;
}

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
