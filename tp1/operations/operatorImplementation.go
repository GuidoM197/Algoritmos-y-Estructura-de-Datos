package operations

import (
	"fmt"
	"math"
)

type operator struct {
	representation arithmetics
	arity          int
	numbers        []int64
}

func CreateOperator(rep arithmetics, arrNumbers []int64) Operator {
	return &operator{representation: rep, arity: rep.getArity(), numbers: arrNumbers}
}

func (op *operator) Operation() (int64, error) {
	switch op.representation {
	case sum{}:
		return operateSum(op.numbers)
	case subtraction{}:
		return operateSubtraction(op.numbers)
	case multiplication{}:
		return operateMultiplication(op.numbers)
	case division{}:
		return operateDivision(op.numbers)
	case raise{}:
		return operateRaise(op.numbers)
	case logarithm{}:
		return operateLogarithm(op.numbers)
	case sqrt{}:
		return operateSqrt(op.numbers)
	case ternary{}:
		return operateTernary(op.numbers)
	default:
		return 0, nil
	}
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
	return operators[1] - operators[0], nil
}

func operateMultiplication(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( * )")
	}
	return operators[1] * operators[0], nil
}

func operateDivision(operators []int64) (int64, error) {
	if len(operators) == 0 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( / )")
	}
	if operators[0] == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return operators[1] / operators[0], nil
}

func operateRaise(operators []int64) (int64, error) {
	if len(operators) < 0 {
		return 0, fmt.Errorf("Raise invalid")
	}
	return int64(math.Pow(float64(operators[1]), float64(operators[0]))), nil
}

func operateLogarithm(operators []int64) (int64, error) {
	if operators[1] < 1 || operators[0] <= 1 {
		return 0, fmt.Errorf("Unable to calculate a log with argument less than 1")
	}
	return int64(math.Log(float64(operators[1])) / math.Log(float64(operators[0]))), nil
}

func operateTernary(operators []int64) (int64, error) {
	if len(operators) < 3 {
		return 0, fmt.Errorf("There are not enough numbers to performance this operation ( Ternary ).")
	}

	if operators[2] == 0 {
		return operators[1], nil
	} else {
		return operators[0], nil
	}
}
