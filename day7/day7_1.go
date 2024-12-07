package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	ans := int64(0)
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		s := strings.Split(string(b), ": ")
		s2 := strings.Split(s[1], " ")
		result, _ := strconv.ParseInt(s[0], 10, 64)
		n := len(s2)
		nums := make([]int64, n)
		for i, val := range s2 {
			nums[i], _ = strconv.ParseInt(val, 10, 64)
		}
		for mask := 0; mask < (1 << (n - 1)); mask++ {
			curr := nums[0]
			for bit := 0; bit < (n - 1); bit++ {
				if mask&(1<<bit) > 0 {
					curr *= nums[bit+1]
				} else {
					curr += nums[bit+1]
				}
				if curr > result {
					break
				}
			}
			if curr == result {
				ans += result
				break
			}
		}
	}
	fmt.Println(ans)
}
