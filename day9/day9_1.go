package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	id := 0
	v := []int{}
	for i, ch := range s {
		num, _ := strconv.Atoi(string(ch))
		if i%2 == 0 {
			for range num {
				v = append(v, id)
			}
			id++
		} else {
			for range num {
				v = append(v, -1)
			}
		}
	}
	l := 0
	for r := len(v) - 1; r >= 0; r-- {
		if v[r] == -1 {
			continue
		}
		for l < r && v[l] != -1 {
			l++
		}
		if l >= r {
			break
		}
		v[l], v[r] = v[r], v[l]
	}
	ans := int64(0)
	for pos, num := range v {
		if num != -1 {
			ans += int64(pos) * int64(num)
		}
	}
	fmt.Println(ans)
}
