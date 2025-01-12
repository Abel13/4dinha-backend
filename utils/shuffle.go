package utils

import (
	"math/rand"
	"time"
)

func Shuffle(deck []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

func CalculateGroup(n int) int {
	pos := (n-1)%14 + 1
	switch pos {
	case 1:
		return 1
	case 2, 14:
		return 2
	case 3, 13:
		return 3
	case 4, 12:
		return 4
	case 5, 11:
		return 5
	case 6, 10:
		return 6
	case 7, 9:
		return 7
	case 8:
		return 8
	default:
		return 1
	}
}
