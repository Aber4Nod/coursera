package main

import (
	"os"
	"fmt"
	"unicode/utf8"
	"bufio"
	"strconv"
)

func operCheck(cnt *uint8, chVal uint8) (bool, string) {
	if *cnt < chVal {
		str := "Invalid operation (drop expression)"
		*cnt = 0
		return false, str
	}
	*cnt--
	return true, ""
}

func calc(expr string) string {
	var sp uint8 = 0
	var str string; var err bool
	var runesNumber int = utf8.RuneCountInString(expr)
	stack := make([]int, runesNumber, runesNumber)
	for _, char := range expr {
		switch char {
		case ' ', '\n': break
		case '=':
			if err, str = operCheck(&sp, 1); err {
				return "Result = " + strconv.Itoa(stack[sp])
			}
			return str
		case '+':
			if err, str = operCheck(&sp, 2 ); err {
				stack[sp-1] = stack[sp-1] + stack[sp]; break
			}
			return str
		case '-':
			if err, str = operCheck(&sp, 2 ); err {
				stack[sp-1] = stack[sp-1] - stack[sp]; break
			}
			return str
		case '*':
			if err, str = operCheck(&sp, 2 ); err {
				stack[sp-1] = stack[sp-1] * stack[sp]; break
			}
			return str
		case '/':
			if err, str = operCheck(&sp, 2 ); err {
				stack[sp-1] = stack[sp-1] / stack[sp]; break
			}
			return str
		default:
			stack[sp] = int(char)-'0'; sp++
		}
	}
	if sp != 0 {
		return "Empty operation (Recoverable)"
	}
	return "Empty expression"
}

func main() {
	out := os.Stdout
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c := scanner.Text()
		res := calc(c)
		fmt.Fprintln(out, res)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}