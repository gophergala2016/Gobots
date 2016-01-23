package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// TODO(bsprague): LITERALLY EVERYTHING. Off the top:

// - The server - It needs to be able to serve pages, handle matches, do OAuth
// with Github, and implement gRPCs bidirectional streaming

// - The library - Needs a base Bot that can be embedded in higher level bots,
// and needs to know how to handshake with the normal server

// - The game - Needs to exist. Should be in a separate subpackage?? In any
// case, gopherjs should be used so that the implementation only needs to be
// written once

var addr = flag.String("addr", ":8000", "server address")
var templates = tmpl{template.Must(template.ParseGlob("templates/*.html"))}

func main() {
	flag.Parse()

	http.HandleFunc("/", withLogin(serveIndex))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Yeah...so about that whole server thing: ", err)
	}
}

func serveIndex(c context) {
	if err := templates.ExecuteTemplate(c, "index.html", struct{}{}); err != nil {
		serveError(c.w, err)
	}
}

func serveError(w http.ResponseWriter, err error) {
	w.Write([]byte("Internal Server Error"))
	log.Printf("Error: %v\n", err)
}
