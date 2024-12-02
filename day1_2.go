package main

import (
	"fmt"
)

func main() {
	a := []int{}
	b := make(map[int]int)
	for {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			break
		}
		a = append(a, x)
		b[y] += 1
	}
	sum := 0
	for _, x := range a {
		sum += (x * b[x])
	}
	fmt.Println(sum)
}
