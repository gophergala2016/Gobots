// Package easyai provides an idiomatic Go wrapper around the bot API.
package easyai

import (
	"net"

	"github.com/gophergala2016/Gobots/botapi"
	"golang.org/x/net/context"
	"zombiezen.com/go/capnproto2/rpc"
)

// Board represents the state of the board in a round.
type Board struct {
	Round int
	Cells [][]*Robot
}

// At returns the robot at a particular cell or nil if none is present.
func (b *Board) At(loc Loc) *Robot {
	return b.Cells[loc.Y][loc.X]
}

// IsInside reports whether loc is inside the board bounds.
func (b *Board) IsInside(loc Loc) bool {
	return loc.X >= 0 && loc.X < len(b.Cells[0]) && loc.Y >= 0 && loc.Y < len(b.Cells)
}

// Find finds a robot on the board that matches the given function.
func (b *Board) Find(f func(*Robot) bool) *Robot {
	for _, row := range b.Cells {
		for _, r := range row {
			if f(r) {
				return r
			}
		}
	}
	return nil
}

// A Robot is a piece on the board.
type Robot struct {
	ID      uint32
	Loc     Loc
	Faction Faction
	Health  int
}

// Faction identifies who owns a robot.
type Faction int

const (
	MyFaction Faction = iota
	OpponentFaction
)

// An AI is an algorithm that makes moves for a particular game.
type AI interface {
	RobotTick(board *Board, r *Robot) Turn
}

// Loc is a coordinate pair.
type Loc struct {
	X, Y int
}

func (loc Loc) Add(d Direction) Loc {
	switch d {
	case North:
		return Loc{X: loc.X, Y: loc.Y - 1}
	case South:
		return Loc{X: loc.X, Y: loc.Y + 1}
	case West:
		return Loc{X: loc.X - 1, Y: loc.Y}
	case East:
		return Loc{X: loc.X + 1, Y: loc.Y}
	default:
		return loc
	}
}

// Distance returns the Manhattan distance between two locations.
func Distance(a, b Loc) int {
	dx := a.X - b.X
	if dx < 0 {
		dx = -dx
	}
	dy := a.Y - b.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

// A Turn represents what a robot will do.  The zero value waits the turn.
type Turn struct {
	Kind      TurnKind
	Direction Direction
}

func (t Turn) toWire(id uint32, wire botapi.Turn) {
	wire.SetId(id)
	switch t.Kind {
	case Wait:
		wire.SetWait()
	case Move:
		wire.SetMove(t.Direction.toWire())
	case Attack:
		wire.SetAttack(t.Direction.toWire())
	case SelfDestruct:
		wire.SetSelfDestruct()
	}
}

// TurnKind is an enumeration of the kinds of turns.
type TurnKind int

// Kinds of turns.
const (
	Wait TurnKind = iota
	Move
	Attack
	SelfDestruct
)

// Direction is a cardinal direction.
type Direction int

// The defined directions.
const (
	North = Direction(botapi.Direction_north)
	South = Direction(botapi.Direction_south)
	East  = Direction(botapi.Direction_east)
	West  = Direction(botapi.Direction_west)
)

func (d Direction) toWire() botapi.Direction {
	return botapi.Direction(d)
}

// Client represents a connection to the game server.
type Client struct {
	conn      *rpc.Conn
	connector botapi.AiConnector
}

// Dial connects to a server at the given TCP address.
func Dial(addr string) (*Client, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn := rpc.NewConn(rpc.StreamTransport(c))
	return &Client{
		conn:      conn,
		connector: botapi.AiConnector{Client: conn.Bootstrap(context.TODO())},
	}, nil
}

// Close terminates the connection to the server.
func (c *Client) Close() error {
	return c.conn.Close()
}

// RegisterAI adds an AI implementation for the token given by the website.
// The AI factory function will be called for each new game encountered.
func (c *Client) RegisterAI(token string, factory Factory) error {
	a := botapi.Ai_ServerToClient(&aiAdapter{
		factory: factory,
		games:   make(map[string]AI),
	})
	_, err := c.connector.Connect(context.TODO(), func(r botapi.ConnectRequest) error {
		creds, err := r.NewCredentials()
		if err != nil {
			return err
		}
		err = creds.SetSecretToken(token)
		if err != nil {
			return err
		}
		r.SetAi(a)
		return nil
	}).Struct()
	return err
}

// Factory is a function that creates an AI per game.
type Factory func(gameID string) AI

// aiAdapter is a type that implements botapi.Ai by mapping turns to
// games and calling the AI interface methods.
//
// Note: since TakeTurn does not call server.Ack, the Cap'n Proto
// concurrency model guarantees that each call to TakeTurn happens after
// the previous return. Thus, we don't need to add any additional locks.
type aiAdapter struct {
	factory Factory
	games   map[string]AI
}

func (a *aiAdapter) TakeTurn(call botapi.Ai_takeTurn) error {
	board, err := call.Params.Board()
	if err != nil {
		return err
	}
	gameID, err := board.GameId()
	if err != nil {
		return err
	}
	ai := a.games[gameID]
	if ai == nil {
		ai = a.factory(gameID)
		a.games[gameID] = ai
	}

	b, robots, err := convertBoard(board)
	if err != nil {
		return err
	}
	turns, err := botapi.NewTurn_List(call.Results.Segment(), int32(len(robots)))
	if err != nil {
		return err
	}
	for i, r := range robots {
		t := ai.RobotTick(b, r)
		t.toWire(r.ID, turns.At(i))
	}
	call.Results.SetTurns(turns)
	return nil
}

func convertBoard(wire botapi.Board) (b *Board, playerBots []*Robot, err error) {
	w, h := int(wire.Width()), int(wire.Height())
	cells := make([]*Robot, w*h)
	rows := make([][]*Robot, h)
	for y := range rows {
		rows[y] = cells[y*w : (y+1)*w]
	}
	robots, err := wire.Robots()
	if err != nil {
		return nil, nil, err
	}
	playerBots = make([]*Robot, 0, robots.Len())
	for i, n := 0, robots.Len(); i < n; i++ {
		r := robots.At(i)
		// TODO(light): check for negative (x,y)
		rr := &Robot{
			ID:     r.Id(),
			Loc:    Loc{int(r.X()), int(r.Y())},
			Health: int(r.Health()),
		}
		switch r.Faction() {
		case botapi.Faction_mine:
			rr.Faction = MyFaction
			playerBots = append(playerBots, rr)
		case botapi.Faction_opponent:
			fallthrough
		default:
			rr.Faction = OpponentFaction
		}
		rows[rr.Loc.Y][rr.Loc.X] = rr
	}
	return &Board{
		Round: int(wire.Round()),
		Cells: rows,
	}, playerBots, nil
}
