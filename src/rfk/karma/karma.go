package karma

import (
	"log"
	"os"
	"path"
	"rfk/config"
	"rfk/library"
)

var logger *log.Logger

func init() {
	logfile, err := os.OpenFile(path.Join(config.Config.DataPath, "impression.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(logfile, "", 0)
}

func Log(song library.Song, impression int) {
	logger.Printf("%s\t%d", song.Hash, impression)
}
