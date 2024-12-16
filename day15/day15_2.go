package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const (
	wall     = -3
	position = -2
	empty    = -1
)

type pair struct {
	first, second int
}

func rotate90(v [][]int) [][]int {
	n := len(v)
	m := len(v[0])
	res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[j][n-i-1] = v[i][j]
		}
	}
	return res
}

func move(v [][]int) [][]int {
	n := len(v)
	m := len(v[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if v[i][j] == position {
				q := list.New()
				toChange := list.New()

				vis := make([][]bool, n)
				for x := range n {
					vis[x] = make([]bool, m)
				}

				q.PushBack(pair{i, j})
				vis[i][j] = true
				canChange := true

				for q.Len() > 0 {
					pos, _ := q.Front().Value.(pair)
					q.Remove(q.Front())
					toChange.PushBack(pos)
					if v[pos.first][pos.second+1] == wall {
						canChange = false
						break
					}
					if pos.first-1 >= 0 && v[pos.first][pos.second] == v[pos.first-1][pos.second] && !vis[pos.first-1][pos.second] {
						vis[pos.first-1][pos.second] = true
						q.PushBack(pair{pos.first - 1, pos.second})
					}
					if pos.first+1 < n && v[pos.first][pos.second] == v[pos.first+1][pos.second] && !vis[pos.first+1][pos.second] {
						vis[pos.first+1][pos.second] = true
						q.PushBack(pair{pos.first + 1, pos.second})
					}
					if v[pos.first][pos.second+1] != empty && !vis[pos.first][pos.second+1] {
						vis[pos.first][pos.second+1] = true
						q.PushBack(pair{pos.first, pos.second + 1})
					}
				}

				if canChange {
					for toChange.Len() > 0 {
						pos, _ := toChange.Back().Value.(pair)
						toChange.Remove(toChange.Back())
						v[pos.first][pos.second+1] = v[pos.first][pos.second]
						v[pos.first][pos.second] = empty
					}
					return v
				}
			}
		}
	}
	return v
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := [][]int{}
	obsId := 0
	for {
		s, _ := scanner.ReadString('\n')
		if len(s) == 1 {
			break
		}
		s = s[:len(s)-1]
		curr := make([]int, 2*len(s))
		for i, ch := range s {
			if ch == '#' {
				curr[2*i] = wall
				curr[2*i+1] = wall
			} else if ch == 'O' {
				curr[2*i] = obsId
				curr[2*i+1] = obsId
				obsId++
			} else if ch == '.' {
				curr[2*i] = empty
				curr[2*i+1] = empty
			} else {
				curr[2*i] = position
				curr[2*i+1] = empty
			}
		}
		v = append(v, curr)
	}
	n := len(v)
	m := len(v[0])
	for {
		moves, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		for _, ch := range moves {
			if ch == '\n' {
				break
			}
			if ch == '<' {
				v = rotate90(v)
				v = rotate90(v)
				v = move(v)
				v = rotate90(v)
				v = rotate90(v)
			} else if ch == '^' {
				v = rotate90(v)
				v = move(v)
				v = rotate90(v)
				v = rotate90(v)
				v = rotate90(v)
			} else if ch == '>' {
				v = move(v)
			} else if ch == 'v' {
				v = rotate90(v)
				v = rotate90(v)
				v = rotate90(v)
				v = move(v)
				v = rotate90(v)
			}
		}
	}
	ans := 0
	considered := make(map[int]bool)
	for i := range n {
		for j := range m {
			if v[i][j] >= 0 && !considered[v[i][j]] {
				ans += 100*i + j
				considered[v[i][j]] = true
			}
		}
	}
	fmt.Println(ans)
}
