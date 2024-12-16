package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate90(v [][]rune) [][]rune {
	n := len(v)
	res := make([][]rune, n)
	for i := range n {
		res[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-i-1] = v[i][j]
		}
	}
	return res
}

func move(v [][]rune) [][]rune {
	n := len(v)
	m := len(v[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if v[i][j] == '@' {
				k := j + 1
				for {
					if v[i][k] == '#' || v[i][k] == '.' {
						break
					}
					k++
				}
				if v[i][k] == '.' {
					for k > j {
						v[i][k] = v[i][k-1]
						k--
					}
					v[i][j] = '.'
					break
				}
				return v
			}
		}
	}
	return v
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	v := [][]rune{}
	for {
		s, _ := scanner.ReadString('\n')
		if len(s) == 1 {
			break
		}
		v = append(v, []rune(s[:len(s)-1]))
	}
	n := len(v)
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
	for i := range n {
		for j := range n {
			if v[i][j] == 'O' {
				ans += 100*i + j
			}
		}
	}
	fmt.Println(ans)
}
