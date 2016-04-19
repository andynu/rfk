extern crate csv;
extern crate rand;

use std::fs::File;
use std::path::Path;
use std::io::BufRead;
use std::io::BufReader;
use std::vec::Vec;

use config::Config;
use song::Song;

pub enum QueueOrder {
    Consecutive,
    Random,
}

pub struct DJ { 
    songs: Vec<Song>
}
impl DJ {

    pub fn new(config: &Config) -> DJ {
        DJ{
            songs: DJ::load_songs(config.data_dir.to_string()), 
        }
    }

    //fn next_request_song(&self, queueOrder: QueueOrder) -> Option<Song> {
    //    //let songs = self.load_songs();

    //    //let songIdx = match queueOrder {
    //    //    QueueOrder::Consecutive => self.nextIdx(),
    //    //    QueueOrder::Random => self.randIdx(),
    //    //};

    //    //songs[songIdx].clone();
    //    None
    //}

    fn next_all_song(&self, queueOrder: QueueOrder) -> Option<Song> {

        let songIdx = match queueOrder {
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

    fn nextIdx(&self) -> usize {
        0
    }

    fn randIdx(&self) -> usize {
        let rng = rand::
        
        let num: f64 = rng.gen_range(0,self.songs.len());
        let idx = num as usize;
        idx
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

        let queueOrder = QueueOrder::Random;

        //match self.next_request_song(queueOrder) {
        //    Some(s) => return Some(s),
        //    None => println!("No requests.")
        //}

        match self.next_all_song(queueOrder) {
            Some(s) => return Some(s),
            None => println!("No songs.")
        }

        panic!("No Song!")
    }
}
