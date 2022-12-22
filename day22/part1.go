package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	right int = iota
	down
	left
	up

	dirN
)

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
	var i int
	var r rune
	for i, r = range s {
		if r == 'R' || r == 'L' {
			n := atoi(s[j:i])
			path = append(path, n)
			if r == 'R' {
				path = append(path, right)
			} else {
				path = append(path, left)
			}
			j = i + 1
		}
	}
	path = append(path, atoi(s[j:]))
	return path
}

func parse() ([]string, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	max := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > max {
			max = len(line)
		}
		lines = append(lines, line)
	}
	for i, line := range lines[:len(lines)-2] {
		if len(line) < max {
			lines[i] = line + strings.Repeat(" ", max-len(line))
		}
	}
	return lines[:len(lines)-2], parsePath(lines[len(lines)-1])
}

func follow(grid []string, path []int) (int, int, int) {
	r, c, d := 0, 0, right
	for i, v := range grid[0] {
		if v == '.' {
			c = i
			break
		}
	}

	for i, v := range path {
		if i%2 != 0 {
			if v == right {
				d = mod(d+1, dirN)
			} else {
				d = mod(d-1, dirN)
			}
			continue
		}
		for i := 0; i < v; i++ {
			switch d {
			case right:
				cc := c + 1
				if cc >= len(grid[r]) || grid[r][cc] == ' ' {
					cc--
					for cc >= 0 && grid[r][cc] != ' ' {
						cc--
					}
					cc++
				}
				if grid[r][cc] == '#' {
					continue
				}
				c = cc
			case down:
				rr := r + 1
				if rr >= len(grid) || grid[rr][c] == ' ' {
					rr--
					for rr >= 0 && grid[rr][c] != ' ' {
						rr--
					}
					rr++
				}
				if grid[rr][c] == '#' {
					continue
				}
				r = rr
			case left:
				cc := c - 1
				if cc < 0 || grid[r][cc] == ' ' {
					cc++
					for cc < len(grid[r]) && grid[r][cc] != ' ' {
						cc++
					}
					cc--
				}
				if grid[r][cc] == '#' {
					continue
				}
				c = cc
			case up:
				rr := r - 1
				if rr < 0 || grid[rr][c] == ' ' {
					rr++
					for rr < len(grid) && grid[rr][c] != ' ' {
						rr++
					}
					rr--
				}
				if grid[rr][c] == '#' {
					continue
				}
				r = rr
			}
		}
	}
	return r, c, d
}

func main() {
	grid, path := parse()
	r, c, d := follow(grid, path)
	fmt.Println((r+1)*1000 + (c+1)*4 + d)
}
