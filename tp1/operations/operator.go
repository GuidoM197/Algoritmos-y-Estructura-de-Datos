package operations

type Operator interface {
	Operation() (int64, error)
	//OperateSum() (int64,error)
	//OperateSubtraction() (int64,error)
	//OperateMultiplication() (int64,error)
	//OperateDivision() (int64,error)
	//OperateLogarithm() (int64,error)
	//OperateSqrt() (int64,error)
	//OperateTernary() (int64,error)
}
