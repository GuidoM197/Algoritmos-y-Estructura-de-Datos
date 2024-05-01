package operations

import (
	"fmt"
	"strconv"
	"strings"
	"tdas/pila"
)

const (
	SUM       = "+"
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
		counterNumbers int
		result         int64
		err            error
	)
	newStack := pila.CrearPilaDinamica[int64]()

	for _, value := range operation {

		number, errStr := strconv.ParseInt(value, 10, 64)

		if errStr != nil && !strings.Contains(OPERATORS, value) {
			err = fmt.Errorf("Could not be converted correctly: %v\n", err)

		} else if errStr != nil && strings.Contains(OPERATORS, value) {
			result, err :=
		} else {
			counterNumbers++
			newStack.Apilar(number)

		}
		newStack.Apilar(result)
		counterNumbers++

	}
	if (!newStack.EstaVacia() && counterNumbers != 1) || newStack.EstaVacia() {
		return result, fmt.Errorf("There are extra operators/numbers or unfinished equations")

	}
	return result, err
}

func processorOperation(stackOfNumbers pila.Pila[int64], operator string, elementsInStack int) (int64, error, int) {

}


func getOperators(stack pila.Pila[int64], quantity int, res []int64) []int64 {
	for _ = range quantity {
		res = append(res, stack.Desapilar())
	}
	return res
}

func Operations(operator string, stack pila.Pila[int64]) (int64, error) {
	var (
		op Operator
		arithmeticOperator arithmetics
		numbers []int64
		result int64
		err error
	)

	switch operator {
	case SUM: arithmeticOperator = sum{}
	case SUBTRACTION: arithmeticOperator = subtract{}
	case MULTIPLICATION: arithmeticOperator = multiplication{}
	case DIVISION: arithmeticOperator = division{}
	case RAISE: arithmeticOperator = raise{}
	case LOGARITHM: arithmeticOperator = logarithm{}
	case SQRT: arithmeticOperator = sqrt{}
	case TERNARY: arithmeticOperator = ternary{}
	}
	op = CreateOperator(arithmeticOperator)
	numbers = getOperators(stack, op.arity, numbers)
	result, err =
}
