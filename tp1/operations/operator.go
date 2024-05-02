package operations

type Operator struct {
	representation string
	arity          int
	operate        func(operators []int64) (int64, error)
}
