package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a1, b1, c1, a2, b2, c2 int64) (int64, int64) {
	det := a1*b2 - a2*b1
	if det == 0 {
		return -1, -1
	}
	x := (c1*b2 - c2*b1) / det
	y := (a1*c2 - a2*c1) / det
	return x, y
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	ans := int64(0)
	for {
		s, _ := scanner.ReadString('\n')
		s2, _ := scanner.ReadString('\n')
		s3, _ := scanner.ReadString('\n')
		var x1, x2, y1, y2, goalx, goaly int64
		fmt.Sscanf(s[:len(s)-1], "Button A: X+%d, Y+%d", &x1, &y1)
		fmt.Sscanf(s2[:len(s2)-1], "Button B: X+%d, Y+%d", &x2, &y2)
		fmt.Sscanf(s3[:len(s3)-1], "Prize: X=%d, Y=%d", &goalx, &goaly)
		goalx += 10000000000000
		goaly += 10000000000000
		a, b := solve(x1, x2, goalx, y1, y2, goaly)
		if (a*x1+b*x2) == goalx && (a*y1+b*y2) == goaly {
			ans += (a*3 + b)
		}
		_, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
	}
	fmt.Println(ans)
}
