package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type pos struct {
	row int
	col int
}

func readLines(r io.Reader) []string {
	s := bufio.NewScanner(r)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func main() {
	grid := readLines(os.Stdin)
	cmax := len(grid[0])
	rmax := len(grid)
	visible := map[pos]bool{}

	const prev0 = '0' - 1
	var prev byte

	for r := 0; r < rmax; r++ {
		prev = prev0
		for c := 0; c < cmax; c++ {
			if grid[r][c] > prev {
				visible[pos{r, c}] = true
				prev = grid[r][c]
			}
		}
		prev = prev0
		for c := cmax - 1; c >= 0; c-- {
			if grid[r][c] > prev {
				visible[pos{r, c}] = true
				prev = grid[r][c]
			}
		}
	}
	for c := 0; c < cmax; c++ {
		prev = prev0
		for r := 0; r < rmax; r++ {
			if grid[r][c] > prev {
				visible[pos{r, c}] = true
				prev = grid[r][c]
			}
		}
		prev = prev0
		for r := rmax - 1; r >= 0; r-- {
			if grid[r][c] > prev {
				visible[pos{r, c}] = true
				prev = grid[r][c]
			}
		}
	}
	fmt.Println(len(visible))
}
