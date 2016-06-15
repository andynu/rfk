// rpc server endpoint
package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/andynu/rfk/server/api"
)

func Listener() {
	rpcPlayer := new(Player)
	rpc.Register(rpcPlayer)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":7777")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}

type In int
type Out int
type Player struct{}

func (t *Player) Skip(in In, out *Out) error {
	log.Println("rpc: skip")
	api.Skip()
	return nil
}

func (t *Player) SkipNoPunish(in In, out *Out) error {
	log.Println("rpc: skip (no punish)")
	api.SkipNoPunish()
	return nil
}

func (t *Player) Reward(in In, out *Out) error {
	log.Println("rpc: reward")
	api.Reward()
	return nil
}

func (t *Player) Pause(in In, out *Out) error {
	log.Println("rpc: play/pause")
	api.Pause()
	return nil
}

func (t *Player) Unpause(in In, out *Out) error {
	log.Println("rpc: play/pause")
	api.Unpause()
	return nil
}
