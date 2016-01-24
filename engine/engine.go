package engine

import (
	"fmt"
	"strconv"

	"github.com/gophergala2016/Gobots/botapi"
)

// TODO: Fix the super inconsistent board API that alternates between taking in
// *Robots and RobotIDs, it's gross

const (
	P1Faction       = 1
	P2Faction       = 2
	InitialHealth   = 5
	CollisionDamage = 1
	AttackDamage    = 2
	DestructDamage  = 2
	SelfDamage      = 1000 // Make them super dead
)

type collisionMap map[Loc][]RobotID

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

func (b *Board) Width() int {
	return b.Size.X
}

func (b *Board) Height() int {
	return b.Size.Y
}

func (b *Board) newID() RobotID {
	b.nextID++
	return b.nextID
}

// NewBoard creates an initialized game board for two factions.
func NewBoard(w, h int) *Board {
	b := EmptyBoard(w, h)

	// Just line the ends with robots
	for i := 0; i < h; i++ {
		la, lb := Loc{0, i}, Loc{w - 1, i}
		ca, cb := b.cellIndex(la), b.cellIndex(lb)
		b.cells[ca] = &Robot{
			ID:      b.newID(),
			Health:  InitialHealth,
			Faction: P1Faction,
		}
		b.cells[cb] = &Robot{
			ID:      b.newID(),
			Health:  InitialHealth,
			Faction: P1Faction,
		}
	}
	return b
}

func (b *Board) Update(ta, tb botapi.Turn_List) {
	c := make(collisionMap)
	b.addCollisions(c, ta)
	b.addCollisions(c, tb)

	// Move the bots to their new locations, unless they collide with something,
	// in which case just subtract 1 from their health and don't move them.

	for loc, botIDs := range c {
		// If there's only one bot trying to get somewhere, just move them there
		if len(botIDs) == 1 {
			b.moveBot(botIDs[0], loc)
		}

		// Multiple bots, hurt 'em
		for _, id := range botIDs {
			// TODO: nil check, for safety
			bot := b.robot(id)
			b.hurtBot(bot, CollisionDamage)
		}
	}
	// Get rid of anyone who died in a collision
	b.clearTheDead()

	// Ok, we've moved everyone into place and hurt them for bumping into each
	// other, now we issue attacks
	// We issue attacks first, because I don't like the idea of robots
	// self-destructing when someone could have killed them, it makes for better
	// strategy this way

	// Allow all attacks to be issued before removing bots, because there's no
	// good, sensical way to order attacks. They all happen simultaneously
	b.issueAttacks(ta)
	b.issueAttacks(tb)

	// Get rid of anyone who was viciously murdered
	b.clearTheDead()

	// Boom goes the dynamite
	b.issueSelfDestructs(ta)
	b.issueSelfDestructs(tb)

	// Get rid of anyone killed in some kamikaze-shenanigans
	b.clearTheDead()

	b.Round++
}

func (b *Board) issueAttacks(ts botapi.Turn_List) {
	for i := 0; i < ts.Len(); i++ {
		t := ts.At(i)
		if t.Which() != botapi.Turn_Which_attack {
			continue
		}

		// They're attacking
		loc := b.robotLoc(RobotID(t.Id()))
		xOff, yOff := directionOffsets(t.Attack())
		attackLoc := Loc{
			X: loc.X + xOff,
			Y: loc.Y + yOff,
		}

		// If there's a bot at the attack location, make them sad
		// NOTE: You *can* hurt attack your own robots
		victim := b.At(attackLoc)
		if victim != nil {
			b.hurtBot(victim, AttackDamage)
		}
	}
}

func (b *Board) issueSelfDestructs(ts botapi.Turn_List) {
	for i := 0; i < ts.Len(); i++ {
		t := ts.At(i)
		if t.Which() != botapi.Turn_Which_selfDestruct {
			continue
		}

		// They're Metro-booming on production:
		// (https://www.youtube.com/watch?v=NiM5ARaexPE)
		loc := b.robotLoc(RobotID(t.Id()))
		for _, boomLoc := range b.surrounding(loc) {
			// If there's a bot in the blast radius
			victim := b.At(boomLoc)
			if victim != nil {
				b.hurtBot(victim, DestructDamage)
			}
		}

		bomber := b.At(loc)
		// Kill 'em
		b.hurtBot(bomber, SelfDamage)
	}
}

