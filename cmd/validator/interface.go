package validator

type Validator interface {
	// Validate action is run by calling ValidateString function
	ValidateString(input string) error
}
