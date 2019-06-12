package validator

import "regexp"

type ValidatorExecutor interface {
	validate(input string) bool
}

func (validator *Validator) validate(input string) bool {
	isValid, error = regexp.MatchString(validator.Regex, input)

	if error != nil && isValid {
		return true
	}

	return false
}
