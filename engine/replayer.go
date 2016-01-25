package engine

import (
	"github.com/gophergala2016/Gobots/botapi"
)

type Replay struct {
	botapi.Replay
	round int
}

func NewReplay(r botapi.Replay) *Replay {
	return &Replay{r, 0}
}

func (r *Replay) NextBoard() *Board {
	// TODO: Stop ignoring errors
	if r.round == 0 {
		w, _ := r.InitialBoard()
		b, _ := boardFromWire(w)
		r.round++
		return b
	} else {
		rs, _ := r.Rounds()
		w, _ := rs.At(r.round).EndBoard()
		b, _ := boardFromWire(w)
		r.round++
		return b
	}
}

// boardFromWire converts the wire representation to the board
func boardFromWire(wire botapi.Board) (*Board, error) {
	b := EmptyBoard(int(wire.Width()), int(wire.Height()))
	b.Round = int(wire.Round())

	bots, err := wire.Robots()
	if err != nil {
		return b, err
	}

	for i := 0; i < bots.Len(); i++ {
		bot := bots.At(i)
		loc := Loc{
			X: int(bot.X()),
			Y: int(bot.Y()),
		}
		b.cells[b.cellIndex(loc)] = robotFromWire(bot)
	}

	return b, nil
}

func robotFromWire(wire botapi.Robot) *Robot {
	var faction int
	if wire.Faction() == botapi.Faction_mine {
		faction = P1Faction
	} else {
		faction = P2Faction
	}

	return &Robot{
		ID:      RobotID(wire.Id()),
		Health:  int(wire.Health()),
		Faction: faction,
	}
}
