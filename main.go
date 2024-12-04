package main

import (
	"MathEvaluator/postfixcalc"
	"fmt"
)

func main() {
	expression := "1+1"
	calc := postfixcalc.NewCalc(expression)

	//test cases
	firstTestCase := [10]string{
		"5 + 3",
		"8 - 4",
		"6 * 2",
		"9 / 3",
		"(7 + 2) - 3",
		"(5 * 2) + 8",
		"9 * (4 + 2)",
		"((8 + 2) * 5) / 3",
		"7 - ((3 + 2) * 4)",
		"(((9 * (6 + 2)) / 3) - 4)",
	}

	secondTestCase := [10]string{
		"12 + 8",
		"25 - 17",
		"15 * 7",
		"36 / 12",
		"(18 + 24) - 10",
		"(30 - 5) * 4",
		"48 / (6 + 6)",
		"((25 + 10) * 3) / 5",
		"50 - ((12 + 8) * 2)",
		"(((40 * 2) / 5) + 30) - 12",
	}

	thirdTestCase := [10]string{
		"123 + 45",
		"500 - 250",
		"36 * 12",
		"144 / 12",
		"(90 + 15) * 2",
		"((45 * 3) + 100) - 30",
		"250 / (10 + 5)",
		"300 - (75 * 4)",
		"(12 + 8) * (50 / 5)",
		"((1000 / 25) + 20) * 3",
	}

	showTestCase(calc, firstTestCase)
	showTestCase(calc, secondTestCase)
	showTestCase(calc, thirdTestCase)
}

func showTestCase(calc *postfixcalc.Calc, arr [10]string) {
	fmt.Printf("\n====================\n")
	for _, expr := range arr {
		calc.SetExpression(expr)
		fmt.Printf("Expression: %s, Result: %.2f\n", expr, calc.Calculate())
	}
	fmt.Printf("====================\n")
}
