use crate::lspcom::{request_methods, LspServer};
use lsp_types::*;
use once_cell::sync::Lazy;
use regex::Regex;
use serde_json::{self, json, Value};
use std::{
    collections::HashMap,
    fs::{self, File},
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
    fn completion(&mut self) -> CompletionResponse {
        CompletionResponse::Array(vec![CompletionItem::new_simple(
            "mytest".to_string(),
            "mydetail".to_string(),
        )])
    }
}

pub fn run_lsp_server() {
    if Path::new("/home/leo/work/wyst/log.txt").exists() {
        fs::remove_file("/home/leo/work/wyst/log.txt").unwrap();
    }
    let mut file = File::create("/home/leo/work/wyst/log.txt").expect("err_log");
    let clpattern = Lazy::new(|| Regex::new(r"^Content-Length: (\d+)").unwrap());
    let mut reader = stdin();
    let mut server = Server {
        documents: HashMap::new(),
    };
    writeln!(file, "It's running!").unwrap();

    loop {
        writeln!(file, "waiting").unwrap();
        let mut input = String::new();
        if reader.read_line(&mut input).is_err() {
            writeln!(file, "Error reading line").unwrap();
            continue;
        }
        writeln!(file, "{}", input).unwrap();
        let stdout = stdout();
        let mut handle = stdout.lock();

        if let Some(caps) = clpattern.captures(&input) {
            let content_len = caps[1].parse::<usize>().unwrap() + 2;
            writeln!(file, "ok {}", content_len).unwrap();
            let mut buf = vec![0u8; content_len];
            if reader.read_exact(&mut buf).is_err() {
                writeln!(file, "Error reading content").unwrap();
                continue;
            }
            let json_string = String::from_utf8(buf).expect("Our bytes should be valid utf8");
            writeln!(file, "{}", &json_string).unwrap();
            let client_json: Value = serde_json::from_str(&json_string).expect("err_json");
            let response = match client_json["method"].as_str().unwrap() {
                request_methods::INITIALIZE => {
                    writeln!(file, "test").unwrap();
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
                    "result": server.completion()
                }))
                .unwrap(),
                request_methods::INITIALIZED => "None".to_string(),
                request_methods::SHUTDOWN => {
                    return;
                }
                _ => {
                    writeln!(
                        file,
                        "other request {}",
                        client_json["method"].as_str().unwrap()
                    )
                    .unwrap();
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
                writeln!(file, "{}", header).unwrap();
            }
        } else {
            writeln!(file, "not match").unwrap();
        }
    }
}
