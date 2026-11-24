package main

import "fmt"

// MinTrials returns the minimum number of trials that guarantees finding
// the critical floor in a building with the given floors, using at most
// the given eggs. Inverted DP: instead of "trials for f floors", compute
// "floors coverable with m moves and e eggs" and grow m until it suffices.
func MinTrials(eggs, floors int) int {
	if floors <= 0 {
		return 0
	}
	if eggs <= 0 {
		return -1
	}
	// cover[e] = floors decidable with the current move budget and e eggs.
	cover := make([]int, eggs+1)
	moves := 0
	for cover[eggs] < floors {
		moves++
		for e := eggs; e >= 1; e-- {
			// Drop once: if it breaks, e-1 eggs explore below (previous
			// budget); if not, e eggs explore above. +1 for this floor.
			cover[e] = cover[e] + cover[e-1] + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(MinTrials(2, 100))
	fmt.Println(MinTrials(3, 14))
}
