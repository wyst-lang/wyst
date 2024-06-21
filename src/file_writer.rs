use std::fs;

use crate::{dll_reader::read_dll, transpiler::Transpiler, variable::Variables};

#[derive(Clone, Debug)]
pub struct Module {
    pub file_ws: String,
    pub file_rs: String,
    pub mod_rs: String,
    pub _code: String,
    pub code_rs: String,
}

#[derive(Clone, Debug)]
pub struct FileWriter {
    pub files: Vec<Module>,
    pub mod_num: u32,
    pub _projpath: String,
}

impl FileWriter {
    pub fn new(_projpath: String) -> FileWriter {
        FileWriter {
            files: Vec::new(),
            mod_num: 0,
            _projpath,
        }
    }
    pub fn write(&mut self) {
        for module in &self.files {
            fs::write(&module.file_rs, &module.code_rs).expect("WriteErr: can't write to file");
        }
    }
    /*Checks if a path has already been imported*/
    pub fn check(&mut self, filepath: String) -> Option<&Module> {
        for module in &self.files {
            if filepath == module.file_ws {
                return Some(module);
            }
        }
        return None;
    }
    pub fn expand(&mut self, writer: FileWriter) {
        for module in writer.files {
            self.files.push(module);
        }
    }
    pub fn add(&mut self, filepath: String, variables: &mut Variables) -> String {
        if filepath.ends_with(".dll") {
            read_dll(filepath.clone());
            return "xyz".to_string();
        }
        if let Some(module) = self.check(filepath.clone()) {
            return module.mod_rs.clone();
        }
        let mut trsp = Transpiler {
            writer: self.clone(),
            ..Default::default()
        };
        let mut vars = Variables::new();
        let code = fs::read_to_string(filepath.clone()).expect("failed to read");
        let code_rs = trsp.transpile(code.clone(), 0, &mut vars);
        let module = Module {
            file_ws: filepath,
            mod_rs: format!("mod_{}", self.mod_num),
            file_rs: format!("build/mod_{}.rs", self.mod_num),
            _code: code,
            code_rs,
        };
        self.files.push(module.clone());
        self.expand(trsp.writer);
        variables.expand(vars);
        return module.mod_rs;
    }
}
