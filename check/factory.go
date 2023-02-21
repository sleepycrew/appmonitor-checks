package check

type Factory interface {
	GetName() string
	BuildCheck(input map[string]string) (Check, error)
}
