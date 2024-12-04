// Package postfixcalc implements Reverse Polish Notation
// for basic arithmetic operations: addition, subtraction, multiplication and division.
package postfixcalc

import (
	"MathEvaluator/stack"
	"log"
	"regexp"
	"strconv"
)

// Calc structure defines the object which implements the RPN algorithm.
// It contains expression field - a field which contains the expression to be computed
// and a map of priorities with the priorities for arithmetic operations.
type Calc struct {
	expression string
	priorities map[string]int
}

// NewCalc function returns a new *Calc object.
// For parameters you need to pass an expression that you want to calculate.
func NewCalc(expression string) *Calc {
	var priorities = map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	return &Calc{expression: expression, priorities: priorities}
}

// Expression func returns Calc.exression value.
func (c *Calc) Expression() string {
	return c.expression
}

// SetExpression func allows a new value to be set for the expression.
func (c *Calc) SetExpression(expression string) {
	c.expression = expression
}

// getTokens func splits the given mathematical expression into tokens for evaluation.
func (c *Calc) getTokens(expression string) (matched []string) {
	re, err := regexp.Compile("[()]|[0-9]+|[+\\-*/]")
	if err != nil {
		log.Fatal(err)
	}
	matched = re.FindAllString(expression, -1)
	return
}

// getPostfixForm func converts a mathematical expression represented as a slice of tokens
// in infix notation into its equivalent in postfix notation (Reverse Polish Notation).
func (c *Calc) getPostfixForm(source []string) string {
	var res string
	st := stack.NewStack[string]()
	for _, ch := range source {
		switch ch {
		case "+", "-", "*", "/":
			for peek, err := st.Peek(); err == nil && c.priorities[peek] >= c.priorities[ch]; peek, err = st.Peek() {
				pop, err := st.Pop()
				if err != nil {
					break
				}
				res += pop + " "
			}
			st.Push(ch)
		case "(":
			st.Push(ch)
		case ")":
			for peek, err := st.Peek(); err == nil && peek != "("; peek, err = st.Peek() {
				pop, err := st.Pop()
				if err != nil {
					break
				}
				res += pop + " "
			}
			_, _ = st.Pop()

		default:
			res += ch + " "
		}
	}

	for pop, err := st.Pop(); err == nil; pop, err = st.Pop() {
		res += pop + " "
	}

	return res
}

// Calculate evaluates the mathematical expression stored in the Calc struct
// and returns the computed result as a float64.
func (c *Calc) Calculate() float64 {
	resStr := c.getPostfixForm(c.getTokens(c.expression))
	tokens := c.getTokens(resStr)
	st := stack.NewStack[float64]()
	var res float64

	for _, str := range tokens {
		var nums [2]float64
		char, err := strconv.ParseFloat(str, 64)
		if err == nil {
			st.Push(char)
		} else {
			i := 0
			for _, err := st.Peek(); err == nil && i < 2; _, err = st.Peek() {
				pop, _ := st.Pop()
				nums[i] = pop
				i++
			}
			switch str {
			case "+":
				res = (nums[0] + nums[1])
				st.Push(res)
			case "-":
				res = (nums[1] - nums[0])
				st.Push(res)
			case "*":
				res = (nums[0] * nums[1])
				st.Push(res)
			case "/":
				res = (nums[1] / nums[0])
				st.Push(res)
			}
		}
	}

	return res
}
