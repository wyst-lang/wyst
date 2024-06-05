use lsp_types;
use once_cell::sync::Lazy;
use regex::Regex;
use std::io::{stdin, stdout, Read};

pub fn run_lsp_server() {
    let clpattern = Lazy::new(|| Regex::new(r"Content-Length: (\d+)\r\n").unwrap());
    let mut sin = stdin();
    let mut sout = stdout();

    println!("ok");

    loop {
        let mut input = String::new();
        let _ = sin.read_line(&mut input);
        if let Some(caps) = clpattern.captures(&input) {
            let x = 30;
            let mut buf = vec![0u8; x];
            sin.read_exact(&mut buf);
            
            println!("{}", &caps[1])
        }
    }
}