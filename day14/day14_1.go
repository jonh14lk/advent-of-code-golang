package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	q1, q2, q3, q4 := 0, 0, 0, 0
	n, m := 101, 103
	for {
		s, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		var x, y, dx, dy int
		fmt.Sscanf(s[:len(s)-1], "p=%d,%d v=%d,%d", &x, &y, &dx, &dy)
		x += 100 * dx
		y += 100 * dy
		x = ((x % n) + n) % n
		y = ((y % m) + m) % m

		if x < (n/2) && y < (m/2) {
			q1++
		} else if x > (n/2) && y < (m/2) {
			q2++
		} else if x < (n/2) && y > (m/2) {
			q3++
		} else if x > (n/2) && y > (m/2) {
			q4++
		}
	}
	ans := q1 * q2 * q3 * q4
	fmt.Println(ans)
}
