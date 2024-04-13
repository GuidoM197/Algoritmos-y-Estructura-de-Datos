package operations

import (
	"fmt"
	"math"
	"strconv"
	"tdas/pila"
)

const (
	OPERATORSQRT    = "sqrt"
	OPERATORTERNARY = "?"
)

func IdentifyOperations(operation []string) (int64, error) {
	var (
		counterNumbers int
		result         int64
		err            error
	)
	newStack := pila.CrearPilaDinamica[int64]()

	for _, value := range operation {

		if value != "+" && value != "-" && value != "*" && value != "/" && value != "^" && value != "log" && value != "sqrt" && value != "?" {
			number, errStr := strconv.ParseInt(value, 10, 64)
			if errStr != nil {
				err = fmt.Errorf("Could not be converted correctly: %v\n", err)

			} else {
				counterNumbers++
				newStack.Apilar(number)

			}
		} else {
			result, err, counterNumbers = processorOperation(newStack, value, counterNumbers)
			newStack.Apilar(result)
			counterNumbers++

		}
	}
	if (!newStack.EstaVacia() && counterNumbers != 1) || newStack.EstaVacia() {
		return result, fmt.Errorf("There are extra operators/numbers or unfinished equations")

	}
	return result, err
}

func processorOperation(stackOfNumbers pila.Pila[int64], operator string, elementsInStack int) (int64, error, int) {

	var (
		result int64
		n1     int64
		n2     int64
		err    error
	)

	if stackOfNumbers.EstaVacia() {
		return 0, fmt.Errorf("There are more operators than numbers"), elementsInStack
	}
	if operator == OPERATORSQRT {
		if elementsInStack >= 1 {
			n1 = stackOfNumbers.Desapilar()
			elementsInStack--
			result, err = sqrt(n1)
			return result, err, elementsInStack

		} else {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( sqrt )"), elementsInStack

		}
	} else if operator == OPERATORTERNARY {
		if elementsInStack >= 3 {
			result, err, elementsInStack = ternary(stackOfNumbers, elementsInStack)
			return result, err, elementsInStack

		} else {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( Ternary )"), elementsInStack

		}
	}
	if !stackOfNumbers.EstaVacia() {
		if elementsInStack >= 2 {
			n2 = stackOfNumbers.Desapilar()
			n1 = stackOfNumbers.Desapilar()
			elementsInStack -= 2
			result, err = Operations(n1, n2, operator)

		} else {
			return 0, fmt.Errorf("There are not enough numbers to performance this operation ( +, -, *, /, log )"), elementsInStack
		}
	}
	return result, err, elementsInStack
}

func Operations(n1, n2 int64, operator string) (int64, error) {
	switch operator {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		return division(n1, n2)
	case "^":
		return power(n1, n2)
	default:
		return logarithm(n1, n2)
	}
}

func division(n1, n2 int64) (int64, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return n1 / n2, nil
}

func power(n1, n2 int64) (int64, error) {
	if n2 < 0 {
		return 0, fmt.Errorf("Pow invalid")
	}
	return int64(math.Pow(float64(n1), float64(n2))), nil
}

func logarithm(n1, n2 int64) (int64, error) {
	if n1 < 1 || n2 <= 1 {
		return 0, fmt.Errorf("No se puede calcular un logaritmo con argumento menor a 1")
	}
	return int64(math.Log(float64(n1)) / math.Log(float64(n2))), nil
}

func ternary(stackNumbers pila.Pila[int64], elementStack int) (int64, error, int) {
	n3 := stackNumbers.Desapilar()
	n2 := stackNumbers.Desapilar()
	n1 := stackNumbers.Desapilar()
	elementStack -= 3

	result := n1
	if n1 != 0 {
		result = n2
	} else if n2 != 0 {
		result = n3
	}
	return result, nil, elementStack
}

func sqrt(n1 int64) (int64, error) {
	if n1 < 0 {
		return 0, fmt.Errorf("Sqrt invalid")
	}
	return int64(math.Sqrt(float64(n1))), nil
}
