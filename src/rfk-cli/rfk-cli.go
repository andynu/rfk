package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	rfkrpc "rfk/rpc"
)

func main() {
	for _, cmd := range os.Args[1:] {
		switch cmd {
		case "skip":
			simpleCall("Player.Skip")
		case "stop":
			simpleCall("Player.Stop")
		case "reward":
			simpleCall("Player.Reward")
		default:
			fmt.Printf("Unkown command %q\n", cmd)
		}
	}
}

func simpleCall(cmd string) {
	client := connect()
	var in rfkrpc.In
	var out rfkrpc.Out
	err := client.Call(cmd, in, &out)
	if err != nil {
		log.Fatal(err)
	}
}

func connect() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
