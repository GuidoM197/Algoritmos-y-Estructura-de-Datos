package operations

import (
	"fmt"
	"math"
)

type operator struct {
	representation arithmetics
	arity          int
}

func CreateOperator(rep arithmetics, numbers []int64) Operator {
	return &operator{representation: rep, arity: rep.getArity()}
}

func (op *operator) SelectOperation() (int64, error) {
	switch op.representation {
	case sum{}:
		return operateSum()
	case subtraction{}:
		return operateSubtraction()
	case multiplication{}:
		return operateMultiplication()
	case division{}:
		return operateDivision()
	case raise{}:
		return operateRaise()
	case logarithm{}:
		return operateLogarithm()
	case sqrt{}:
		return operateSqrt()
	case ternary{}:
		return operateTernary()
	}
	return 0, nil
}

func operateSqrt(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( sqrt )")
	}
	if operators[0] < 0 {
		return 0, fmt.Errorf("Sqrt invalid")
	}
	return int64(math.Sqrt(float64(operators[0]))), nil
}

func operateSum(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( + )")
	}
	return operators[0] + operators[1], nil
}

func operateSubtraction(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( - )")
	}
	return operators[0] - operators[1], nil
}

func operateMultiplication(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( * )")
	}
	return operators[0] * operators[1], nil
}

func operateDivision(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return operators[0] / operators[1], nil
}

func operateRaise(operators []int64) (int64, error) {
	if len(operators) < 0 {
		return 0, fmt.Errorf("Raise invalid")
	}
	return int64(math.Pow(float64(operators[0]), float64(operators[1]))), nil
}

func operateLogarithm(operators []int64) (int64, error) {
	if operators[0] < 1 || operators[1] <= 1 {
		return 0, fmt.Errorf("Unable to calculate a log with argument less than 1")
	}
	return int64(math.Log(float64(operators[0])) / math.Log(float64(operators[1]))), nil
}

func operateTernary(operators []int64) (int64, error) {
	if len(operators) < 3 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( Ternary )")
	}

	if operators[0] == 0 {
		return operators[2], nil
	} else {
		return operators[1], nil
	}
}
