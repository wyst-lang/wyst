use std::{collections::HashMap, fs, path::Path};

use dirs::home_dir;

use crate::{
    dllmgr::{read_dll, HeaderConfig},
    lspcom::{Problem, ProblemType},
    transpiler::Transpiler,
    variable::{Variable, Variables},
};

fn join_directories(dir1: &str, dir2: &str) -> String {
    let mut main_seperator = '/';
    if cfg!(windows) {
        main_seperator = '\\';
    }
    let mut result = dir1.to_string();

    if !result.ends_with(main_seperator) {
        result.push(main_seperator);
    }

    let dir2_trimmed = if dir2.starts_with(main_seperator) {
        &dir2[1..]
    } else {
        dir2
    };

    result.push_str(dir2_trimmed);
    result
}

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
    pub fn add(
        &mut self,
        filepath: String,
        variables: &mut Variables,
        global: bool,
    ) -> Result<String, Problem> {
        let mut filepath = filepath;
        if global {
            let homedir_ = home_dir().expect("Err_HOMEDIR_NOTFOUND");
            let homedir = homedir_.to_str().expect("ERR_HOMEDIR_STR");
            let gdir = join_directories(&homedir, &join_directories("wyst", "lib"));
            let lpath = join_directories("lib", &filepath);
            let gpath = join_directories(&gdir, &filepath);
            if Path::new(&lpath).exists() {
                filepath = lpath;
            } else if Path::new(&gpath).exists() {
                filepath = gpath;
            } else {
                return Err(Problem {
                    problem_type: ProblemType::FileNotFound,
                    problem_msg: format!("failed to import {}: Not found", filepath),
                });
            }
        }
        if !Path::new(&filepath).exists() {
            return Err(Problem {
                problem_type: ProblemType::FileNotFound,
                problem_msg: format!("failed to import {}: Not found", filepath),
            });
        }
        if let Some(module) = self.check(filepath.clone()) {
            return Ok(module.mod_rs.clone());
        }
        match filepath.rsplit_once('.').expect("Err_SPLIT").1 {
            "wt" => {
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
                self.mod_num += 1;
                Ok(module.mod_rs)
            }
            "h" => {
                let hdr_code = fs::read_to_string(filepath.clone()).expect("failed to read");
                let mut parent = String::new();
                match Path::new(&filepath).parent() {
                    Some(p) => parent = format!("{}", p.display()),
                    None => {}
                };
                match HeaderConfig::new(hdr_code) {
                    Some(hdrcnf) => {
                        let file_path = Path::new(&parent).join(hdrcnf.file_path.clone());
                        let map_path = Path::new(&parent).join(hdrcnf.map_path.clone());
                        match fs::read_to_string(file_path) {
                            Ok(code) => match fs::read_to_string(map_path) {
                                Ok(map_string) => {
                                    let map: HashMap<String, Variable> =
                                        serde_json::from_str(&map_string)
                                            .expect(format!("err_hdr.map: {}", filepath).as_str());
                                    let vars = Variables { vars: map };
                                    let module = Module {
                                        file_ws: filepath,
                                        mod_rs: format!("mod_{}", self.mod_num),
                                        file_rs: format!("build/mod_{}.rs", self.mod_num),
                                        _code: code.clone(),
                                        code_rs: code,
                                    };
                                    self.files.push(module.clone());
                                    variables.expand(vars);
                                    Ok(module.mod_rs)
                                }
                                Err(_) => Err(Problem {
                                    problem_type: ProblemType::FileNotFound,
                                    problem_msg: format!(
                                        "failed to import {}: Not found",
                                        hdrcnf.map_path
                                    ),
                                }),
                            },
                            Err(_) => Err(Problem {
                                problem_type: ProblemType::FileNotFound,
                                problem_msg: format!(
                                    "failed to import {}: Not found",
                                    hdrcnf.file_path
                                ),
                            }),
                        }
                    }
                    None => Err(Problem {
                        problem_type: ProblemType::HeaderSyntaxError,
                        problem_msg: format!("failed to deserialize"),
                    }),
                }
            }
            "dll" => {
                read_dll(filepath.clone());
                return Ok("xyz".to_string());
            }
            "rs" => Err(Problem {
                problem_type: ProblemType::FileNotFound,
                problem_msg: format!("failed to import {}: Not found", filepath),
            }),
            _ => Err(Problem {
                problem_type: ProblemType::FileNotFound,
                problem_msg: format!("failed to import {}: Not found", filepath),
            }),
        }
    }
}
