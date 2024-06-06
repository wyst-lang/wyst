use once_cell::sync::Lazy;
use regex::Regex;
use serde::{Deserialize, Serialize};
use serde_json::{self, Value};

use std::{
    fs::{self, File},
    io::{stdin, Read, Write},
    path::Path,
};

pub fn run_method<F>(func: F) {}

pub fn run_lsp_server() {
    if Path::new("/Users/klestiselimaj/work/wyst/log.txt").exists() {
        fs::remove_file("/Users/klestiselimaj/work/wyst/log.txt");
    }
    let mut file = File::create("/Users/klestiselimaj/work/wyst/log.txt").expect("err_log");
    let clpattern = Lazy::new(|| Regex::new(r"^Content-Length: (\d+)").unwrap());
    let mut reader = stdin();
    // let mut sout = stdout();

    writeln!(file, "It's running!");

    loop {
        let mut input = String::new();
        let _ = reader.read_line(&mut input);
        writeln!(file, "{}", &input);
        if let Some(caps) = clpattern.captures(&input) {
            let conent_len = *(&caps[1].parse::<usize>().unwrap()) + 2;
            writeln!(file, "ok {}", conent_len);
            let mut buf = vec![0u8; conent_len];
            let _ = reader.read_exact(&mut buf);
            let json_string = String::from_utf8(buf).expect("Our bytes should be valid utf8");
            writeln!(file, "{}", &json_string);
            let client_json: Value = serde_json::from_str(&json_string).expect("err_json");
            writeln!(file, "bruh {}", &client_json["method"].as_str().unwrap());
        }
    }
}
