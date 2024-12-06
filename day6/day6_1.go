package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := []string{}
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		v = append(v, string(b))
	}
	n := len(v)
	x, y := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v[i][j] == '^' {
				x, y = i, j
			}
		}
	}
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, n)
	}
	visited[x][y] = true
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	ans, dir := 1, 0
	for {
		out := false
		for {
			nxt_x, nxt_y := x+dx[dir], y+dy[dir]
			if nxt_x < 0 || nxt_x >= n || nxt_y < 0 || nxt_y >= n {
				out = true
				break
			}
			if v[nxt_x][nxt_y] != '#' {
				x, y = nxt_x, nxt_y
				break
			}
			dir = (dir + 1) % 4
		}
		if out {
			break
		}
		if !visited[x][y] {
			visited[x][y] = true
			ans++
		}
	}
	fmt.Println(ans)
}
