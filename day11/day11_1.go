package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	first  int64
	second int
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	nums := strings.Split(s, " ")
	ans := int64(0)
	for _, ch := range nums {
		q := list.New()
		num, _ := strconv.ParseInt(ch, 10, 64)
		q.PushBack(pair{num, 0})
		for q.Len() > 0 {
			front := q.Front()
			q.Remove(front)
			curr, _ := front.Value.(pair)
			str := strconv.FormatInt(curr.first, 10)
			if curr.second == 25 {
				ans++
			} else if curr.first == 0 {
				q.PushBack(pair{1, curr.second + 1})
			} else if len(str)%2 == 0 {
				div := len(str) / 2
				na, _ := strconv.ParseInt(str[:div], 10, 64)
				nb, _ := strconv.ParseInt(str[div:], 10, 64)
				q.PushBack(pair{na, curr.second + 1})
				q.PushBack(pair{nb, curr.second + 1})
			} else {
				q.PushBack(pair{curr.first * 2024, curr.second + 1})
			}
		}
	}
	fmt.Println(ans)
}
