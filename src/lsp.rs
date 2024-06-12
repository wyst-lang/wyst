use crate::{
    lspcom::{get_completion, get_items, request_methods, LspServer, TextDocumentChangeParams},
    variable::Variables,
};
use lsp_types::*;
use once_cell::sync::Lazy;
use regex::Regex;
use serde_json::{self, json, Value};
use std::{
    collections::HashMap,
    fs,
    io::{stdin, stdout, Read, Write},
    path::Path,
};

#[derive(Debug)]
pub struct Server {
    documents: HashMap<String, String>,
}

impl LspServer for Server {
    fn initialize(&mut self) -> InitializeResult {
        InitializeResult {
            capabilities: ServerCapabilities {
                completion_provider: Some(CompletionOptions {
                    trigger_characters: Some(vec![".".to_string()]),
                    ..Default::default()
                }),
                ..Default::default()
            },
            ..Default::default()
        }
    }
    fn completion(&mut self, params: CompletionParams) -> CompletionResponse {
        let text = self
            .documents
            .get(params.text_document_position.text_document.uri.as_str())
            .expect("err_textdoc");
        let items: Variables = get_completion(
            text.clone(),
            params.text_document_position.position.line as usize + 1,
            params.text_document_position.position.character as usize,
        );
        CompletionResponse::Array(get_items(items, "".to_string()))
    }
    fn did_change(&mut self, params: TextDocumentChangeParams) {
        self.documents.insert(params.uri, params.text);
    }
}

pub fn run_lsp_server() {
    if Path::new("/home/leo/work/wyst/log.txt").exists() {
        fs::remove_file("/home/leo/work/wyst/log.txt").unwrap();
    }
    let clpattern = Lazy::new(|| Regex::new(r"^Content-Length: (\d+)").unwrap());
    let mut reader = stdin();
    let mut server = Server {
        documents: HashMap::new(),
    };
    loop {
        let mut input = String::new();
        if reader.read_line(&mut input).is_err() {
            continue;
        }
        let stdout = stdout();
        let mut handle = stdout.lock();

        if let Some(caps) = clpattern.captures(&input) {
            let content_len = caps[1].parse::<usize>().unwrap() + 2;
            let mut buf = vec![0u8; content_len];
            if reader.read_exact(&mut buf).is_err() {
                continue;
            }
            let json_string = String::from_utf8(buf).expect("Our bytes should be valid utf8");
            let client_json: Value = serde_json::from_str(&json_string).expect("err_json");
            let response = match client_json["method"].as_str().unwrap() {
                request_methods::INITIALIZE => {
                    serde_json::to_string(&json!({
                        "jsonrpc": "2.0",
                        "id": client_json["id"].as_u64().unwrap(),
                        "result": server.initialize()
                    }))
                    .unwrap()
                }
                request_methods::COMPLETION => serde_json::to_string(&json!({
                    "jsonrpc": "2.0",
                    "id": client_json["id"].as_u64().unwrap(),
                    "result": server.completion(serde_json::from_value(serde_json::to_value(client_json["params"].as_object()).expect("err_pars2")).unwrap())
                }))
                .unwrap(),
                request_methods::DID_CHANGE =>{
                    server.did_change(serde_json::from_value(serde_json::to_value(client_json["params"].as_object()).expect("err_pars2")).expect("err_pars3"));
                    "None".to_string()},
                request_methods::INITIALIZED => "None".to_string(),
                request_methods::SHUTDOWN => {
                    return;
                }
                _ => {
                    continue;
                }
            };
            if response != "None" {
                let header = format!(
                    "Content-Length: {}\r\n\r\n{}",
                    response.trim().len(),
                    response.trim()
                );
                handle
                    .write_all(header.as_bytes())
                    .expect("err_write_stdin");
                handle.flush().expect("err_flush_stdin");
            }
        } else {
        }
    }
}
