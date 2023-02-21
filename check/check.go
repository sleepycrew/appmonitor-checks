package check

type Result struct {
	Result int8
	Value  string
}

type Check interface {
	RunCheck(result chan<- Result)
}
