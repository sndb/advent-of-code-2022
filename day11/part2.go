package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items []int
	op    func(int) int
	div   int
	if1   int
	if0   int
}

func parse(input string) *monkey {
	r := &monkey{}
	lines := strings.Split(input, "\n")

	for _, item := range strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ") {
		n, _ := strconv.Atoi(item)
		r.items = append(r.items, n)
	}

	operation := strings.Split(strings.TrimPrefix(lines[2], "  Operation: new = old "), " ")
	arg, err := strconv.Atoi(operation[1])
	r.op = func(old int) int {
		if err != nil {
			arg = old
		}
		if operation[0] == "+" {
			return old + arg
		}
		return old * arg
	}

	r.div, _ = strconv.Atoi(strings.TrimPrefix(lines[3], "  Test: divisible by "))
	r.if1, _ = strconv.Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	r.if0, _ = strconv.Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))

	return r
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	mod := 1
	monkeys := []*monkey{}
	for i, m := range strings.Split(string(input), "\n\n") {
		monkeys = append(monkeys, parse(m))
		mod *= monkeys[i].div
	}
	inspects := make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		for j, m := range monkeys {
			for _, n := range m.items {
				inspects[j]++
				k := m.op(n) % mod
				var l *monkey
				if k%m.div == 0 {
					l = monkeys[m.if1]
				} else {
					l = monkeys[m.if0]
				}
				l.items = append(l.items, k)
			}
			m.items = nil
		}
	}
	sort.Ints(inspects)
	fmt.Println(inspects[len(inspects)-1] * inspects[len(inspects)-2])
}
