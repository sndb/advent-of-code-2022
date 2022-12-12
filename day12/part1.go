package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x int
	y int
}

type path struct {
	pos pos
	len int
}

func main() {
	grid := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	var src, dst pos
	for r, line := range grid {
		for c, char := range line {
			if char == 'S' {
				src = pos{r, c}
			} else if char == 'E' {
				dst = pos{r, c}
			}
		}
	}
	fmt.Println(find(grid, src, dst))
}

func find(grid []string, src, dst pos) (int, bool) {
	paths := []path{{src, 0}}
	for {
		if len(paths) == 0 {
			return 0, false
		}
		paths = update(grid, paths)
		for _, p := range paths {
			if p.pos == dst {
				return p.len, true
			}
		}
	}
}

func update(grid []string, paths []path) []path {
	next := []path{}
	for _, p := range paths {
		np := neighbors(grid, p)
		next = append(next, np...)
	}
	return next
}

var seen = map[pos]bool{}

func neighbors(grid []string, p path) []path {
	up := pos{p.pos.x - 1, p.pos.y}
	down := pos{p.pos.x + 1, p.pos.y}
	right := pos{p.pos.x, p.pos.y + 1}
	left := pos{p.pos.x, p.pos.y - 1}
	unvisited := []pos{}
	for _, v := range []pos{up, down, right, left} {
		unvisited = append(unvisited, v)
	}
	next := []path{}
	for _, v := range unvisited {
		if accessible(grid, p.pos, v) && !seen[v] {
			seen[v] = true
			next = append(next, path{v, p.len + 1})
		}
	}
	return next
}

func accessible(grid []string, src, dst pos) bool {
	return inside(grid, src) && inside(grid, dst) && distance(src, dst) <= 1 &&
		elevation(grid[dst.x][dst.y])-elevation(grid[src.x][src.y]) <= 1
}

func elevation(b byte) int {
	if b == 'S' {
		b = 'a'
	} else if b == 'E' {
		b = 'z'
	}
	return int(b - 'a')
}

func inside(grid []string, p pos) bool {
	return p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[0])
}

func distance(p, q pos) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
