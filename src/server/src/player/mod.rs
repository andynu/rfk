use std::time::Duration;
use std::thread::sleep;
use std::process::Command;

use config::Config;
use song::Song;

// player
#[derive(PartialEq)]
pub enum PlayerState {
    Playing,
    Paused,
    Stopped,
}

pub enum PlayerType {
    Simulated,
    Mp3
}



pub struct Player {
    state: PlayerState,
    song_player: PlayerType
}

impl Player {
    pub fn new(config: &Config) -> Player {
        Player{
            state: PlayerState::Paused, 
            song_player: match config.simulated {
                true => PlayerType::Simulated,
                false => PlayerType::Mp3
            }
        }
    }

    pub fn play(&self, song: Song) {
        println!("playing song");
        match self.song_player {
            PlayerType::Simulated => SimulatedPlayer::play(song),
            PlayerType::Mp3 => Mp3Player::play(song),
        }
    }

    pub fn pause(&mut self) {
        self.state = PlayerState::Paused;
    }

    pub fn wait_to_play(&self) {
        if self.state == PlayerState::Paused {
            let dur = Duration::from_secs(1);
            println!("paused: waiting {:?}", dur);
            sleep(dur);
        }
    }

}


trait SongPlayer {
    fn play(Song);
}

struct Mp3Player;
impl SongPlayer for Mp3Player {
    fn play(song: Song){
        let output = Command::new("mpg123").arg(&song.path).output().unwrap_or_else(|e|{
            panic!("failed to execute process: {}", e)
        });
        println!("> {}", String::from_utf8_lossy(&output.stdout));
    }
}

struct SimulatedPlayer;
impl SongPlayer for SimulatedPlayer {
    fn play(song: Song){
        let dur = Duration::from_secs(5);
        println!("simulated play: waiting {:?}", dur);
        sleep(dur);
    }
}
