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

type knots [10]pos

func delta(i, j int) int {
	if i < j {
		i, j = j, i
	}
	return i - j
}

func keep(head, tail pos) pos {
	if head.x > tail.x {
		tail.x++
	}
	if head.x < tail.x {
		tail.x--
	}
	if head.y > tail.y {
		tail.y++
	}
	if head.y < tail.y {
		tail.y--
	}
	return tail
}

func move(k knots, d string) knots {
	switch d {
	case "U":
		k[0].y++
	case "D":
		k[0].y--
	case "R":
		k[0].x++
	case "L":
		k[0].x--
	}
	for i := 1; i < len(k); i++ {
		if !adjacent(k[i-1], k[i]) {
			k[i] = keep(k[i-1], k[i])
		}
	}
	return k
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
	var k knots
	visited := map[pos]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		d, u := parse(line)

		for i := 0; i < u; i++ {
			k = move(k, d)
			visited[k[len(k)-1]] = true
		}
	}
	fmt.Println(len(visited))
}
