use lsp_types;
use once_cell::sync::Lazy;
use regex::Regex;
use std::io::{stdin, stdout, Read};

pub fn run_lsp_server() {
    let clpattern = Lazy::new(|| Regex::new(r"Content-Length: (\d+)\n").unwrap());
    let mut sin = stdin();
    let mut sout = stdout();

    println!("ok");

    loop {
        let mut input = String::new();
        let _ = sin.read_line(&mut input);
        if let Some(caps) = clpattern.captures(&input) {
            let conent_len = *(&caps[1].parse::<usize>().unwrap());
            let mut buf = vec![0u8; conent_len];
            sin.read_exact(&mut buf);
            let json_string = String::from_utf8(buf).expect("Our bytes should be valid utf8");
            
        }
    }
}