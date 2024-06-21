use std::{
    fs::File,
    io::{BufReader, Read},
};

use zip::{read::ZipFile, ZipArchive};

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
