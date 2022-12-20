package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	v    int
	prev *node
	next *node
}

func mod(x, y int) int {
	r := x % y
	if r < 0 {
		return y + r
	}
	return r
}

func left(n *node) {
	npp := n.prev.prev
	np := n.prev
	nn := n.next

	npp.next = n
	n.prev = npp

	n.next = np
	np.prev = n

	np.next = nn
	nn.prev = np
}

func right(n *node) {
	nnn := n.next.next
	nn := n.next
	np := n.prev

	n.next = nnn
	nnn.prev = n

	nn.next = n
	n.prev = nn

	np.next = nn
	nn.prev = np
}

func parse() []*node {
	scanner := bufio.NewScanner(os.Stdin)
	nums := []*node{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, &node{v: n})
	}
	for i := 0; i < len(nums); i++ {
		nums[i].prev = nums[mod(i-1, len(nums))]
		nums[i].next = nums[mod(i+1, len(nums))]
	}
	return nums
}

func main() {
	nums := parse()
	var zero *node
	for _, n := range nums {
		if n.v == 0 {
			zero = n
		}
		v := n.v
		if v < 0 {
			for v = -v; v > 0; v-- {
				left(n)
			}
		} else {
			for ; v > 0; v-- {
				right(n)
			}
		}
	}
	sum := 0
	for i := 0; i < 3000; i++ {
		zero = zero.next
		if i%1000 == 999 {
			sum += zero.v
		}
	}
	fmt.Println(sum)
}
