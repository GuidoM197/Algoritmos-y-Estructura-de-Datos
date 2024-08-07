package operations

type Operator interface {
	Operation() (int64, error)
}
