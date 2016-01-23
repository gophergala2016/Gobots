package main

// TODO(bsprague): LITERALLY EVERYTHING. Off the top:

// - The server - It needs to be able to serve pages, handle matches, do OAuth
// with Github, and implement gRPCs bidirectional streaming

// - The library - Needs a base Bot that can be embedded in higher level bots,
// and needs to know how to handshake with the normal server

// - The game - Needs to exist. Should be in a separate subpackage?? In any
// case, gopherjs should be used so that the implementation only needs to be
// written once

func main() {
}
