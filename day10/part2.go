package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(line string) (string, int) {
	var o string
	var v int
	fmt.Sscanf(line, "%s %d", &o, &v)
	return o, v
}

func main() {
	x, c := 1, 0
	cycle := func() {
		p := c % 40
		if p >= x-1 && p <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if c%40 == 39 {
			fmt.Println()
		}
		c++
	}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		op, v := parse(line)
		switch op {
		case "noop":
			cycle()
		case "addx":
			cycle()
			cycle()
			x += v
		}
	}
}
