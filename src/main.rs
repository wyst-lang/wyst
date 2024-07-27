use std::str;
use std::thread::sleep;
use std::time::Duration;
use std::{fs, io::Read, process};
mod compiler;
mod transpiler;
mod utils;

use clap::{App, Arg};
use compiler::{compile_rust, transpile_file};

fn main() {
    let matches = App::new("Wyst")
        .version("0.51")
        .author("Orus")
        .about("The wyst compiler")
        .arg(
            Arg::with_name("file")
                .required_unless_one(&["stdio", "hex", "logo"])
                .help("Sets the input file to use"),
        )
        .arg(
            Arg::with_name("output")
                .short('o')
                .long("output")
                .takes_value(true)
                .help("Specifies the output binary file name"),
        )
        .arg(
            Arg::with_name("build")
                .short('b')
                .long("build")
                .help("Puts the rust files in the build directory"),
        )
        .arg(
            Arg::with_name("stdio")
                .short('s')
                .long("stdio")
                .help("LSP Mode"),
        )
        .arg(
            Arg::with_name("hex")
                .short('h')
                .long("hex")
                .takes_value(true)
                .help("Turn an identifier into hex"),
        )
        .arg(Arg::with_name("logo").long("logo").hide(true))
        .get_matches();
    if matches.is_present("logo") {
        println!("{}", utils::LOGO);
        let mut byte = [4; 4];
        std::io::stdin().read_exact(&mut byte).unwrap();
        let egg = str::from_utf8(&byte).unwrap().to_string();
        let mut counter: u8 = 30;
        if egg == "logo" {
            loop {
                if counter == 35 {
                    counter = 30;
                }
                println!("\x1B[2J\x1B[{}m{}", counter, utils::LOGO);
                sleep(Duration::from_millis(400));
                counter += 1;
            }
        }
    } else if matches.is_present("stdio") {
        println!("stdio mode!");
    } else if matches.is_present("hex") {
        let ident = matches.value_of("hex").unwrap().to_string();
        println!("_0x{}", hex::encode(ident));
    } else {
        let file = matches.value_of("file").unwrap().to_string();
        let output = matches.value_of("output").unwrap_or("out").to_string();
        if let Ok(_) = fs::create_dir("build") {}
        let ecode = transpile_file(file, "main.rs".to_string(), true);
        if ecode == 0 {
            compile_rust(output);
        }
        if !matches.is_present("build") {
            fs::remove_dir_all("build").expect("Error removing build");
        }
        process::exit(ecode as i32);
    }
}
