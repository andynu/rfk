use song::Song;
use std::time::Duration;
use std::thread::sleep;
use std::process::Command;

// player
pub enum PlayerState {
    Playing,
    Paused,
    Stopped,
}

pub struct Player {
    paused: bool,
}

impl Player {
    pub fn new() -> Player {
        Player{paused: true}
    }

    pub fn play(&self, song: Song) {
        println!("playing song");
        let output = Command::new("mpg123").arg(&song.path).output().unwrap_or_else(|e|{
            panic!("failed to execute process: {}", e)
        });
        println!("> {}", String::from_utf8_lossy(&output.stdout));
    }

    pub fn pause(&mut self) {
        self.paused = false;
    }

    pub fn wait_to_play(&self) {
        if self.paused {
            let dur = Duration::from_secs(5);
            println!("waiting {:?}", dur);
            sleep(dur);
        } else {
            println!("nothing to wait for. proceed.");

        }
    }

}
