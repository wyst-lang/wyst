use regex::Regex;
use std::fs;
use std::io::Write;
use std::{
    fs::File,
    io::{BufReader, Read},
};
use zip::{write::SimpleFileOptions, ZipArchive, ZipWriter};

use crate::variable::Variables;

pub fn read_dll(dll_path: String) -> Variables {
    let mut vars = Variables::new();
    let dll_file = File::open(dll_path).expect("err_dll_read");
    let dll_reader = BufReader::new(dll_file);
    let mut dll_zip = ZipArchive::new(dll_reader).expect("err_archive");
    for i in 0..dll_zip.len() {
        let mut buf = Vec::new();
        let mut file = dll_zip.by_index(i).expect("err_dll_idx");
        let filename = file.name();
        if filename == "map.json" && file.is_file() {
            file.read_to_end(&mut buf).expect("err_read_zip_file");
            println!(
                "Filename: {}\nFileContents: {}",
                file.name(),
                String::from_utf8_lossy(&buf)
            );
        }
    }
    vars
}

pub fn write_dll(vars: Variables, exe_path: String, dll_path: String) {
    let dll_file = File::create(&dll_path).expect("Err_DllCREATE");
    let mut dll = ZipWriter::new(dll_file);
    dll.start_file("run", SimpleFileOptions::default())
        .expect("ERR_STARTFILE_RUN");
    let buf_vec = fs::read(exe_path).expect("ERR_READ_EXE");
    let buf_u8: &[u8] = buf_vec.as_slice();
    dll.write_all(buf_u8).expect("ERR_DLL_EXEWRITE");
    dll.start_file("map.json", SimpleFileOptions::default())
        .expect("ERR_STARTFILE_MAP");
    let json_data = serde_json::to_string(&vars.vars).expect("ERR_MAP_JSON");
    dll.write_all(json_data.as_bytes())
        .expect("ERR_DLL_JSONWRITE");
    println!("{:?}", serde_json::to_string(&vars.vars));
}

pub struct HeaderConfig {
    pub file_path: String,
    pub map_path: String,
}

impl HeaderConfig {
    // Function to parse the input string
    pub fn new(input: String) -> Option<Self> {
        let re = Regex::new(r#"FILE_PATH\s*=\s*"([^"]*)""#).expect("ERR_REGEX_HDR");
        let re1 = Regex::new(r#"MAP_PATH\s*=\s*"([^"]*)""#).expect("ERR_REGEX_HDR1");
        let mut fp: String = String::new();
        let mut mp: String = String::new();
        if let Some(caps) = re.captures(input.as_str()) {
            fp = caps.get(1)?.as_str().to_string();
        } else {
            return None;
        }
        if let Some(caps) = re1.captures(input.as_str()) {
            mp = caps.get(1)?.as_str().to_string();
        } else {
            return None;
        }
        return Some(HeaderConfig {
            file_path: fp,
            map_path: mp,
        });
    }
}
