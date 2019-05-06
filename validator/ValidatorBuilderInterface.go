package validator

type ValidatorBuilderInterface interface {
	// Add one field to validate on internal map with one validator
	AddFieldToValidate(fieldName string, validator Validator) ValidatorBuilder
	ValidateAll() bool
}
