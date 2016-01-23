package main

import (
	"net/http"
)

// I'm sure I'll need this eventually, and am definitely not prematurely
// optimizing.
type context struct {
	r *http.Request
	w http.ResponseWriter

	p *player
}

func newContext(w http.ResponseWriter, r *http.Request) context {
	return context{
		w: w,
		r: r,
		p: nil,
	}
}

type player struct {
	name      string
	endpoints []string
}
