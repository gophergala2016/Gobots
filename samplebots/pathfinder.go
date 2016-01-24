package main

import "github.com/gophergala2016/Gobots/easyai"

type pathfinder struct {
	targets map[uint32]uint32
}

func (pf *pathfinder) RobotTick(b *easyai.Board, r *easyai.Robot) easyai.Turn {
	// Immediate surrounding attacks
	ds := []easyai.Direction{
		easyai.North,
		easyai.South,
		easyai.East,
		easyai.West,
	}
	for _, d := range ds {
		loc := r.Loc.Add(d)
		if opponentAt(b, loc) {
			return easyai.Turn{
				Kind:      easyai.Attack,
				Direction: d,
			}
		}
	}

	// Acquire target
	tgt, ok := pf.targets[r.ID]
	var opp *easyai.Robot
	if ok {
		opp = b.Find(func(q *easyai.Robot) bool {
			return q.ID == tgt
		})
	}
	if !ok || opp == nil {
		if pf.targets == nil {
			pf.targets = make(map[uint32]uint32)
		}
		opp := nearestOpponent(b, r.Loc)
		if opp == nil {
			return easyai.Turn{Kind: easyai.Wait}
		}
		pf.targets[r.ID] = opp.ID
	}

	// Move to target.
	// Don't worry about collisions, since we already shot at all neighbors.
	// TODO: but what about friends?
	// TODO: and why don't you compute the vector angle?
	switch {
	case opp.Loc.X < r.Loc.X:
		return easyai.Turn{
			Kind:      easyai.Move,
			Direction: easyai.West,
		}
	case opp.Loc.X > r.Loc.X:
		return easyai.Turn{
			Kind:      easyai.Move,
			Direction: easyai.East,
		}
	case opp.Loc.Y < r.Loc.Y:
		return easyai.Turn{
			Kind:      easyai.Move,
			Direction: easyai.North,
		}
	case opp.Loc.Y > r.Loc.Y:
		return easyai.Turn{
			Kind:      easyai.Move,
			Direction: easyai.South,
		}
	}
	// TODO: impossibru?
	return easyai.Turn{Kind: easyai.Wait}
}

func nearestOpponent(b *easyai.Board, loc easyai.Loc) *easyai.Robot {
	// Probably faster ways of doing this.. traversing outward
	var closest *easyai.Robot
	var closestDist int
	for y, row := range b.Cells {
		for x, r := range row {
			curr := easyai.Loc{x, y}
			if r == nil || r.Faction != easyai.OpponentFaction {
				continue
			}
			d := easyai.Distance(loc, curr)
			if closest == nil || d < closestDist {
				closest, closestDist = r, d
			}
		}
	}
	return closest
}
