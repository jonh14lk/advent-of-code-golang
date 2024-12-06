package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := [][]rune{}
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		v = append(v, []rune(string(b)))
	}
	n := len(v)
	inix, iniy := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v[i][j] == '^' {
				inix, iniy = i, j
			}
		}
	}
	ans := 0
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v[i][j] == '.' {
				v[i][j] = '#'
				iters := 0
				x, y := inix, iniy
				dir := 0
				for {
					iters++
					if iters > (5 * n * n) {
						ans++
						break
					}
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
				}
				v[i][j] = '.'
			}
		}
	}
	fmt.Println(ans)
}
