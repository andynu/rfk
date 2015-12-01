// Interprets impressions into library.Song.Rank
package karma

import (
	"github.com/andynu/rfk/server/config"
	"github.com/andynu/rfk/server/library"
	"log"
	"os"
	"path"
	"sync"
)

var logger *log.Logger
var mu sync.Mutex

func init() {
	logfile, err := os.OpenFile(path.Join(config.Config.DataPath, "impression.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	logger = log.New(logfile, "", 0)
}

// Record a positive/negative impression of a Song.
func Log(song library.Song, impression int) {
	mu.Lock()
	logger.Printf("%s\t%d", song.Hash, impression)
	mu.Unlock()
}
