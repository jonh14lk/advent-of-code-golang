package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	s := ""
	for {
		b, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		s += string(b)
	}
	i := 0
	expr := []string{}
	for i < len(s) {
		if i+2 < len(s) && s[i:i+3] == "mul" {
			expr = append(expr, s[i:i+3])
			i += 3
		} else if i+3 < len(s) && s[i:i+4] == "do()" {
			expr = append(expr, s[i:i+4])
			i += 4
		} else if i+6 < len(s) && s[i:i+7] == "don't()" {
			expr = append(expr, s[i:i+7])
			i += 7
		} else if s[i] >= '0' && s[i] <= '9' {
			curr_num := string(s[i])
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				curr_num += string(s[j])
				j += 1
			}
			expr = append(expr, curr_num)
			i = j

		} else {
			expr = append(expr, string(s[i]))
			i += 1
		}
	}
	ans := 0
	enabled := true
	for i := 0; i+5 < len(expr); i++ {
		if enabled && expr[i] == "mul" &&
			expr[i+1] == "(" &&
			expr[i+2][0] >= '0' && expr[i+2][0] <= '9' &&
			expr[i+3] == "," &&
			expr[i+4][0] >= '0' && expr[i+4][0] <= '9' &&
			expr[i+5] == ")" {
			numa, _ := strconv.Atoi(expr[i+2])
			numb, _ := strconv.Atoi(expr[i+4])
			ans += (numa * numb)
		} else if expr[i] == "do()" {
			enabled = true
		} else if expr[i] == "don't()" {
			enabled = false
		}
	}
	fmt.Println(ans)
}
