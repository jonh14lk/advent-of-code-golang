package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var pairs [][]int
	n := 0
	for {
		b, _, _ := scanner.ReadLine()
		if len(b) == 0 {
			break
		}
		nums := strings.Split(string(b), "|")
		vec := make([]int, 2)
		for i := range 2 {
			vec[i], _ = strconv.Atoi(nums[i])
		}
		pairs = append(pairs, vec)
		n = int(math.Max(float64(n), float64(vec[0]+1)))
		n = int(math.Max(float64(n), float64(vec[1]+1)))
	}
	before := make([][]bool, n)
	for i := range n {
		before[i] = make([]bool, n)
	}
	for _, vec := range pairs {
		before[vec[0]][vec[1]] = true
	}
	ans := 0
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		nums := strings.Split(string(b), ",")
		vec := make([]int, len(nums))
		for i := range len(nums) {
			vec[i], _ = strconv.Atoi(nums[i])
		}
		ok := true
		for i := 0; i < len(vec); i++ {
			for j := i + 1; j < len(vec); j++ {
				if before[vec[j]][vec[i]] {
					ok = false
				}
			}
		}
		if !ok {
			for i := 0; i < len(vec); i++ {
				for j := i + 1; j < len(vec); j++ {
					if before[vec[j]][vec[i]] {
						vec[i], vec[j] = vec[j], vec[i]
					}
				}
			}
			ans += vec[len(vec)/2]
		}
	}
	fmt.Println(ans)
}
