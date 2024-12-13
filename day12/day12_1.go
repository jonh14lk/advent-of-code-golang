package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

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
				area := 1
				edges := 0
				q := list.New()
				q.PushBack(pair{i, j})
				vis[i][j] = true
				for q.Len() > 0 {
					x := q.Front()
					q.Remove(x)
					pos, _ := x.Value.(pair)
					for dir := 0; dir < 4; dir++ {
						r := pos.first + dx[dir]
						c := pos.second + dy[dir]
						if r >= 0 && r < n && c >= 0 && c < n && v[r][c] == v[pos.first][pos.second] {
							edges++
							if !vis[r][c] {
								area++
								vis[r][c] = true
								q.PushBack(pair{r, c})
							}
						}
					}
				}
				perimeter := (4 * area) - edges
				ans += perimeter * area
			}
		}
	}
	fmt.Println(ans)
}
