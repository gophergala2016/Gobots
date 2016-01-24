package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Note: The comment under this became irrelevant like 10 minutes into
// development, it's only here for posterity

// TODO(bsprague): LITERALLY EVERYTHING. Off the top:

// - The server - It needs to be able to serve pages, handle matches, do OAuth
// with Github, and implement gRPCs bidirectional streaming

// - The library - Needs a base Bot that can be embedded in higher level bots,
// and needs to know how to handshake with the normal server

// - The game - Needs to exist. Should be in a separate subpackage?? In any
// case, gopherjs should be used so that the implementation only needs to be
// written once

var (
	addr      = flag.String("addr", ":8000", "server address")
	templates = tmpl{template.Must(template.ParseGlob("templates/*.html"))}

	// You know, this should really be unique to each player, but I'm
	// not ready to do cookie stuff yet
	// TODO: Make this cookie-based...and not super easily hackable
	imAnIdiot = genName(128)
	secretz   string
)

const (
	clientId = "07ef388cb32ffbbd5146"
)

func main() {
	flag.Parse()

	if dat, err := ioutil.ReadFile("secretz"); err == nil {
		secretz = string(dat)
	} else {
		log.Fatal("Ain't got no GitHub client secret!!")
	}

	http.HandleFunc("/", withLogin(serveIndex))
	http.HandleFunc("/auth", withLogin(serveAuth))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Yeah...so about that whole server thing: ", err)
	}
}

func serveIndex(c context) {
	if err := templates.ExecuteTemplate(c, "index.html", map[string]interface{}{}); err != nil {
		serveError(c.w, err)
	}
}

func serveError(w http.ResponseWriter, err error) {
	w.Write([]byte("Internal Server Error"))
	log.Printf("Error: %v\n", err)
}

func serveAuth(c context) {
	if c.r.FormValue("state") != imAnIdiot {
		// Pssh they couldn't even fake my publically-accessible secret? Laaaaaame
		return
	}

	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"client_id":     []string{clientId},
		"client_secret": []string{secretz},
		"code":          []string{c.r.FormValue("code")},
	})
	if err != nil {
		return
	}

	defer resp.Body.Close()
	// TODO: Read the access_token from the body, and then use it to get the
	// player's username. Store that in an encrypted cookie
}
