package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type stack []string

func (s *stack) isempty() bool {
	return len(*s) == 0
}
func (s *stack) push(str string) {

	*s = append(*s, str)

}
func (s *stack) pop() bool {
	if s.isempty() {
		return false
	} else {
		index := len(*s) - 1
		*s = (*s)[:index]
		return true
	}
}
func (s *stack) top() string {
	if s.isempty() {
		return ""
	} else {
		index := len(*s) - 1
		item := (*s)[index]
		return item

	}
}

func prec(s string) int {
	if s == "^" {
		return 3
	} else if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}
func infixtopostfix(expression string) string {
	var sta stack
	var postfix string
	for i := 0; i < len(expression); i++ {
		s := expression[i]
		if s == ' ' {
			continue
		} else if unicode.IsDigit(rune(s)) {
			var temp string
			for unicode.IsDigit(rune(expression[i])) {
				temp = temp + string(expression[i])
				i++
				if i < len(expression) {
					s = expression[i]
				} else {
					break
				}
			}

			i--
			postfix = postfix + temp + ","
		} else if s == '(' {
			sta.push(string(s))
		} else if expression[i] == ')' {
			for sta.top() != "(" {
				postfix = postfix + sta.top() + ","
				sta.pop()
			}
			sta.pop()
		} else {
			for !sta.isempty() && prec(string(s)) <= prec(sta.top()) {
				postfix = postfix + sta.top() + ","
				sta.pop()
			}
			sta.push(string(s))
		}
	}

	for !sta.isempty() {
		postfix = postfix + sta.top() + ","
		sta.pop()
	}
	return postfix
}

func postfixevaluation(expression string) string {
	var sta stack
	for i := 0; i < len(expression); i++ {

		if expression[i] == ',' {
			continue
		} else if unicode.IsDigit(rune(expression[i])) {
			var temp string
			if i < len(expression)-1 {
				for unicode.IsDigit(rune(expression[i])) {
					temp = temp + string(expression[i])
					i++
					if i >= len(expression) {
						break
					}
				}
				i--
			} else {
				temp = string(expression[i])
			}
			sta.push(temp)
		} else {
			val1 := sta.top()
			sta.pop()
			val2 := sta.top()
			sta.pop()
			a, err1 := strconv.Atoi(val1)
			b, err2 := strconv.Atoi(val2)

			if err1 == err2 {
				fmt.Print()
			}

			if expression[i] == '+' {
				res := b + a
				sta.push(strconv.Itoa(res))
			} else if expression[i] == '-' {
				res := b - a
				sta.push(strconv.Itoa(res))
			} else if expression[i] == '*' {
				res := b * a
				sta.push(strconv.Itoa(res))
			} else if expression[i] == '/' {
				res := b / a
				sta.push(strconv.Itoa(res))
			}

		}
	}
	res := sta.top()
	return res
}
func main() {
	//infix := "100 * ( 2 + 12 ) / 14"
	var infix string
	fmt.Scanln(&infix)
	postfix := infixtopostfix(infix)
	evaluation := postfixevaluation(postfix)
	fmt.Println(evaluation)

}
