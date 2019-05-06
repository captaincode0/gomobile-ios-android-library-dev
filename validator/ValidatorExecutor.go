package validator

type ValidatorExecutor interface {
	//
	validate(input string) bool
}
