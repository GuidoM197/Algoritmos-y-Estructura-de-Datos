package operations

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tdas/pila"
)

type Operator struct {
	representation string
	arity          int
	operate        func(operators []int64) (int64, error)
}

var (
	SUM            = "+"
	SUBTRACTION    = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
	SQRT           = "sqrt"
	RAISE          = "^"
	LOGARITHM      = "log"
	TERNARY        = "?"
	OPSUM          = Operator{SUM, 2, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( + )")
		}
		return operands[0] + operands[1], nil
	}}
	OPSUBTRACTION = Operator{SUBTRACTION, 2, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( - )")
		}
		return operands[1] - operands[0], nil
	}}
	OPMULTIPLICATION = Operator{MULTIPLICATION, 2, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( * )")
		}
		return operands[1] * operands[0], nil
	}}
	OPDIVISION = Operator{DIVISION, 2, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( / )")
		}
		if operands[0] == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return operands[1] / operands[0], nil
	}}
	OPSQRT = Operator{SQRT, 1, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( sqrt )")
		}
		if operands[0] < 0 {
			return 0, fmt.Errorf("Sqrt invalid")
		}
		return int64(math.Sqrt(float64(operands[0]))), nil
	}}
	OPRAISE = Operator{RAISE, 2, func(operands []int64) (int64, error) {
		if len(operands) == 0 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( ^ )")
		}
		if operands[0] < 0 {
			return 0, fmt.Errorf("Raise invalid")
		}
		return int64(math.Pow(float64(operands[1]), float64(operands[0]))), nil
	}}
	OPLOGARITHM = Operator{LOGARITHM, 2, func(operands []int64) (int64, error) {
		if operands[1] < 1 || operands[0] <= 1 {
			return 0, fmt.Errorf("Unable to calculate a log with argument less than 1")
		}
		return int64(math.Log(float64(operands[1])) / math.Log(float64(operands[0]))), nil
	}}
	OPTERNARY = Operator{TERNARY, 3, func(operands []int64) (int64, error) {
		if len(operands) < 3 {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( Ternary )")
		}
		if operands[2] != 0 {
			return operands[1], nil
		}
		return operands[0], nil
	}}
	OPERATORS = []Operator{OPSUM, OPSUBTRACTION, OPMULTIPLICATION, OPDIVISION, OPSQRT, OPRAISE, OPLOGARITHM, OPTERNARY}
)

func IdentifyOperations(operation []string) (int64, error) {
	var (
		result int64
		err    error
	)
	stack := pila.CrearPilaDinamica[int64]()

	for _, value := range operation {

		number, errStr := strconv.ParseInt(value, 10, 64)

		if errStr != nil && strings.Contains("+-*/sqrt^log?", value) {
			result, err = Operations(value, stack)
			stack.Apilar(result)

		} else if errStr != nil && !strings.Contains("+-*/sqrt^log?", value) {
			err = fmt.Errorf("Could not be converted correctly: %v\n", errStr)

		} else {
			stack.Apilar(number)

		}
	}

	if stack.EstaVacia() {
		return result, fmt.Errorf("There are extra operators/numbers or unfinished equations")

	}
	result = stack.Desapilar()
	if !stack.EstaVacia() {
		return result, fmt.Errorf("There are extra operators/numbers or unfinished equations")

	}
	return result, err
}

func getOperators(stack pila.Pila[int64], quantity int, res []int64) ([]int64, error) {
	for _ = range quantity {
		if stack.EstaVacia() {
			return res, fmt.Errorf("There arent enough numbers.")
		}
		res = append(res, stack.Desapilar())
	}
	return res, nil
}

func Operations(operator string, stack pila.Pila[int64]) (int64, error) {
	var (
		err     error
		numbers []int64
	)

	for _, value := range OPERATORS {
		if value.representation == operator {
			numbers, err = getOperators(stack, value.arity, numbers)
			if err != nil {
				return 0, fmt.Errorf("%v\n", err)
			}
			return value.operate(numbers)
		}
	}
	return 0, fmt.Errorf("Error to processing the equation.")
}
