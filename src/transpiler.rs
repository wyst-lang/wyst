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
                                let mut ident_rs = String::from("_0x");
                                let ident_chars: Vec<char> = ident.chars().collect();
                                for i in 0..ident_chars.len() {
                                    let current = ident_chars[i];
                                    if ident_chars.len() > i + 1 {
                                        let next = ident_chars[i + 1];
                                        if current == '.' {
                                            ident_rs += "._0x";
                                        } else if current == ':' && next == ':' {
                                            ident_rs += "::_0x";
                                        } else {
                                            if current != ':' && current != '.' {
                                                ident_rs +=
                                                    hex::encode(current.to_string()).as_str();
                                            }
                                        }
                                    } else {
                                        if current == '.' {
                                            ident_rs += "._0x";
                                        } else {
                                            ident_rs += hex::encode(current.to_string()).as_str();
                                        }
                                    }
                                }
                                if self.mod_manager.macros.contains(&ident_rs) {
                                    output += ident_rs.as_str();
                                    output += "!";
                                } else {
                                    output += match ident {
                                        "void" => "()",
                                        "int" => "i32",
                                        "float" => "f64",
                                        "Vec" => "Vec",
                                        "bool" => "bool",
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
    pub fn transpile_vec(&mut self, pairs: Vec<Pair<Rule>>) -> String {
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
                        .as_str()
                        .replace("let mut", ""),
                    self.transpile(tokens[0].clone()),
                    self.transpile(tokens[3].clone())
                )
                .as_str();
            }
            Rule::ptr_fn_def => {
                res += format!(
                    "pub fn {}({}) -> &mut {} {}\n",
                    self.transpile(tokens[1].clone()),
                    self.transpile_pairs(tokens[2].clone().into_inner())
                        .as_str()
                        .replace("let mut", ""),
                    self.transpile(tokens[0].clone()),
                    self.transpile(tokens[3].clone())
                )
                .as_str();
            }
            Rule::var_def_ptr => {
                res += format!(
                    "let mut {}: &mut {}",
                    self.transpile(tokens[1].clone()),
                    self.transpile(tokens[0].clone())
                )
                .as_str();
            }
            Rule::var_def => {
                if tokens.len() == 1 {
                    res += self.transpile(tokens[0].clone()).as_str();
                } else {
                    res += format!(
                        "let mut {}: {}",
                        self.transpile(tokens[1].clone()),
                        self.transpile(tokens[0].clone())
                    )
                    .as_str();
                }
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

            Rule::term => {
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

            Rule::enum_def => {
                res += format!(
                    "enum {} {}{}{}",
                    self.transpile(tokens[0].clone()),
                    "{",
                    self.transpile_pairs(tokens[1].clone().into_inner()),
                    "}\n"
                )
                .as_str();
            }

            Rule::ecall => {
                res += self.transpile(tokens[0].clone()).as_str();
                if tokens.len() > 1 {
                    let mut params = tokens.clone();
                    params.remove(0);
                    res += "(";
                    res += self.transpile_vec(params).as_str();
                    res += ")";
                }
            }

            Rule::comma => {
                res += ",";
            }

            Rule::ptr_set => {
                res += format!(
                    "*{} = {}",
                    self.transpile(tokens[0].clone()),
                    self.transpile(tokens[1].clone()),
                )
                .as_str();
            }

            Rule::int => {
                res += " ";
                res += pair.as_str();
                res += " ";
            }

            Rule::refrence => {
                res += "&mut ";
                res += self.transpile(tokens[0].clone()).as_str();
            }

            Rule::return_stm => {
                res += "return ";
                res += self.transpile(tokens[0].clone()).as_str();
            }

            Rule::if_stm => {
                res += "if ";
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::else_if_stm => {
                res += "else if ";
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::else_stm => {
                res += "else ";
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::bool => {
                res += pair.as_str();
            }

            Rule::operator => {
                res += pair.as_str();
            }

            Rule::if_operator => {
                res += pair.as_str();
            }

            Rule::if_expr => {
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::namespace_def => {
                res += "mod ";
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::top_code => {
                res += "{";
                res += self.transpile_pairs(pair.into_inner()).as_str();
                res += "}";
            }

            Rule::top_expr => {
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::impl_def => {
                res += "impl ";
                res += self.transpile_pairs(pair.into_inner()).as_str();
            }

            Rule::match_name => {
                let to_match = tokens[0].clone().as_str();
                let match_id = tokens[1].clone().as_str();
                res += "\n";
                res += self.encode_ident(to_match).as_str();
                res += "<match:";
                res += match_id;
                res += ">";
            }

            Rule::operation => {
                for i in 0..tokens.len() {
                    let token = tokens[i].clone();
                    if i % 2 == 0 {
                        res += self.transpile(token).as_str();
                    } else {
                        res += token.as_str();
                    }
                }
            }

            _ => {
                println!("\x1b[33mNot Handled: {}\x1b[0m", pair);
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
