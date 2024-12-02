package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{}
	b := []int{}
	for {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			break
		}
		a = append(a, x)
		b = append(b, y)
	}
	sort.Ints(a)
	sort.Ints(b)
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += int(math.Abs(float64(a[i] - b[i])))
	}
	fmt.Println(sum)
}
