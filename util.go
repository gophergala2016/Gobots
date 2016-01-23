package main

import "net/http"

func withLogin(handler func(c context)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO(bsprague): Make sure they're authenticated with
		// GitHub
		c := newContext(w, r)
		handler(c)
	}
}
