package operations

type Operator interface {
	// Using an arithmetic operator, it solves the equation and returns an (int64, error).
	Operation() (int64, error)
}
