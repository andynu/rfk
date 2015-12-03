// rpc server endpoint
package rpc

import (
	"github.com/andynu/rfk/server/karma"
	"github.com/andynu/rfk/server/player"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type In int
type Out int
type Player struct{}

func (t *Player) Skip(in In, out *Out) error {
	log.Println("rpc: skip")
	player.Skip()
	return nil
}

func (t *Player) Stop(in In, out *Out) error {
	log.Println("rpc: stop")
	player.Stop()
	return nil
}

func (t *Player) Reward(in In, out *Out) error {
	log.Println("rpc: reward")
	karma.Log(player.CurrentSong, 1)
	return nil
}

func SetupRPC() {
	rpcPlayer := new(Player)
	rpc.Register(rpcPlayer)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":7777")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
