// rpc client to the rfk/rpc server.
//
//     ./rfk-cli skip
//
//     ./rfk-cli reward
//
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	rfkrpc "github.com/andynu/rfk/server/api/rpc"
)

func main() {
	for _, cmd := range os.Args[1:] {
		switch cmd {
		case "skip":
			simpleCall("Player.Skip")
		case "next":
			simpleCall("Player.SkipNoPunish")
		case "reward":
			simpleCall("Player.Reward")
		default:
			fmt.Printf("Unkown command %q\n", cmd)
		}
	}
}

// string command caller to the rpc with dummy in/out args.
func simpleCall(cmd string) {
	client := connect()
	var in rfkrpc.In
	var out rfkrpc.Out
	err := client.Call(cmd, in, &out)
	if err != nil {
		log.Fatal(err)
	}
}

// Connects to the running rpc server
func connect() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
