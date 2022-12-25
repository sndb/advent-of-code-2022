package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const digits = "=-012"

func snafuToDec(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n *= 5
		n += strings.IndexByte(digits, s[i]) - 2
	}
	return n
}

func decToSnafu(n int) string {
	s := ""
	for n > 0 {
		k := n % 5
		n /= 5
		if k < 3 {
			s = strconv.Itoa(k) + s
		} else {
			s = string(digits[k-3]) + s
			n++
		}
	}
	return s
}

func parse() []int {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	nums := []int{}
	for _, line := range strings.Fields(string(b)) {
		n := snafuToDec(line)
		nums = append(nums, n)
	}
	return nums
}

func main() {
	total := 0
	nums := parse()
	for _, n := range nums {
		total += n
	}
	fmt.Println(decToSnafu(total))
}
