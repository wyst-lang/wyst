use pest::iterators::{Pair, Pairs};
use crate::utils::{ProblemCap};
use serde::{Deserialize, Serialize};

// The Parser
use pest::Parser;
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

#[derive(Deserialize, Serialize, Debug, Clone)]
pub struct Transpiler {
    pub problems: Vec<ProblemCap>,
    pub state: State,
    pub inject: Inject,
}
#[derive(Deserialize, Serialize, Debug, Clone)]
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
    for pair in pairs {arr.push(pair)}
    arr

}

impl Transpiler {
    pub fn transpile_pairs(&mut self, pairs: Pairs<Rule>) -> String {
        let mut res = String::new();
        for pair in pairs {
            res+=self.transpile(pair).as_str();
        }
        res
    }
    pub fn transpile(&mut self, pair: Pair<Rule>) -> String {
        let mut res = String::new();
        let tokens = to_indexable(pair.clone().into_inner());
        match pair.as_rule() {
            Rule::func_def => {
                res += format!("fn {}({}) -> {} {}", tokens[1].as_str(),
                               self.transpile_pairs(tokens[2].clone().into_inner()).as_str(),
                               tokens[0].as_str(),
                               self.transpile(tokens[3].clone())
                ).as_str();
            }
            Rule::var_def => {
                res += format!("let mut {}: {}", tokens[1].as_str(), tokens[0].as_str()).as_str();
            }
            Rule::curly => {
                res += "{";
                res += &self.transpile_pairs(pair.clone().into_inner());
                res += "}";
            }
            Rule::code_expr => {
                res += &self.transpile_pairs(pair.clone().into_inner());
            }
            //TODO: use LibManager to manage the includes
            Rule::include_global => {}
            Rule::include => {}

            _ => {
                res += " ";
                res += pair.as_str();
                res += " ";
            }
        }
        res
    }
    pub fn transpile_code(&mut self, code: String, indent: usize) -> String {
        let mut res = String::new();
        match WystParser::parse(Rule::top_expr, &code) {
            Ok(parsed) => {
                for pair in parsed {
                    res += self.transpile_pairs(pair.into_inner()).as_str();
                }
            }
            Err(e) => {
                // println!("err: {}", e.line());
                println!("{}", e);
            }
        }
        if indent > 0 {
            res += " ".repeat(indent * 2).as_str();
        }

        if indent > 0 {
            res += "\n";
            res += " ".repeat((indent - 1) * 2).as_str();
        }
        res
    }
}
