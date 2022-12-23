package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	a, b string
}

type state struct {
	time  int
	valve string
	mask  int
}

var (
	valves    = map[string]int{}            // valves[x] == flow rate of x
	rich      = []string{}                  // list of valves with flow rate > 0
	indices   = map[string]int{}            // indices[x] == index of valve x in rich list
	neighbors = map[string][]string{}       // neighbors[x] == neighbors of valve x
	distances = map[string]map[string]int{} // distances[x][y] == distance between valves x and y
	cache     = map[state]int{}
)

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func addDistance(a, b string, n int) {
	if distances[a] == nil {
		distances[a] = map[string]int{}
	}
	if distances[b] == nil {
		distances[b] = map[string]int{}
	}
	distances[a][b] = n
	distances[b][a] = n
}

func bfs() {
	for k, v := range valves {
		if v == 0 && k != "AA" {
			continue
		}
		queue := []string{k}
		seen := map[string]bool{k: true}
		parents := map[string]string{} // parents[x] == parent of valve x
		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			for _, m := range neighbors[n] {
				if seen[m] {
					continue
				}
				seen[m] = true
				queue = append(queue, m)
				parents[m] = n
			}
		}
		for p := range parents {
			if valves[p] == 0 {
				continue
			}
			i := 0
			for q := p; q != k; q = parents[q] {
				i++
			}
			addDistance(p, k, i)
		}
	}
}

func dfs(s state) int {
	if x, ok := cache[s]; ok {
		return x
	}
	max := 0
	for neighbor, distance := range distances[s.valve] {
		bit := 1 << indices[neighbor]
		if s.mask&bit > 0 {
			continue
		}
		rem := s.time - distance - 1
		if rem <= 0 {
			continue
		}
		x := dfs(state{rem, neighbor, s.mask | bit}) + valves[neighbor]*rem
		if x > max {
			max = x
		}
	}
	cache[s] = max
	return max
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		name := fields[1]
		rate := atoi(strings.TrimSuffix(strings.Split(fields[4], "=")[1], ";"))
		neigh := []string{}
		for _, v := range fields[9:] {
			neigh = append(neigh, strings.TrimSuffix(v, ","))
		}
		if rate > 0 {
			rich = append(rich, name)
		}
		valves[name] = rate
		neighbors[name] = neigh
	}
	for i, v := range rich {
		indices[v] = i
	}
}

func main() {
	parse()
	bfs()

	part1 := dfs(state{30, "AA", 0})
	part2 := 0

	bits := (1 << len(rich)) - 1
	for i := 0; i < bits+1; i++ {
		x := dfs(state{26, "AA", i}) + dfs(state{26, "AA", i ^ bits})
		if x > part2 {
			part2 = x
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
