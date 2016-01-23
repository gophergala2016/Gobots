package engine

import (
	"fmt"
	"strconv"
)

type Board struct {
	cells []*Robot
	Size  Loc
	Round int

	nextID RobotID
}

// NewBoard creates an empty board of the given size.
func NewBoard(w, h int) *Board {
	return &Board{
		cells: make([]*Robot, w*h),
		Size:  Loc{w, h},
	}
}

func (b *Board) Update() {
	// TODO
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
