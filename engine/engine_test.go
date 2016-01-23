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

func TestUpdate(t *testing.T) {
	tests := []struct {
		size      Loc
		init      map[Loc]Robot
		initRound int

		// TODO: parameters for round

		want      map[Loc]Robot
		wantRound int
	}{
		// TODO: This is no-op update change detector test.
		{
			size: Loc{5, 5},
			init: map[Loc]Robot{
				Loc{1, 1}: Robot{ID: 123, Health: 10, Faction: 0},
				Loc{2, 2}: Robot{ID: 456, Health: 10, Faction: 1},
			},
			initRound: 0,
			want: map[Loc]Robot{
				Loc{1, 1}: Robot{ID: 123, Health: 10, Faction: 0},
				Loc{2, 2}: Robot{ID: 456, Health: 10, Faction: 1},
			},
			wantRound: 1,
		},
	}
	for i, test := range tests {
		t.Logf("tests[%d], size = %v, round = %d", i, test.size, test.initRound)
		b := NewBoard(test.size.X, test.size.Y)
		b.Round = test.initRound
		for l, r := range test.init {
			t.Logf("  -> set %v to %#v", l, r)
			rr := new(Robot)
			*rr = r
			b.Set(l, rr)
		}

		t.Logf("  -> Update()")
		// TODO: add parameters
		b.Update()

		if b.Round != test.wantRound {
			t.Errorf("  !! b.Round = %d; want %d", b.Round, test.wantRound)
		}

		for y := 0; y < test.size.Y; y++ {
			for x := 0; x < test.size.X; x++ {
				loc := Loc{x, y}
				r := b.At(loc)
				want, ok := test.want[loc]
				if (r != nil) != ok {
					if ok {
						t.Errorf("  !! b.At(%v) = nil; want %#v", loc, want)
					} else {
						t.Errorf("  !! b.At(%v) = %#v; want nil", loc, r)
					}
					continue
				}
				if !ok {
					continue
				}
				if *r != want {
					t.Errorf("  !! b.At(%v) = %#v; want %#v", loc, r, want)
				}
			}
		}
	}
}
