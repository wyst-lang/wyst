use crate::transpiler::{Rule, State, Transpiler, WystParser};
use dirs::home_dir;
use pest::Parser;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs;
use std::path::Path;

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum ProblemType {
    VariableNotFound,
    FileNotFound,
    HeaderSyntaxError,
    SyntaxError,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum ProblemCap {
    Error(Problem),
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Problem {
    pub problem_type: ProblemType,
    pub problem_msg: String,
    pub state: State,
}

#[derive(Debug, Clone)]
pub struct Module {
    pub rs_mod: String,
    pub file_path: String,
    pub code: String,
}

#[derive(Debug, Clone)]
pub struct ModManager {
    pub modules: Vec<Module>,
    pub macros: Vec<String>,
    pub map: HashMap<String, String>,
}

impl ModManager {
    pub fn new() -> ModManager {
        ModManager {
            modules: Vec::new(),
            macros: Vec::new(),
            map: HashMap::new(),
        }
    }
    pub fn check(&mut self, path: String) -> Option<Module> {
        for module in self.modules.clone() {
            if module.file_path == path {
                return Some(module);
            }
        }
        None
    }
    pub fn add(&mut self, root: &mut Transpiler, file_path: String, global: bool) -> String {
        let mut problems: Vec<ProblemCap> = Vec::new();
        let mod_name = format!("mod_{}", self.modules.len());
        let home_dirpb = home_dir().unwrap();
        let home_dir = home_dirpb.to_str().unwrap();
        let mut file_path = file_path;
        if global {
            let mut tmp = Path::new(home_dir)
                .join("wyst")
                .join("lib")
                .join(file_path.clone())
                .to_str()
                .unwrap()
                .to_string();
            if !Path::new(&tmp).exists() {
                tmp = Path::new("lib")
                    .join(file_path)
                    .to_str()
                    .unwrap()
                    .to_string();
            }
            file_path = tmp;
        }
        if let Some(modl) = self.check(file_path.clone()) {
            return modl.rs_mod;
        }
        if let Ok(code) = fs::read_to_string(file_path.clone()) {
            match file_path.rsplit_once('.').unwrap_or(("", "")).1 {
                "rs" => match WystParser::parse(Rule::rust_code, &code) {
                    Ok(parsed) => {
                        self.modules.push(Module {
                            rs_mod: mod_name.clone(),
                            file_path: file_path.clone(),
                            code: code.clone(),
                        });
                        for pair in parsed {
                            for p in pair.into_inner() {
                                if p.as_rule() == Rule::rust_macro_lint {
                                    self.macros.push(p.into_inner().as_str().to_string());
                                }
                            }
                        }
                    }
                    Err(_) => {}
                },
                "h" => {
                    file_path = file_path.rsplit_once('.').unwrap_or(("", "")).0.to_string();
                    file_path += ".rs";
                    let parsed = WystParser::parse(Rule::header_expr, &code).unwrap();
                    for pair in parsed {
                        root.transpile_pairs(pair.into_inner());
                    }
                    self.map.extend(root.mod_manager.map.clone().into_iter());
                    if let Ok(code1) = fs::read_to_string(file_path.clone()) {
                        let parsed = WystParser::parse(Rule::rust_code, &code1).unwrap();
                        self.modules.push(Module {
                            rs_mod: mod_name.clone(),
                            file_path: file_path.clone(),
                            code: code1.clone(),
                        });
                        for pair in parsed {
                            for p in pair.into_inner() {
                                if p.as_rule() == Rule::rust_macro_lint {
                                    self.macros.push(p.into_inner().as_str().to_string());
                                }
                            }
                        }
                    } else {
                        problems.push(ProblemCap::Error(Problem {
                            problem_msg: format!("Sub Header not found '{}'", file_path),
                            problem_type: ProblemType::FileNotFound,
                            state: State {
                                line: 0,
                                column: 0,
                                file: None,
                            },
                        }))
                    }
                }
                _ => {}
            }
        } else {
            problems.push(ProblemCap::Error(Problem {
                problem_msg: format!("File not found '{}'", file_path),
                problem_type: ProblemType::FileNotFound,
                state: State {
                    line: 0,
                    column: 0,
                    file: None,
                },
            }))
        }
        mod_name
    }
    pub fn write(&mut self) -> Vec<ProblemCap> {
        let mut problems: Vec<ProblemCap> = Vec::new();
        for module in self.modules.clone() {
            let file_path = module.file_path;
            match file_path.rsplit_once('.').unwrap_or(("", "")).1 {
                "wst" => {
                    let mut trsp = Transpiler {
                        mod_manager: self.clone(),
                        ..Default::default()
                    };
                    let rust_code = trsp.transpile_code(module.code, 0).expect("");
                    fs::write(
                        Path::new("build").join((module.rs_mod + ".rs").as_str()),
                        rust_code,
                    )
                    .expect("Error writing rust file: ");
                    problems.append(&mut trsp.problems);
                    self.modules = trsp.mod_manager.modules;
                }

                "rs" => fs::write(
                    Path::new("build").join((module.rs_mod + ".rs").as_str()),
                    &module.code,
                )
                .unwrap(),

                _ => {}
            }
        }
        problems
    }
}
