package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check2(a []int) bool {
	for i := 1; i < len(a); i++ {
		if (a[i]-a[i-1] < 1) || (a[i]-a[i-1] > 3) {
			return false
		}
	}
	return true
}

func check(a []int) bool {
	for i := 1; i < len(a); i++ {
		if (a[i]-a[i-1] < 1) || (a[i]-a[i-1] > 3) {
			return check2(erase_id(a, i)) || check2(erase_id(a, i-1))
		}
	}
	return true
}

func erase_id(a []int, i int) []int {
	b := make([]int, len(a)-1)
	curr_id := 0
	for id, x := range a {
		if id != i {
			b[curr_id] = x
			curr_id += 1
		}
	}
	return b
}

func reverse(a []int) []int {
	l := 0
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return a
}

func main() {
	ans := 0
	scanner := bufio.NewReader(os.Stdin)
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		a := strings.Split(string(b), " ")

		x := make([]int, len(a))
		for i, str := range a {
			num, _ := strconv.Atoi(str)
			x[i] = num
		}

		ok := check(x)
		ok = ok || check(reverse(x))
		if ok {
			ans += 1
		}
	}
	fmt.Println(ans)
}
