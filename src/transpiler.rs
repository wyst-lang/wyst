use crate::utils::{ModManager, ProblemCap};
use pest::iterators::{Pair, Pairs};
use serde::{Deserialize, Serialize};

// The Parser
use pest::{error::Error, Parser};
use pest_derive::Parser;
#[derive(Parser)]
#[grammar = "wyst.pest"]
pub struct WystParser;

#[derive(Deserialize, Serialize, Debug, Clone)]
pub struct State {
    pub line: u32,
    pub column: u32,
    pub file: Option<String>,
}

#[derive(Debug, Clone)]
#[allow(dead_code)]
pub struct Transpiler {
    pub problems: Vec<ProblemCap>,
    pub state: State,
    pub inject: Inject,
    pub mod_manager: ModManager,
}
#[derive(Debug, Clone)]
#[allow(dead_code)]
pub struct Inject {
    pub state: State,
    pub inject: bool,
}

impl Default for Inject {
    fn default() -> Self {
        Inject {
            inject: false,
            state: State {
                line: 0,
                column: 0,
                file: None,
            },
        }
    }
}

impl Default for Transpiler {
    fn default() -> Self {
        Transpiler {
            inject: Inject::default(),
            problems: Vec::new(),
            mod_manager: ModManager::new(),
            state: State {
                line: 1,
                column: 0,
                file: None,
            },
        }
    }
}

pub fn to_indexable(pairs: Pairs<Rule>) -> Vec<Pair<Rule>> {
    let mut arr: Vec<Pair<Rule>> = Vec::new();
    for pair in pairs {
        arr.push(pair)
    }
    arr
}

impl Transpiler {
    pub fn encode_ident(&mut self, ident: &str) -> String {
        match WystParser::parse(Rule::def_identifier, ident) {
            Ok(parsed) => {
                let mut output = String::new();
                for pair in parsed {
                    let tokens = to_indexable(pair.clone().into_inner());
                    for t in tokens {
                        match t.as_rule() {
                            Rule::identifier => {
                                let ident = t.as_str().trim();
                                let ident_rs = format!("_0x{}", hex::encode(ident.as_bytes()));
                                if self.mod_manager.macros.contains(&ident_rs) {
                                    output += ident_rs.as_str();
                                    output += "!";
                                } else {
                                    output += match ident {
                                        "void" => "()",
                                        "int" => "i32",
                                        "float" => "f64",
                                        "Vec" => "Vec",
                                        _ => {
                                            if self.mod_manager.map.contains_key(ident) {
                                                self.mod_manager.map.get(ident).unwrap()
                                            } else {
                                                ident_rs.as_str()
                                            }
                                        }
                                    };
                                }
                            }
                            Rule::def_identifier => {
                                output += "<";
                                output += self.encode_ident(t.as_str()).as_str();
                                output += ">";
                            }
                            _ => {}
                        }
                    }
                }
                output
            }
            Err(_) => String::new(),
        }
    }
    pub fn transpile_pairs(&mut self, pairs: Pairs<Rule>) -> String {
        let mut res = String::new();
        for pair in pairs {
            res += self.transpile(pair).as_str();
        }
        res
    }
    pub fn transpile(&mut self, pair: Pair<Rule>) -> String {
        let mut res = String::new();
        let tokens = to_indexable(pair.clone().into_inner());
        match pair.as_rule() {
            Rule::func_def => {
                res += format!(
                    "pub fn {}({}) -> {} {}\n",
                    self.transpile(tokens[1].clone()),
                    self.transpile_pairs(tokens[2].clone().into_inner())
                        .as_str(),
                    self.transpile(tokens[0].clone()),
                    self.transpile(tokens[3].clone())
                )
                .as_str();
            }
            Rule::var_def => {
                res += format!(
                    "let mut {}: {}",
                    self.transpile(tokens[1].clone()),
                    self.transpile(tokens[0].clone())
                )
                .as_str();
            }
            Rule::var_def_set => {
                res += format!(
                    "{} = {}",
                    self.transpile(tokens[0].clone()),
                    self.transpile(tokens[1].clone()),
                )
                .as_str();
            }
            Rule::curly => {
                res += "{";
                res += &self.transpile_pairs(pair.clone().into_inner());
                res += "}";
            }
            Rule::code_expr => {
                res += &self.transpile_pairs(pair.clone().into_inner());
            }

            Rule::include_global => {
                res += "#[allow(unused_imports)]\nuse ";
                res += self
                    .mod_manager
                    .add(&mut self.clone(), tokens[0].as_str().to_string(), true)
                    .as_str();
                res += "::*;\n";
            }
            Rule::include => {
                res += "#[allow(unused_imports)]\nuse ";
                res += self
                    .mod_manager
                    .add(&mut self.clone(), tokens[0].as_str().to_string(), false)
                    .as_str();
                res += "::*;\n";
            }

            Rule::identifier | Rule::def_identifier => {
                res += self.encode_ident(pair.as_str().trim()).as_str();
            }

            Rule::call => {
                res += " ";
                res += self.transpile(tokens[0].clone()).as_str();
                res += self.transpile(tokens[1].clone()).as_str();
            }

            Rule::round => {
                res += "(";
                res += self.transpile_pairs(pair.into_inner()).as_str();
                res += ")";
            }

            Rule::expr => {
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::expr_ => {
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::struct_def => {
                res += "struct ";
                res += self.transpile(tokens[0].clone()).as_str();
                res += "{";
                res += self.transpile(tokens[1].clone()).as_str();
                res += "}";
            }

            Rule::semicolon => {
                res += ";\n";
            }

            Rule::string => {
                res += pair.as_str();
            }

            Rule::map => {
                self.mod_manager.map.insert(
                    tokens[0].as_str().to_string(),
                    tokens[2].as_str().to_string(),
                );
            }

            _ => {
                res += " ";
                res += pair.as_str();
                res += " ";
            }
        }
        res
    }
    pub fn transpile_code(&mut self, code: String, indent: usize) -> Result<String, Error<Rule>> {
        let mut res = String::new();
        match WystParser::parse(Rule::top_expr, &code) {
            Ok(parsed) => {
                for pair in parsed {
                    res += self.transpile_pairs(pair.into_inner()).as_str();
                }
            }
            Err(e) => {
                return Err(e);
            }
        }
        if indent > 0 {
            res += " ".repeat(indent * 2).as_str();
        }

        if indent > 0 {
            res += "\n";
            res += " ".repeat((indent - 1) * 2).as_str();
        }
        Ok(res)
    }
}
