package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	r, c int
}

type state struct {
	minute int
	pos    pos
	stage  int
}

const (
	up int = iota
	down
	left
	right
)

var (
	maxR, maxC = 0, 0
	blizzards  = map[pos][]int{}
	cache      = map[int]map[pos][]int{}
)

func situation(minute int) map[pos][]int {
	if minute == 0 {
		return blizzards
	}
	if ret, ok := cache[minute]; ok {
		return ret
	}
	ret := map[pos][]int{}
	prev := situation(minute - 1)
	for p, dd := range prev {
		for _, d := range dd {
			var r, c int
			switch d {
			case up:
				r, c = p.r-1, p.c
			case down:
				r, c = p.r+1, p.c
			case left:
				r, c = p.r, p.c-1
			case right:
				r, c = p.r, p.c+1
			}
			switch {
			case r < 0:
				r = maxR
			case r > maxR:
				r = 0
			case c < 0:
				c = maxC
			case c > maxC:
				c = 0
			}
			ret[pos{r, c}] = append(ret[pos{r, c}], d)
		}
	}
	cache[minute] = ret
	return ret
}

func valid(p pos) bool {
	return (p.r == -1 && p.c == 0) || (p.r >= 0 && p.r <= maxR && p.c >= 0 && p.c <= maxC)
}

func positions(p pos) []pos {
	candidates := []pos{
		{p.r, p.c},
		{p.r + 1, p.c},
		{p.r - 1, p.c},
		{p.r, p.c + 1},
		{p.r, p.c - 1},
	}
	ret := []pos{}
	for _, c := range candidates {
		if valid(c) {
			ret = append(ret, c)
		}
	}
	return ret
}

func bfs(s state) int {
	queue := []state{s}
	seen := map[state]bool{}
	for {
		s, queue = queue[0], queue[1:]
		if seen[s] {
			continue
		}
		seen[s] = true

		if s.pos.r == maxR && s.pos.c == maxC {
			if s.stage == 2 {
				return s.minute + 1
			}
			if s.stage == 0 {
				s.stage = 1
				s.minute += 2
			}
		} else if s.pos.r == 0 && s.pos.c == 0 && s.stage == 1 {
			s.stage = 2
		}
		for _, p := range positions(s.pos) {
			next := state{s.minute + 1, p, s.stage}
			if _, occupied := situation(next.minute)[p]; !occupied {
				queue = append(queue, next)
			}
		}
	}
}

func parse() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	grid := strings.Fields(string(in))
	for r, row := range grid[1 : len(grid)-1] {
		maxR = r
		for c, dir := range row[1 : len(row)-1] {
			maxC = c
			switch dir {
			case '^':
				blizzards[pos{r, c}] = []int{up}
			case 'v':
				blizzards[pos{r, c}] = []int{down}
			case '<':
				blizzards[pos{r, c}] = []int{left}
			case '>':
				blizzards[pos{r, c}] = []int{right}
			}
		}
	}
}

func main() {
	parse()
	fmt.Println(bfs(state{0, pos{-1, 0}, 0}))
}
