package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// TODO(bsprague): Add things to this handy global JS object when you get
	// even remotely that far
	js.Global.Set("Gobot", map[string]interface{}{})
}
