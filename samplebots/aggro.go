package main

import "github.com/gophergala2016/Gobots/easyai"

type aggro struct{}

func (aggro) RobotTick(b *easyai.Board, r *easyai.Robot) easyai.Turn {
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
	return easyai.Turn{Kind: easyai.Wait}
}

func opponentAt(b *easyai.Board, loc easyai.Loc) bool {
	if b.IsInside(loc) {
		return false
	}
	r := b.At(loc)
	if r == nil {
		return false
	}
	return r.Faction == easyai.OpponentFaction
}
