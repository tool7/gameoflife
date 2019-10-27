package main

type patternType int
const (
	rPentomino patternType = iota + 1
	diehard
	acorn
	gosperGliderGun
)

type offset struct {
	x int
	y int
}

var rPentominoPattern = []offset {
	offset { 0, 0 },
	offset { 0, 1 },
	offset { 0, -1 },
	offset { -1, 0 },
	offset { 1, 1 },
}

var diehardPattern = []offset {
	offset { 2, 0 },
	offset { 3, 0 },
	offset { 4, 0 },
	offset { 3, 2 },
	offset { -2, 0 },
	offset { -2, 1 },
	offset { -3, 1 },
}

var acornPattern = []offset {
	offset { 0, 1 },
	offset { 1, 0 },
	offset { 2, 0 },
	offset { 3, 0 },
	offset { -2, 0 },
	offset { -3, 0 },
	offset { -2, 2 },
}

var gosperGliderGunPattern = []offset {
	offset { 0, 0 },
	offset { -4, 0 },
	offset { -4, 1 },
	offset { -4, -1 },
	offset { -3, 2 },
	offset { -3, -2 },
	offset { -2, 3 },
	offset { -2, -3 },
	offset { -1, 3 },
	offset { -1, -3 },
	offset { 1, 2 },
	offset { 1, -2 },
	offset { 2, 0 },
	offset { 2, 1 },
	offset { 2, -1 },
	offset { 3, 0 },
	offset { -13, 0 },
	offset { -14, 0 },
	offset { -13, 1 },
	offset { -14, 1 },
	offset { 6, 1 },
	offset { 6, 2 },
	offset { 6, 3 },
	offset { 7, 1 },
	offset { 7, 2 },
	offset { 7, 3 },
	offset { 8, 0 },
	offset { 8, 4 },
	offset { 10, 0 },
	offset { 10, -1 },
	offset { 10, 4 },
	offset { 10, 5 },
	offset { 20, 2 },
	offset { 20, 3 },
	offset { 21, 2 },
	offset { 21, 3 },
}

func getPatternOffsets(patternType patternType) []offset {
	switch patternType {
	case rPentomino:
		return rPentominoPattern
	case diehard:
		return diehardPattern
	case acorn:
		return acornPattern
	case gosperGliderGun:
		return gosperGliderGunPattern
	}

	return nil
}
