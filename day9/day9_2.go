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
	r := len(v) - 1
	for r >= 0 {
		if v[r] == -1 {
			r--
			continue
		}
		rr := r
		for rr-1 >= 0 && v[rr-1] == v[r] {
			rr--
		}
		sz := r - rr + 1
		cnt := 0
		for l := 0; l < rr; l++ {
			if v[l] == -1 {
				cnt++
			} else {
				cnt = 0
			}
			if cnt == sz {
				for x := 0; x < sz; x++ {
					v[l-x], v[r-x] = v[r-x], v[l-x]
				}
				break
			}
		}
		r = rr - 1
	}
	ans := int64(0)
	for pos, num := range v {
		if num != -1 {
			ans += int64(pos) * int64(num)
		}
	}
	fmt.Println(ans)
}
