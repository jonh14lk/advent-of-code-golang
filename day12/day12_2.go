package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func rotate90(v [][]bool) [][]bool {
	n := len(v)
	res := make([][]bool, n)
	for i := range n {
		res[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-i-1] = v[i][j]
		}
	}
	return res
}

type pair struct {
	first, second int
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := []string{}
	for {
		s, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		v = append(v, s[:len(s)-1])
	}
	n := len(v)
	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, n)
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] {
				n := len(v)
				on := make([][]bool, n)
				for i := range n {
					on[i] = make([]bool, n)
				}
				area := 1
				q := list.New()
				q.PushBack(pair{i, j})
				vis[i][j] = true
				for q.Len() > 0 {
					x := q.Front()
					q.Remove(x)
					pos, _ := x.Value.(pair)
					on[pos.first][pos.second] = true
					for dir := 0; dir < 4; dir++ {
						r := pos.first + dx[dir]
						c := pos.second + dy[dir]
						if r >= 0 && r < n && c >= 0 && c < n && v[r][c] == v[pos.first][pos.second] {
							if !vis[r][c] {
								area++
								vis[r][c] = true
								q.PushBack(pair{r, c})
							}
						}
					}
				}
				sides := 0
				for a := 0; a < 4; a++ {
					for i := n - 1; i >= 0; i-- {
						prv := false
						for j := 0; j < n; j++ {
							if on[i][j] && (i+1 == n || !on[i+1][j]) {
								if !prv {
									sides++
								}
								prv = true
							} else {
								prv = false
							}
						}
					}
					on = rotate90(on)
				}
				ans += sides * area
			}
		}
	}
	fmt.Println(ans)
}
