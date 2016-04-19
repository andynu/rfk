use std::time::Duration;
use std::thread::sleep;
use std::process::Command;

use config::Config;
use song::Song;

// player
pub enum PlayerState {
    Playing,
    Paused,
    Stopped,
}

pub struct Player {
    paused: bool,
    simulated: bool,
}

impl Player {
    pub fn new(config: &Config) -> Player {
        Player{
            paused: false, 
            simulated: config.simulated,
        }
    }

    pub fn play(&self, song: Song) {
        println!("playing song");
        if self.simulated {
            let dur = Duration::from_secs(5);
            println!("simulated play: waiting {:?}", dur);
            sleep(dur);
        } else {
            let output = Command::new("mpg123").arg(&song.path).output().unwrap_or_else(|e|{
                panic!("failed to execute process: {}", e)
            });
            println!("> {}", String::from_utf8_lossy(&output.stdout));
        }
    }

    pub fn pause(&mut self) {
        self.paused = false;
    }

    pub fn wait_to_play(&self) {
        if self.paused {
            let dur = Duration::from_secs(1);
            println!("paused: waiting {:?}", dur);
            sleep(dur);
        }
    }

}
