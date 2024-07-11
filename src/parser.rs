use pest::Parser;
use pest_derive::Parser;

#[derive(Parser)]
#[grammar = "wyst.pest"]
pub struct WystParser;

pub fn parse(input: String, rule: Rule) {
    match WystParser::parse(rule, &input) {
        Ok(parsed) => {
            for pair in parsed {
                println!("---------START--------");
                println!("{:?}", pair);
                println!("---------END--------");
            }
            // println!("{:#?}", parsed);
        }
        Err(e) => {
            println!("Ok: {}", e);
        }
    }
}
