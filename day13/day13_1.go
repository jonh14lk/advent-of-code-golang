package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	inf, ans := int(1e9), 0
	for {
		s, _ := scanner.ReadString('\n')
		s2, _ := scanner.ReadString('\n')
		s3, _ := scanner.ReadString('\n')
		var x1, x2, y1, y2, goalx, goaly int
		fmt.Sscanf(s[:len(s)-1], "Button A: X+%d, Y+%d", &x1, &y1)
		fmt.Sscanf(s2[:len(s2)-1], "Button B: X+%d, Y+%d", &x2, &y2)
		fmt.Sscanf(s3[:len(s3)-1], "Prize: X=%d, Y=%d", &goalx, &goaly)
		cost, t := inf, 0
		for (t*x1) <= goalx && (t*y1) <= goaly {
			difx, dify := goalx-(t*x1), goaly-(t*y1)
			if difx%x2 == 0 && dify%y2 == 0 && difx/x2 == dify/y2 {
				curr := (3 * t) + (difx / x2)
				if curr < cost {
					cost = curr
				}
			}
			t++
		}
		if cost < inf {
			ans += cost
		}
		_, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
	}
	fmt.Println(ans)
}
