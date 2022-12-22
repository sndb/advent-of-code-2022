package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Input format:
//
// .AB
// .C.
// ED.
// F..

type side [50][50]byte

const (
	sideA int = iota
	sideB
	sideC
	sideD
	sideE
	sideF
)

const (
	dirRight int = iota
	dirDown
	dirLeft
	dirUp
	dirN
)

type edge struct {
	side, a, b int
}

var edges = map[edge]edge{
	{sideA, 1, 2}: {sideF, 3, 2},
	{sideA, 2, 3}: {sideE, 3, 2},
	{sideA, 3, 4}: {sideC, 2, 1},
	{sideA, 1, 4}: {sideB, 2, 3},

	{sideB, 1, 2}: {sideF, 4, 3},
	{sideB, 2, 3}: {sideA, 1, 4},
	{sideB, 3, 4}: {sideC, 1, 4},
	{sideB, 1, 4}: {sideD, 4, 1},

	{sideC, 1, 2}: {sideA, 4, 3},
	{sideC, 2, 3}: {sideE, 2, 1},
	{sideC, 3, 4}: {sideD, 2, 1},
	{sideC, 1, 4}: {sideB, 3, 4},

	{sideD, 1, 2}: {sideC, 4, 3},
	{sideD, 2, 3}: {sideE, 1, 4},
	{sideD, 3, 4}: {sideF, 1, 4},
	{sideD, 1, 4}: {sideB, 4, 1},

	{sideE, 1, 2}: {sideC, 3, 2},
	{sideE, 2, 3}: {sideA, 3, 2},
	{sideE, 3, 4}: {sideF, 2, 1},
	{sideE, 1, 4}: {sideD, 2, 3},

	{sideF, 1, 2}: {sideE, 4, 3},
	{sideF, 2, 3}: {sideA, 2, 1},
	{sideF, 3, 4}: {sideB, 2, 1},
	{sideF, 1, 4}: {sideD, 3, 4},
}

var offsets = map[int]struct{ r, c int }{
	sideA: {0, 50},
	sideB: {0, 100},
	sideC: {50, 50},
	sideD: {100, 50},
	sideE: {100, 0},
	sideF: {150, 0},
}

func mod(a, b int) int {
	r := a % b
	if r < 0 {
		return b + r
	}
	return r
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parsePath(s string) []int {
	path := []int{}
	j := 0
	for i, r := range s {
		if r == 'R' || r == 'L' {
			path = append(path, atoi(s[j:i]))
			if r == 'R' {
				path = append(path, dirRight)
			} else {
				path = append(path, dirLeft)
			}
			j = i + 1
		}
	}
	return append(path, atoi(s[j:]))
}

func parseSide(lines []string, r, c int) side {
	s := side{}
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			s[i][j] = lines[i+r][j+c]
		}
	}
	return s
}

func parseSides(lines []string) map[int]side {
	return map[int]side{
		sideA: parseSide(lines, offsets[sideA].r, offsets[sideA].c),
		sideB: parseSide(lines, offsets[sideB].r, offsets[sideB].c),
		sideC: parseSide(lines, offsets[sideC].r, offsets[sideC].c),
		sideD: parseSide(lines, offsets[sideD].r, offsets[sideD].c),
		sideE: parseSide(lines, offsets[sideE].r, offsets[sideE].c),
		sideF: parseSide(lines, offsets[sideF].r, offsets[sideF].c),
	}
}

func parse() (map[int]side, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return parseSides(lines[:len(lines)-2]), parsePath(lines[len(lines)-1])
}

func move(sides map[int]side, r, c, side, dir int) (newr, newc, newside, newdir int) {
	newr, newc, newside, newdir, ok := wrap(sides, r, c, side, dir)
	if !ok {
		if dir == dirUp && sides[side][r-1][c] != '#' {
			newr--
		} else if dir == dirDown && sides[side][r+1][c] != '#' {
			newr++
		} else if dir == dirRight && sides[side][r][c+1] != '#' {
			newc++
		} else if dir == dirLeft && sides[side][r][c-1] != '#' {
			newc--
		}
	}
	return newr, newc, newside, newdir
}

func wrap(sides map[int]side, r, c, side, dir int) (newr, newc, newside, newdir int, ok bool) {
	var x, y, p int
	if r == 0 && dir == dirUp {
		x, y, p = 1, 2, c
	} else if r == 49 && dir == dirDown {
		x, y, p = 3, 4, c
	} else if c == 0 && dir == dirLeft {
		x, y, p = 2, 3, r
	} else if c == 49 && dir == dirRight {
		x, y, p = 1, 4, r
	} else {
		return r, c, side, dir, false
	}
	e := edges[edge{side, x, y}]
	if e.a > e.b {
		e.a, e.b = e.b, e.a
	}
	if e.a == x && e.b == y {
		p = 49 - p
	}

	if e.a == 1 && e.b == 2 {
		newr, newc, newdir = 0, p, dirDown
	} else if e.a == 2 && e.b == 3 {
		newr, newc, newdir = p, 0, dirRight
	} else if e.a == 3 && e.b == 4 {
		newr, newc, newdir = 49, p, dirUp
	} else if e.a == 1 && e.b == 4 {
		newr, newc, newdir = p, 49, dirLeft
	} else {
		panic(0)
	}
	if sides[e.side][newr][newc] != '#' {
		return newr, newc, e.side, newdir, true
	}
	return r, c, side, dir, true
}

func follow(sides map[int]side, path []int) (r, c, dir int) {
	r, c, dir = 0, 0, dirRight
	side := sideA
	for i, v := range path {
		if i%2 == 0 {
			for i := 0; i < v; i++ {
				r, c, side, dir = move(sides, r, c, side, dir)
			}
		} else {
			if v == dirRight {
				dir = mod(dir+1, dirN)
			} else {
				dir = mod(dir-1, dirN)
			}
		}
	}
	return offsets[side].r + r, offsets[side].c + c, dir
}

func main() {
	sides, path := parse()
	r, c, dir := follow(sides, path)
	fmt.Println((r+1)*1000 + (c+1)*4 + dir)
}
