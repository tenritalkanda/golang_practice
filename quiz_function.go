package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)

	fmt.Println(total)

	math_operation := []string{"+", "-", "*", "/", "="}

	for _, value := range math_operation {
		result, err := calculate(10, 2, value)
		if err != nil {
			fmt.Println(err, 10, value, 2)
		} else {
			fmt.Println("result = ", result)
		}

	}

}

func sum(all_scores []int) (sum_all int) {
	for _, value := range all_scores {
		sum_all += value
	}

	return
}

func calculate(a int, b int, operand string) (int, error) {
	switch operand {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	case "*":
		return a * b, nil
	default:
		return 0, errors.New("Error : can't do math operation with")
	}
}
