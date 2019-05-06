package validator

type ValidatorMessageBuilder interface {
  GetErrorMessage(context string) string
}

func (validator *Validator) GetErrorMessage(context string) string {
  if (context == "") {
    context = validator.ErrorMessage
  }

  return fmt.Sprintf(
    "Validation Error [Type: %] context: %",
    validator.Name,
    context
  )
}
