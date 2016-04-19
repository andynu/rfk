# RFK

A jukebox that plays random songs. (for now)

**This is a toy project as I learn rust.**

## Prereqs

A few assumptions:

* You're running linux, or a POSIX compatible shell.
* You've cloned this repository to ~/rfk


### Install go
You'll need go. See https://golang.org/dl/

    cd
    wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz    
    tar xvfz go1.5.1.linux-amd64.tar.gz

Setting up your go env for this project:

    export PATH=$PATH:~/go/bin
    export GO15VENDOREXPERIMENT=1
    export GOPATH=~/go

### Install rfk

Grab a copy of this repository via the `go get` command:

    go get github.com/andynu/rfk

This will put the code in ~/go/src/github.com/andynu/rfk

### Install mpg123

For the moment only files that mpg123 can play are supported. mpg123 is expected to be in your path.

    sudo apt-get install mpg123

or download it from http://www.mpg123.de/

### Configure your music library

RFK knows about your music from a songs.txt file that contains absolute paths to your song files (one per line).

    mkdir -p ~/rfk/data/$(hostname)/
    find / -name '*.mp3' > ~/rfk/data/$(hostname)/songs.txt

## Build 


    cd ~/go/src/github.com/andynu/rfk
    make


## Run Server

    ./rfk-server

You should now be listening to random music.

### Run Client


    ./rfk-cli skip

or

    ./rfk-cli reward

### Hash your songs

    cat songs.txt | ./rfk-ident > song_hashes.txt

----

## Attributions

* Thanks to David Howden for the dhowden/tag library, used for audio only checksums and mp3 metadata extraction. See [LICENSE](src/github.com/dhowden/tag/LICENSE)
