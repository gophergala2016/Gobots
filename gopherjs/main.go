package main

import (
	"io/ioutil"
	"net/http"

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
	return js.MakeWrapper(engine.NewReplay(r))
}

func GetReplay(url string) *js.Object {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	d, _ := ioutil.ReadAll(resp.Body)
	msg, _ := capnp.Unmarshal(d)
	r, _ := botapi.ReadRootReplay(msg)
	return js.MakeWrapper(engine.NewReplay(r))
}
