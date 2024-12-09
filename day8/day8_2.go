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
	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for i2 := i; i2 < n; i2++ {
				ini := 0
				if i2 == i {
					ini = j + 1
				}
				for j2 := ini; j2 < n; j2++ {
					if (v[i][j] != '.') && (v[i][j] == v[i2][j2]) {
						di := i2 - i
						dj := j2 - j
						x := i2
						y := j2
						for x >= 0 && x < n && y >= 0 && y < n {
							vis[x][y] = 1
							x += di
							y += dj
						}
						x = i
						y = j
						for x >= 0 && x < n && y >= 0 && y < n {
							vis[x][y] = 1
							x -= di
							y -= dj
						}
					}
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += vis[i][j]
		}
	}
	fmt.Println(ans)
}
