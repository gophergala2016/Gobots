package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/gophergala2016/Gobots/botapi"
	"github.com/gophergala2016/Gobots/engine"
	gocontext "golang.org/x/net/context"
	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/rpc"
)

type aiEndpoint struct {
	ds datastore

	// fields below are protected by mu
	mu     sync.Mutex
	online map[aiID]botapi.Ai
}

func startAIEndpoint(addr string) (*aiEndpoint, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	e := &aiEndpoint{online: make(map[aiID]botapi.Ai)}
	go e.listen(l)
	return e, nil
}

// listen runs in its own goroutine, listening for connections.
func (e *aiEndpoint) listen(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("ai endpoint: accept:", err)
			return
		}
		go e.handleConn(c)
	}
}

// handleConn runs in its own goroutine, started by listen.
func (e *aiEndpoint) handleConn(c net.Conn) {
	aic := &aiConnector{e: e}
	rc := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(botapi.AiConnector_ServerToClient(aic).Client))
	rc.Wait()
	aic.drop()
}

// TODO: list online AIs method

// connect adds an online AI, given the secret auth token.
func (e *aiEndpoint) connect(token string, ai botapi.Ai) (aiID, error) {
	// TODO
	return "", nil
}

// removeAIs drops AIs from online, usually via disconnection.
func (e *aiEndpoint) removeAIs(ids []aiID) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, i := range ids {
		delete(e.online, i)
	}
}

type aiConnector struct {
	e   *aiEndpoint
	ais []aiID
}

func (aic *aiConnector) Connect(call botapi.AiConnector_connect) error {
	creds, _ := call.Params.Credentials()
	tok, _ := creds.SecretToken()
	id, err := aic.e.connect(tok, call.Params.Ai())
	if err != nil {
		return err
	}
	aic.ais = append(aic.ais, id)
	return nil
}

func (aic *aiConnector) drop() {
	aic.e.removeAIs(aic.ais)
}

func runMatch(ctx gocontext.Context, ds datastore, aiA, aiB *onlineAI) error {
	// Create new board and store it.
	b := engine.NewBoard(20, 20)
	_, seg, _ := capnp.NewMessage(capnp.SingleSegment(nil))
	wb, _ := botapi.NewRootBoard(seg)
	b.ToWire(wb, 0)
	gid, err := ds.startGame(aiA.info.id, aiB.info.id, wb)
	if err != nil {
		return err
	}

	// Run the game
	for !b.IsFinished() {
		turnCtx, _ := gocontext.WithTimeout(ctx, 30*time.Second)
		chA, chB := make(chan turnResult), make(chan turnResult)
		go aiA.takeTurn(turnCtx, gid, b, 0, chA)
		go aiB.takeTurn(turnCtx, gid, b, 1, chB)
		ra, rb := <-chA, <-chB
		if ra.err.HasError() || rb.err.HasError() {
			// TODO: Something with errors
		}
		b.Update(ra.results, rb.results)
		// TODO: db.addRound
	}

	return nil
}

type onlineAI struct {
	info   aiInfo
	client botapi.Ai
}

type turnResult struct {
	results botapi.Turn_List
	err     turnError
}

func (oa *onlineAI) takeTurn(ctx gocontext.Context, gid gameID, b *engine.Board, faction int, ch chan<- turnResult) {
	results, err := oa.client.TakeTurn(ctx, func(p botapi.Ai_takeTurn_Params) error {
		wb, err := p.NewBoard()
		if err != nil {
			return err
		}
		wb.SetGameId(string(gid))
		return b.ToWire(wb, faction)
	}).Struct()
	var te turnError
	if err != nil {
		te = append(te, err)
	}

	tl, err := results.Turns()
	if err != nil {
		te = append(te, err)
	}
	ch <- turnResult{tl, te}
}

type turnError []error

func (t turnError) Error() string {
	var e string
	for _, err := range t {
		e += err.Error()
	}
	return e
}

func (t turnError) HasError() bool {
	return len(t) > 0
}
