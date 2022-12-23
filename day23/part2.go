package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	north int = iota
	south
	west
	east
)

type pos struct {
	r, c int
}

var directions = []int{north, south, west, east}

func rotate() {
	head := directions[0]
	directions = append(directions[1:], head)
}

func parse() map[pos]bool {
	grove := map[pos]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	j := 0
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			if v == '#' {
				grove[pos{j, i}] = true
			}
		}
		j++
	}
	return grove
}

func add(p, q pos) pos {
	p.r += q.r
	p.c += q.c
	return p
}

func occupied(grove map[pos]bool, p pos, d int) bool {
	switch d {
	case north:
		return grove[add(p, pos{-1, -1})] ||
			grove[add(p, pos{-1, 0})] ||
			grove[add(p, pos{-1, 1})]
	case south:
		return grove[add(p, pos{1, -1})] ||
			grove[add(p, pos{1, 0})] ||
			grove[add(p, pos{1, 1})]
	case west:
		return grove[add(p, pos{-1, -1})] ||
			grove[add(p, pos{0, -1})] ||
			grove[add(p, pos{1, -1})]
	case east:
		return grove[add(p, pos{-1, 1})] ||
			grove[add(p, pos{0, 1})] ||
			grove[add(p, pos{1, 1})]
	}
	panic("bad direction")
}

func adjacent(grove map[pos]bool, p pos) bool {
	return occupied(grove, p, north) ||
		occupied(grove, p, south) ||
		occupied(grove, p, west) ||
		occupied(grove, p, east)
}

func free(grove map[pos]bool, p pos) (pos, bool) {
	for _, d := range directions {
		if !occupied(grove, p, d) {
			switch d {
			case north:
				return add(p, pos{-1, 0}), true
			case south:
				return add(p, pos{1, 0}), true
			case west:
				return add(p, pos{0, -1}), true
			case east:
				return add(p, pos{0, 1}), true
			}
		}
	}
	return pos{0, 0}, false
}

func move(grove map[pos]bool) bool {
	propositions := map[pos][]pos{} // map from dst to list of src

	// 1st half
	for k := range grove {
		if adjacent(grove, k) {
			p, ok := free(grove, k)
			if !ok {
				continue
			}
			propositions[p] = append(propositions[p], k)
		}
	}
	rotate()

	// 2nd half
	ok := false
	for dst, srcs := range propositions {
		if len(srcs) == 1 {
			delete(grove, srcs[0])
			grove[dst] = true
			ok = true
		}
	}
	return ok
}

func main() {
	m := parse()
	i := 1
	for move(m) {
		i++
	}
	fmt.Println(i)
}
