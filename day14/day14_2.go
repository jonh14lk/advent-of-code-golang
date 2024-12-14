package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type pair struct {
	x, y int
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	n, m := 101, 103
	vec := []pair{}
	dir := []pair{}
	for {
		s, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		var x, y, dx, dy int
		fmt.Sscanf(s[:len(s)-1], "p=%d,%d v=%d,%d", &x, &y, &dx, &dy)
		vec = append(vec, pair{x, y})
		dir = append(dir, pair{dx, dy})
	}
	// 10000 is this a good limit????
	for id := range 10000 {
		vis := make([][]int, n)
		for i := range n {
			vis[i] = make([]int, m)
			for j := range m {
				vis[i][j] = 0
			}
		}
		for _, val := range vec {
			vis[val.x][val.y] = 1
		}

		img := image.NewRGBA(image.Rect(0, 0, n, m))

		for x := 0; x < n; x++ {
			cnt := 0
			for y := 0; y < m; y++ {
				if vis[x][y] == 1 {
					cnt++
					img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 0})
				} else {
					cnt = 0
					img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
				}
				if cnt > 10 {
					// guess: potential image to be the answer
					fmt.Println(id)
				}
			}
		}
		filename := "generated/output" + strconv.FormatInt(int64(id), 10) + ".png"
		outFile, _ := os.Create(filename)
		png.Encode(outFile, img)

		for i := range vec {
			vec[i].x += dir[i].x
			vec[i].y += dir[i].y
			vec[i].x = ((vec[i].x % n) + n) % n
			vec[i].y = ((vec[i].y % m) + m) % m
		}
	}
}
