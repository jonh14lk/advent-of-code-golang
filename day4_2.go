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
	ans := 0
	for t := 0; t < 4; t++ {
		for i := 0; i+2 < n; i++ {
			for j := 0; j+2 < n; j++ {
				if v[i][j] == 'M' && v[i][j+2] == 'S' && v[i+1][j+1] == 'A' && v[i+2][j] == 'M' && v[i+2][j+2] == 'S' {
					ans++
				}
			}
		}
		v = rotate90(v)
	}
	fmt.Println(ans)
}
