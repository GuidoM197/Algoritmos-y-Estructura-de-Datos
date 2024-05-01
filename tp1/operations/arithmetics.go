package operations

type arithmetics interface {
	// Return the arity.
	getArity() int
}

type sqrt struct{}

func (sqrt) getArity() int {
	return 1
}

type sum struct{}

func (sum) getArity() int {
	return 2
}

type subtraction struct{}

func (subtraction) getArity() int {
	return 2
}

type multiplication struct{}

func (multiplication) getArity() int {
	return 2
}

type division struct{}

func (division) getArity() int {
	return 2
}

type raise struct{}

func (raise) getArity() int {
	return 2
}

type logarithm struct{}

func (logarithm) getArity() int {
	return 2
}

type ternary struct{}

func (ternary) getArity() int {
	return 3
}
