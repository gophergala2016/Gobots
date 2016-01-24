package main

import (
	"github.com/gophergala2016/Gobots/engine"
	"github.com/gopherjs/gopherjs/js"
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
	// TODO: replayString -> []byte -> Replay -> js.Object
	// In the mean time, fake board for testing
	return js.MakeWrapper(engine.NewBoard(8, 8))
}

func GetReplay(gameID string) *js.Object {
	// TODO: get bytes from server -> Replay -> js.Object
	return js.MakeWrapper(&engine.Board{})
}
