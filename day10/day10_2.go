package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := [][]int{}
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		s := string(b)
		row := make([]int, len(s))
		for i, ch := range s {
			row[i], _ = strconv.Atoi(string(ch))
		}
		v = append(v, row)
	}
	n := len(v)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	ans := 0
	for x := 9; x >= 0; x-- {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if v[i][j] == x {
					if x == 9 {
						dp[i][j] = 1
					} else {
						for dir := 0; dir < 4; dir++ {
							x, y := i+dx[dir], j+dy[dir]
							if x >= 0 && x < n && y >= 0 && y < n && v[x][y]-1 == v[i][j] {
								dp[i][j] += dp[x][y]
							}
						}
					}
					if x == 0 {
						ans += dp[i][j]
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
