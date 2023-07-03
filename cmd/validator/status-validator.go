package validator

type StatusValidator struct {
	//regular expression of allowed form of stock input
	validPatten string
}

func NewStatusValidator() *StatusValidator {
	v := &StatusValidator{
		validPatten: "",
	}

	return v
}

func (v *StatusValidator) ValidateString(input string) error {
	//TODO: confirm status and
	return nil
}
