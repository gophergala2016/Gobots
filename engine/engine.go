package engine

import (
	"fmt"
	"strconv"

	"github.com/gophergala2016/Gobots/botapi"
)

type Board struct {
	cells []*Robot
	Size  Loc
	Round int

	nextID RobotID
}

// EmptyBoard creates an empty board of the given size.
func EmptyBoard(w, h int) *Board {
	return &Board{
		cells: make([]*Robot, w*h),
		Size:  Loc{w, h},
	}
}

// NewBoard creates an initialized game board for two factions.
func NewBoard(w, h int) *Board {
	// TODO: fill in initial conditions
	return EmptyBoard(w, h)
}

func (b *Board) Update() {
	// TODO
	b.Round++
}

// IsFinished reports whether the game is finished.
func (b *Board) IsFinished() bool {
	return b.Round >= 100
}

// At returns the robot at a location or nil if not found.
func (b *Board) At(loc Loc) *Robot {
	if !b.isValidLoc(loc) {
		panic("location out of bounds")
	}
	return b.cells[loc.Y*b.Size.X+loc.X]
}

// Set sets a robot at a particular location.
// Used for initialization.
func (b *Board) Set(loc Loc, r *Robot) {
	if !b.isValidLoc(loc) {
		panic("location out of bounds")
	}
	b.cells[loc.Y*b.Size.X+loc.X] = r
}

func (b *Board) isValidLoc(loc Loc) bool {
	return loc.X >= 0 && loc.X < b.Size.X && loc.Y >= 0 && loc.Y < b.Size.Y
}

// ToWire converts the board to the wire representation with respect to the
// given faction (since the wire factions are us vs. them).
func (b *Board) ToWire(out botapi.Board, faction int) error {
	out.SetWidth(uint16(b.Size.X))
	out.SetHeight(uint16(b.Size.Y))
	out.SetRound(int32(b.Round))

	n := 0
	for _, r := range b.cells {
		if r != nil {
			n++
		}
	}
	robots, err := botapi.NewRobot_List(out.Segment(), int32(n))
	if err != nil {
		return err
	}
	if err = out.SetRobots(robots); err != nil {
		return err
	}
	n = 0
	for i, r := range b.cells {
		if r == nil {
			continue
		}
		outr := robots.At(n)
		outr.SetId(uint32(r.ID))
		outr.SetX(uint16(i % b.Size.X))
		outr.SetY(uint16(i / b.Size.X))
		outr.SetHealth(int16(r.Health))
		if r.Faction == faction {
			outr.SetFaction(botapi.Faction_mine)
		} else {
			outr.SetFaction(botapi.Faction_opponent)
		}
		n++
	}
	return nil
}

// A Robot is a single piece on a board.
type Robot struct {
	ID      RobotID
	Health  int
	Faction int
}

type RobotID uint32

func (id RobotID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id RobotID) GoString() string {
	return id.String()
}

// Loc is a position on a board.
type Loc struct {
	X, Y int
}

func (loc Loc) String() string {
	return fmt.Sprintf("(%d, %d)", loc.X, loc.Y)
}
