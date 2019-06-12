package validator

type ValidatorBuilderInterface interface {
	AddFieldToValidate(fieldName string, validator Validator) ValidatorBuilder
	ValidateAll() bool
	GetAllMessages() (messages []string)
}

func (validatorBuilder *ValidatorBuilder) AddFieldToValidate(fieldName string, validator Validator) ValidatorBuilder {
	if value, keyExists := validatorBuilder.FieldsToValidate[fieldName]; not(keyExists) {
		validatorBuilder.Validators[fieldName] = validator
	}

	return validatorBuilder
}

func (validatorBuilder *ValidatorBuilder) ValidateAll() bool {

}
