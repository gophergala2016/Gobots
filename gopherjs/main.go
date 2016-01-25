package main

import (
	"github.com/gophergala2016/Gobots/botapi"
	"github.com/gophergala2016/Gobots/engine"
	"github.com/gopherjs/gopherjs/js"
	"zombiezen.com/go/capnproto2"
)

func main() {
	// TODO(bsprague): Add things to this handy global JS object when you get
	// even remotely that far
	js.Global.Set("Gobot", map[string]interface{}{
		"GetReplayFromString": GetReplayFromString,
		"GetReplay":           GetReplay,
	})
}

func GetReplayFromString(replayString string) *js.Object {
	// Will it work? It's not Unicode clean
	msg, _ := capnp.Unmarshal([]byte(replayString))
	r, _ := botapi.ReadRootReplay(msg)
	return js.MakeWrapper(r)
}

func GetReplay(gameID string) *js.Object {
	// TODO: get bytes from server -> Replay -> js.Object
	return js.MakeWrapper(&engine.Board{})
}
