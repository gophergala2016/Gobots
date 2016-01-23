package engine

import "testing"

func TestNewBoardIsEmpty(t *testing.T) {
	b := NewBoard(3, 5)
	if b.Size.X != 3 {
		t.Errorf("b.Size.X = %d; want 3", b.Size.X)
	}
	if b.Size.Y != 5 {
		t.Errorf("b.Size.Y = %d; want 5", b.Size.Y)
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 3; x++ {
			loc := Loc{x, y}
			if r := b.At(loc); r != nil {
				t.Errorf("b.At(%v) = %#v; want nil", loc, r)
			}
		}
	}
}

func TestBoard_Set(t *testing.T) {
	b := NewBoard(3, 5)
	loc := Loc{1, 2}
	b.Set(loc, &Robot{
		ID:      1234,
		Health:  50,
		Faction: 3,
	})
	if r := b.At(loc); r != nil {
		if r.ID != 1234 {
			t.Errorf("b.At(%v).ID = %d; want 1234", loc, r.ID)
		}
		if r.Health != 50 {
			t.Errorf("b.At(%v).Health = %d; want 50", loc, r.Health)
		}
		if r.Faction != 3 {
			t.Errorf("b.At(%v).Faction = %d; want 3", loc, r.Faction)
		}
	} else {
		t.Errorf("b.At(%v) = nil", r)
	}
}
