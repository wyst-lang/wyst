// use std::collections::HashMap;
// use std::fs;

use std::fs;

use lspower;
use lspower::jsonrpc::Result;
use lspower::lsp::{self, *};
use lspower::{Client, LanguageServer};

use crate::transpiler::Transpiler;

#[derive(Debug)]
struct Backend {
    client: Client,
}

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

    async fn completion(
        &self,
        _params: lsp::CompletionParams,
    ) -> lspower::jsonrpc::Result<Option<lsp::CompletionResponse>> {
        Ok(Some(CompletionResponse::Array(vec![
            CompletionItem::new_simple("mylabel".to_string(), "mydetail".to_string()),
        ])))
    }

    async fn shutdown(&self) -> Result<()> {
        Ok(())
    }
}

// pub async fn run_lsp_server() {
//     let stdin = tokio::io::stdin();
//     let stdout = tokio::io::stdout();

//     let (service, messages) = LspService::new(|client| Backend { client });
//     Server::new(stdin, stdout)
//         .interleave(messages)
//         .serve(service)
//         .await;
// }

#[tokio::main]
pub async fn run_lsp_server() {
    let mut transpiler = Transpiler {
        var_match: String::from("x"),
        ..Default::default()
    };
    let file_content = fs::read_to_string("main.wt").expect("Error reading file");
    transpiler.transpile(file_content, 0);
}
