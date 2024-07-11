use pest::Parser;
use pest_derive::Parser;

#[derive(Parser)]
#[grammar = "wyst.pest"]
pub struct WystParser;

pub fn parse(input: String, rule: Rule) {
    match WystParser::parse(rule, &input) {
        Ok(parsed) => {
            for pair in parsed {
                for inner in pair.into_inner() {
                    println!("{:?}", inner.into_inner())
                }
            }
            // println!("{:#?}", parsed);
        }
        Err(e) => {
            println!("Ok: {}", e);
        }
    }
}
