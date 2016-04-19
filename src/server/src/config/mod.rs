extern crate getopts;
use std::env;
use std::path::Path;

pub struct Config {
    pub play_count_limit: usize,
    pub data_dir: String,
}

impl Config {

    pub fn new() -> Config {
        Config{ 
            play_count_limit: 0,
            data_dir: "/home/andy/rfk/data".to_string()
        }
    }

    pub fn load() -> Config {
        let mut config = Config::new();
        config.apply_args();
        config
    }

    pub fn apply_args(&mut self){

        let args: Vec<String> = env::args().collect();
        let program = args[0].clone();

        let mut opts = getopts::Options::new();
        opts.optopt("l", "limit", "Limit the number of songs played.", "PLAY_COUNT");

        let matches = match opts.parse(&args[1..]) {
            Ok(m) => { m },
            Err(f) => { panic!(f.to_string()) },
        };

        if matches.opt_present("limit") {
            self.play_count_limit = matches.opt_str("l").unwrap()
                .parse().unwrap();
        }

    }
}
