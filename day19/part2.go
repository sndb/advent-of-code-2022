package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mineral int

const (
	ore mineral = iota
	clay
	obsidian
	geode

	mineralN
)

type blueprint [mineralN][mineralN]int

type state struct {
	resources [mineralN]int
	robots    [mineralN]int
	time      int
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func parse(s string) blueprint {
	fields := strings.Split(s, " ")
	return blueprint{
		ore:      {ore: atoi(fields[6])},
		clay:     {ore: atoi(fields[12])},
		obsidian: {ore: atoi(fields[18]), clay: atoi(fields[21])},
		geode:    {ore: atoi(fields[27]), obsidian: atoi(fields[30])},
	}
}

func genesis() state {
	return state{
		robots:    [mineralN]int{ore: 1},
		resources: [mineralN]int{},
		time:      32,
	}
}

func add(a, b [mineralN]int) [mineralN]int {
	for i, v := range b {
		a[i] += v
	}
	return a
}

func sub(a, b [mineralN]int) [mineralN]int {
	for i, v := range b {
		a[i] -= v
	}
	return a
}

func wait(s state) state {
	s.time--
	s.resources = add(s.resources, s.robots)
	return s
}

func buy(s state, bp blueprint, robot mineral) state {
	s.time--
	s.resources = sub(s.resources, bp[robot])
	s.resources = add(s.resources, s.robots)
	s.robots[robot]++
	return s
}

func enough(s state, bp blueprint, robot mineral) bool {
	for i, v := range bp[robot] {
		if s.resources[i]-v < 0 {
			return false
		}
	}
	return true
}

func maxspend(bp blueprint, m mineral) int {
	n := 0
	for _, cost := range bp {
		if cost[m] > n {
			n = cost[m]
		}
	}
	return n
}

func solve(bp blueprint) int {
	q := []state{genesis()}
	seen := map[state]bool{}
	ms := [mineralN]int{}
	for i := range ms {
		ms[i] = maxspend(bp, mineral(i))
	}
	max := 0
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		for i, v := range s.resources[:geode] {
			if v > s.time*ms[i]-s.robots[i]*(s.time-1) {
				s.resources[i] = s.time*ms[i] - s.robots[i]*(s.time-1)
			}
		}
		if seen[s] {
			continue
		}
		seen[s] = true
		if s.resources[geode] > max {
			max = s.resources[geode]
		}
		if s.time == 0 {
			continue
		}

		q = append(q, wait(s))
		if enough(s, bp, ore) && s.robots[ore] < ms[ore] {
			q = append(q, buy(s, bp, ore))
		}
		if enough(s, bp, clay) && s.robots[clay] < ms[clay] {
			q = append(q, buy(s, bp, clay))
		}
		if enough(s, bp, obsidian) && s.robots[obsidian] < ms[obsidian] {
			q = append(q, buy(s, bp, obsidian))
		}
		if enough(s, bp, geode) {
			q = append(q, buy(s, bp, geode))
		}
	}
	return max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	blueprints := []blueprint{}
	for scanner.Scan() {
		blueprints = append(blueprints, parse(scanner.Text()))
	}

	total := 1
	for _, bp := range blueprints[:3] {
		n := solve(bp)
		total *= n
	}
	fmt.Println(total)
}
