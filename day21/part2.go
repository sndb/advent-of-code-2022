package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op int

const (
	mul op = iota
	div
	add
	sub
	eq
)

type job any // expr | int | string

type expr struct {
	a, b job
	op
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func parse() map[string]job {
	scanner := bufio.NewScanner(os.Stdin)
	dict := map[string]job{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		name := strings.TrimSuffix(fields[0], ":")
		var j job
		if name == "humn" {
			j = "x"
		} else if len(fields) == 2 {
			j = atoi(fields[1])
		} else {
			e := expr{a: fields[1], b: fields[3]}
			if name == "root" {
				e.op = eq
			} else {
				switch fields[2] {
				case "*":
					e.op = mul
				case "/":
					e.op = div
				case "+":
					e.op = add
				case "-":
					e.op = sub
				}
			}
			j = e
		}
		dict[name] = j
	}
	return dict
}

func solve(dict map[string]job, j job) job {
	switch x := j.(type) {
	case string:
		return x
	case int:
		return x
	case expr:
		k, ok0 := x.a.(string)
		l, ok1 := x.b.(string)
		if !ok0 || !ok1 {
			return j
		}
		a, ok2 := dict[k].(int)
		b, ok3 := dict[l].(int)
		if !ok2 || !ok3 {
			return j
		}
		switch x.op {
		case mul:
			return a * b
		case div:
			return a / b
		case add:
			return a + b
		case sub:
			return a - b
		case eq:
			return j
		}
	}
	panic(j)
}

func print(dict map[string]job, j job) string {
	switch x := j.(type) {
	case string:
		if x == "x" {
			return "x"
		}
		return print(dict, dict[x])
	case int:
		return fmt.Sprintf("%d", x)
	case expr:
		var op string
		switch x.op {
		case mul:
			op = "*"
		case div:
			op = "/"
		case add:
			op = "+"
		case sub:
			op = "-"
		case eq:
			op = "="
		}
		return fmt.Sprintf("(%s %s %s)", print(dict, x.a), op, print(dict, x.b))
	}
	panic(j)
}

func main() {
	dict := parse()
	for i := 0; i < 1000; i++ {
		for k, v := range dict {
			dict[k] = solve(dict, v)
		}
	}
	// solve by hand
	fmt.Println(print(dict, dict["root"]))

}
