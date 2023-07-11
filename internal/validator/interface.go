package validator

type Validator interface {
	// Validate action is run by calling ValidateString function.
	// All validator must implement below function.
	ValidateString(input string) error
}
