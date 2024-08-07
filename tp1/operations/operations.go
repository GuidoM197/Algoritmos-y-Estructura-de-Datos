package operations

import (
	"fmt"
	"strconv"
	"strings"
	"tdas/pila"
)

const (
	SUM            = "+"
	SUBTRACTION    = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
	SQRT           = "sqrt"
	RAISE          = "^"
	LOGARITHM      = "log"
	TERNARY        = "?"
	OPERATORS      = "+-*/sqrt^log?"
)

func IdentifyOperations(operation []string) (int64, error) {
	var (
		result int64
		err    error
	)
	newStack := pila.CrearPilaDinamica[int64]()

	for _, value := range operation {

		number, errStr := strconv.ParseInt(value, 10, 64)

		if errStr != nil && strings.Contains(OPERATORS, value) {
			result, err = Operations(value, newStack)
			newStack.Apilar(result)

		} else {
			newStack.Apilar(number)

		}

	}
	if newStack.EstaVacia() {
		return result, fmt.Errorf("There are extra operators/numbers or unfinished equations.")

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
		op                 Operator
		arithmeticOperator arithmetics
		numbers            []int64
		err                error
	)

	switch operator {
	case SUM:
		arithmeticOperator = sum{}
	case SUBTRACTION:
		arithmeticOperator = subtraction{}
	case MULTIPLICATION:
		arithmeticOperator = multiplication{}
	case DIVISION:
		arithmeticOperator = division{}
	case RAISE:
		arithmeticOperator = raise{}
	case LOGARITHM:
		arithmeticOperator = logarithm{}
	case SQRT:
		arithmeticOperator = sqrt{}
	case TERNARY:
		arithmeticOperator = ternary{}
	}
	numbers, err = getOperators(stack, arithmeticOperator.getArity(), numbers)
	if err != nil {
		return 0, fmt.Errorf("%v\n", err)
	}
	op = CreateOperator(arithmeticOperator, numbers)
	return op.Operation()
}
