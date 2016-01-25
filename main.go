package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/securecookie"
)

var (
	addr      = flag.String("addr", ":8000", "HTTP server address")
	apiAddr   = flag.String("api_addr", ":8001", "RPC server address")
	templates = tmpl{template.Must(template.ParseGlob("templates/*.html"))}

	db               datastore
	secretz          string
	s                *securecookie.SecureCookie
	globalAIEndpoint *aiEndpoint
)

const (
	clientId = "07ef388cb32ffbbd5146"
)

func main() {
	flag.Parse()
	var err error

	if db, err = initDB("gobots.db"); err != nil {
		log.Fatal("Couldn't open the database, SHUT IT DOWN")
	}

	if secretz, err = initSecretz(); err != nil {
		log.Fatal("Ain't got no GitHub client secret!!")
	}

	if s, err = initKeys(); err != nil {
		log.Fatal("Can't encrypt the cookies! WHATEVER WILL WE DO")
	}

	http.HandleFunc("/", withLogin(serveIndex))
	http.HandleFunc("/game/", withLogin(serveGame))
	http.HandleFunc("/auth", withLogin(serveAuth))
	http.HandleFunc("/loadBots", withLogin(loadBots))

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	globalAIEndpoint, err = startAIEndpoint(*apiAddr, db)
	if err != nil {
		log.Fatal("AI RPC endpoint failed to start:", err)
	}

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Yeah...so about that whole server thing: ", err)
	}
}

func serveIndex(c context) {
	data := tmplData{
		Data: map[string]interface{}{
			"Bots": globalAIEndpoint.listOnlineAIs(),
		},
		Scripts: []template.URL{
			"/js/main.js",
		},
	}

	if err := templates.ExecuteTemplate(c, "index.html", data); err != nil {
		serveError(c.w, err)
	}
}

func serveGame(c context) {
	replay, _ := db.lookupGame(c.gameID())

	data := tmplData{
		Data: map[string]interface{}{
			"Replay": replay,
			"GameID": c.gameID(),
			// TODO: Use this one when not testing
			//"Exists": err != errDatastoreNotFound,
			"Exists": true,
		},
	}
	if err := templates.ExecuteTemplate(c, "game.html", data); err != nil {
		serveError(c.w, err)
	}
}

func serveError(w http.ResponseWriter, err error) {
	w.Write([]byte("Internal Server Error"))
	log.Printf("Error: %v\n", err)
}

func loadBots(c context) {
	uid := userID(c.p.Name)
	_, token, err := db.createAI(uid, &aiInfo{Nick: c.r.PostFormValue("endpoint")})
	if err != nil {
		serveError(c.w, err)
		return
	}
	fmt.Fprintln(c.w, "Congrats, your token is:", token)
}

func serveAuth(c context) {
	if c.r.FormValue("state") != c.magicToken {
		log.Println("They're spoofing GitHub's API. I AM THE ONE WHO KNOCKS (on GitHub's API server)")
		return
	}

	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"client_id":     {clientId},
		"client_secret": {secretz},
		"code":          {c.r.FormValue("code")},
	})
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	sRep := string(d)
	// http://i.imgur.com/c4jt321.png
	is, ie := strings.Index(sRep, "=")+1, strings.Index(sRep, "&")
	accessToken := sRep[is:ie]

	nutFact, err := loadCookie(c.r)
	if err != nil {
		log.Println(err)
		// This is weird, they must have cookies off, which makes it hard for us to
		// validate them. Screw 'em for now
		return
	}

	// Set their access token to what we just got, and create a user with that token
	if nutFact.AccessToken == "" {
		nutFact.AccessToken = accessToken
		go db.createUser(userID(accessToken))
	}

	if encoded, err := s.Encode("info", nutFact); err == nil {
		cookie := &http.Cookie{
			Name:  "info",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(c.w, cookie)
	}

	http.Redirect(c.w, c.r, "/", http.StatusFound)
}

func username(uID userID) string {
	resp, err := http.Get("https://api.github.com/user?access_token=" + string(uID))
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		log.Println(err)
		return ""
	}

	return data["login"].(string)
}
