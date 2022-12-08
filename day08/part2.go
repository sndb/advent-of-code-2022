package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readLines(r io.Reader) []string {
	s := bufio.NewScanner(r)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func score(grid []string, r, c int) int {
	var s0, s1, s2, s3 int
	k := grid[r][c]
	for i := c - 1; i >= 0; i-- {
		s0++
		if grid[r][i] >= k {
			break
		}
	}
	for i := c + 1; i < len(grid[0]); i++ {
		s1++
		if grid[r][i] >= k {
			break
		}
	}
	for j := r - 1; j >= 0; j-- {
		s2++
		if grid[j][c] >= k {
			break
		}
	}
	for j := r + 1; j < len(grid); j++ {
		s3++
		if grid[j][c] >= k {
			break
		}
	}
	return s0 * s1 * s2 * s3
}

func main() {
	grid := readLines(os.Stdin)
	smax := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			s := score(grid, r, c)
			if s > smax {
				smax = s
			}
		}
	}
	fmt.Println(smax)
}
