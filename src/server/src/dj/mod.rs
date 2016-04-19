extern crate csv;
extern crate rand;

use std::fs::File;
use std::path::Path;
use std::io::BufRead;
use std::io::BufReader;
use std::vec::Vec;

use config::Config;
use song::Song;

#[derive(Clone)]
pub enum QueueOrder {
    Consecutive,
    Random,
}

pub struct DJ { 
    songs: Vec<Song>,
    consecutive_idx: usize,
    play_order: QueueOrder,
}
impl DJ {

    pub fn new(config: &Config) -> DJ {
        DJ{
            songs: DJ::load_songs(config.data_dir.to_string()), 
            consecutive_idx: 0,
            play_order: config.play_order.clone(),
        }
    }

    //fn next_request_song(&self, queueOrder: QueueOrder) -> Option<Song> {

    //    let songIdx = match queueOrder {
    //        QueueOrder::Consecutive => self.nextIdx(),
    //        QueueOrder::Random => self.randIdx(),
    //    };

    //    songs[songIdx].clone();
    //    None
    //}

    fn next_all_song(&mut self) -> Option<Song> {

        let songIdx = match self.play_order {
            QueueOrder::Consecutive => self.nextIdx(),
            QueueOrder::Random => self.randIdx(),
        };
        println!("songIdx: {}", songIdx);

        if songIdx < self.songs.len() {
            Some(self.songs[songIdx].clone())
        } else {
            None
        }
    }

    fn nextIdx(&mut self) -> usize {
        self.consecutive_idx += 1;
        self.consecutive_idx
    }

    fn randIdx(&self) -> usize {
        let rn = rand::random::<f64>();
        let maxIdx = self.songs.len() as f64;
        let idx = rn * maxIdx;
        idx as usize
    }

    fn load_songs(data_dir: String) -> Vec<Song> {
        let songs_txt = data_dir + "/song_hashes.ascii.txt";
        let mut song_count = 0;
        let mut songs = Vec::new();
        let mut reader = csv::Reader::from_file(songs_txt).unwrap();
        for record in reader.decode() {
            let (hash,path,_,_,_): 
                (String, String, String, String, String) 
                 = record.unwrap();
            if Path::new(&path).exists() {
                let song = Song{
                    id: song_count,
                    path: path,
                };
                //println!("{:?}", song);
                songs.push(song);
                song_count += 1;
            }
        }
        println!("song_count = {}", song_count);
        songs
    }

}


impl Iterator for DJ {
    type Item = Song;
    fn next(&mut self) -> Option<Song> {

        //match self.next_request_song(queueOrder) {
        //    Some(s) => return Some(s),
        //    None => println!("No requests.")
        //}

        match self.next_all_song() {
            Some(s) => return Some(s),
            None => println!("No songs.")
        }

        panic!("No Song!")
    }
}
