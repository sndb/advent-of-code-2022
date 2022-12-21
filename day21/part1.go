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
)

type job any // expr | int

type expr struct {
	a, b string
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
		if len(fields) == 2 {
			j = atoi(fields[1])
		} else {
			e := expr{a: fields[1], b: fields[3]}
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
			j = e
		}
		dict[name] = j
	}
	return dict
}

func solve(dict map[string]job, j job) job {
	switch x := j.(type) {
	case int:
		return j
	case expr:
		a, ok := dict[x.a].(int)
		if !ok {
			return j
		}
		b, ok := dict[x.b].(int)
		if !ok {
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
		}
	}
	panic(j)
}

func main() {
	dict := parse()
	for {
		n, ok := dict["root"].(int)
		if ok {
			fmt.Println(n)
			return
		}
		for k, v := range dict {
			dict[k] = solve(dict, v)
		}
	}
}
