package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	first  int64
	second int
}

var dp = make(map[pair]int64)

func dfs(curr pair) int64 {
	if val, exists := dp[curr]; exists {
		return val
	}
	ans := int64(0)
	str := strconv.FormatInt(curr.first, 10)
	if curr.second == 75 {
		return 1
	} else if curr.first == 0 {
		ans += dfs(pair{1, curr.second + 1})
	} else if len(str)%2 == 0 {
		div := len(str) / 2
		na, _ := strconv.ParseInt(str[:div], 10, 64)
		nb, _ := strconv.ParseInt(str[div:], 10, 64)
		ans += dfs(pair{na, curr.second + 1})
		ans += dfs(pair{nb, curr.second + 1})
	} else {
		ans += dfs(pair{curr.first * 2024, curr.second + 1})
	}
	dp[curr] = ans
	return ans
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	nums := strings.Split(s, " ")
	ans := int64(0)
	for _, ch := range nums {
		num, _ := strconv.ParseInt(ch, 10, 64)
		ans += dfs(pair{num, 0})
	}
	fmt.Println(ans)
}
