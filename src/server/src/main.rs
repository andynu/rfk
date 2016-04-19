mod player;
mod dj;
mod config;
mod song;
use config::Config;
use dj::DJ;
use player::Player;
use song::Song;

fn main() {
    let mut config = Config::load();

    let dj = DJ::new(&config);
    let player = Player::new(&config);

    let play_count_limit = config.play_count_limit;
    let limit_plays = play_count_limit > 0;
    let mut play_count = 0;
    for song in dj {
        player.wait_to_play();
        player.play(song);

        play_count += 1;
        if limit_plays && &play_count >= &play_count_limit {
            println!("reached play_count limit {}, stopping.", &play_count);
            break;
        }
    }
}

