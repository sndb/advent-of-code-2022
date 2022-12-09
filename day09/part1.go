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

func delta(i, j int) int {
	if i < j {
		i, j = j, i
	}
	return i - j
}

func move(head, tail pos, d string) (pos, pos) {
	prev := head
	switch d {
	case "U":
		head.y++
	case "D":
		head.y--
	case "R":
		head.x++
	case "L":
		head.x--
	}
	if !adjacent(head, tail) {
		tail = prev
	}
	return head, tail
}

func adjacent(head, tail pos) bool {
	if delta(head.x, tail.x) < 2 && delta(head.y, tail.y) < 2 {
		return true
	}
	return false
}

func parse(line string) (string, int) {
	var d string
	var u int
	fmt.Sscanf(line, "%s %d", &d, &u)
	return d, u
}

func main() {
	var head, tail pos
	visited := map[pos]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		d, u := parse(line)

		for i := 0; i < u; i++ {
			head, tail = move(head, tail, d)
			visited[tail] = true
		}
	}
	fmt.Println(len(visited))
}
