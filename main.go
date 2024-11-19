package main

import (
	"MathEvaluator/stack"
	"fmt"
	"log"
	"regexp"
)

var priorities = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func main() {
	expression := "(1 + 2) * (3 - 4) / 5"
	fmt.Println(GetTokens(expression))
	fmt.Println(Calculate(GetTokens(expression)))
}

func GetTokens(expression string) (matched []string) {
	re, err := regexp.Compile("[()]|[0-9]|[+\\-*/]|[()]")
	if err != nil {
		log.Fatal(err)
	}
	matched = re.FindAllString(expression, -1)
	return
}

func Calculate(source []string) string {
	var res string
	st := stack.NewStack()
	for _, ch := range source {
		switch ch {
		case "+", "-", "*", "/":
			for peek, err := st.Peek(); !err || priorities[peek] > priorities[ch]; peek, err = st.Peek() {
				pop, err2 := st.Pop()
				if !err2 {
					break
				}
				res += pop
			}
			st.Push(ch)
		case "(":
			st.Push(ch)
		case ")":
			for peek, err := st.Peek(); !err || peek != "("; peek, err = st.Peek() {
				pop, err2 := st.Pop()
				if !err2 {
					break
				}
				res += pop
			}
		default:
			res += ch
		}
	}

	for _, err := st.Peek(); !err; _, err = st.Peek() {
		pop, err2 := st.Pop()
		if !err2 {
			break
		}
		res += pop
	}

	return res
}

// st := stack.NewStack()
// for index, r := range source {
// 	switch r {
// 	case '+', '-', '*', '/':
// 		st.Push(string(r))
// 	default:
// 		res += string(r)
// 	}
// }
// return
