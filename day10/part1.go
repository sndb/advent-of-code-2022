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
	x, c, t := 1, 1, 0
	cycle := func() {
		if (c-20)%40 == 0 {
			t += x * c
		}
		c++
	}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		o, v := parse(line)
		switch o {
		case "noop":
			cycle()
		case "addx":
			cycle()
			cycle()
			x += v
		}
	}
	fmt.Println(t)
}
