package main

import "testing"

func TestMinTrials(t *testing.T) {
	cases := []struct {
		eggs, floors, want int
	}{
		{1, 10, 10},  // one egg: pure linear scan
		{2, 10, 4},   // 4 moves cover 4+3+2+1 = 10 floors
		{2, 100, 14}, // the famous answer
		{2, 36, 8},
		{3, 14, 4}, // C(4,1)+C(4,2)+C(4,3) = 14
		{5, 1, 1},
		{2, 0, 0},
		{0, 5, -1},
	}
	for _, c := range cases {
		if got := MinTrials(c.eggs, c.floors); got != c.want {
			t.Errorf("MinTrials(%d,%d) = %d, want %d", c.eggs, c.floors, got, c.want)
		}
	}
}
