package Validator

/**
  Idea of validators

  Validate a structure instance (AKA Object)

  import v "googlemapsearch/validator"

  var location Location = Location{-12.178271, 20.10210}
  var validatorBuilder ValidatorBuilder

  validatorBuilder.FieldsToValidate = location.ToMap()

  validatorBuilder.AddFieldToValidate("Location", FloatValidator{})
  validatorBuilder.Add
*/

type ValidatorBuilder struct {
	FieldsToValidate map[string]string
	Validators       map[string]Validator
}