func (b *Board) surrounding(loc Loc) []Loc {
	offs := []int{-1, 0, 1}

	// At most 8 surrounding locations
	vLocs := make([]Loc, 0, 8)
	for _, ox := range offs {
		for _, oy := range offs {
			// Skip the explosion location
			if ox == 0 && oy == 0 {
				continue
			}
			l := Loc{
				X: loc.X + ox,
				Y: loc.Y + oy,
			}
			if b.isValidLoc(l) {
				vLocs = append(vLocs, l)
			}
		}
	}
	return vLocs
}

func (b *Board) addCollisions(c collisionMap, ts botapi.Turn_List) {
	for i := 0; i < ts.Len(); i++ {
		t := ts.At(i)
		id := RobotID(t.Id())
		nextLoc := b.nextLoc(id, t)
		// Add where they want to move
		c[nextLoc] = append(c[nextLoc], id)
	}
}

func (b *Board) hurtBot(r *Robot, damage int) {
	r.Health -= damage
}

func (b *Board) clearTheDead() {
	for _, bot := range b.cells {
		if bot == nil {
			continue
		}

		// Smite them
		if bot.Health <= 0 {
			loc := b.robotLoc(bot.ID)
			ind := b.cellIndex(loc)
			b.cells[ind] = nil // BOoOoOM, roasted
		}
	}
}

// TODO: Maybe make sure they're not teleporting across the board
func (b *Board) moveBot(id RobotID, loc Loc) {
	oldLoc := b.robotLoc(id)
	ind := b.cellIndex(oldLoc)

	bot := b.cells[ind]
	b.cells[ind] = nil
	b.cells[b.cellIndex(loc)] = bot
}

func (b *Board) cellIndex(loc Loc) int {
	return loc.Y*b.Size.X + loc.X
}

func (b *Board) nextLoc(id RobotID, t botapi.Turn) Loc {
	currentLoc := b.robotLoc(id)
	// If they aren't moving, return their current loc
	if t.Which() != botapi.Turn_Which_move {
		return currentLoc
	}

	// They're moving, return where they're going

	xOff, yOff := directionOffsets(t.Move())
	nextLoc := Loc{
		X: currentLoc.X + xOff,
		Y: currentLoc.Y + yOff,
	}

	if b.isValidLoc(nextLoc) {
		return nextLoc
	}

	// TODO: Penalize people for creating incompetent bots that like travelling
	// to invalid locations, which is the case if we've reached here.
	return currentLoc
}

func directionOffsets(dir botapi.Direction) (x, y int) {
	var xOff, yOff int
	switch dir {
	case botapi.Direction_north:
		yOff = -1
	case botapi.Direction_south:
		yOff = 1
	case botapi.Direction_east:
		xOff = 1
	case botapi.Direction_west:
		xOff = -1
	}
	return xOff, yOff
}

// TODO: Jesus Christ this is inefficient, we should have a map from ids to
// locations for O(1) lookups, our turn algorithm is going to be like O(n^3),
// which doesn't matter for like 10 bots, but still.
func (b *Board) robotLoc(id RobotID) Loc {
	for i, c := range b.cells {
		if c == nil {
			continue
		}
		if c.ID != id {
			continue
		}

		return Loc{
			X: i % b.Size.X,
			Y: i / b.Size.X,
		}
	}
	return Loc{}
}

func (b *Board) robot(id RobotID) *Robot {
	for _, c := range b.cells {
		if c == nil {
			continue
		}
		if c.ID != id {
			continue
		}

		return c
	}
	return nil
}

// IsFinished reports whether the game is finished.
func (b *Board) IsFinished() bool {
	return b.Round >= 100
}

// At returns the robot at a location or nil if not found.
func (b *Board) At(loc Loc) *Robot {
	if !b.isValidLoc(loc) {
		// TODO: Is panic the right thing to do here?
		panic("location out of bounds")
	}
	return b.cells[loc.Y*b.Size.X+loc.X]
}

// At returns the robot at a location or nil if not found.
func (b *Board) AtXY(x, y int) *Robot {
	if !b.isValidLoc(Loc{x, y}) {
		// TODO: Is panic the right thing to do here?
		panic("location out of bounds")
	}
	return b.cells[y*b.Size.X+x]
}

// Set sets a robot at a particular location.
// Used for initialization.
func (b *Board) Set(loc Loc, r *Robot) {
	if !b.isValidLoc(loc) {
		// TODO: Is panic the right thing to do here?
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
