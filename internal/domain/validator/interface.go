package validator

type Validator interface {
	// Validate action is run by calling Validate function.
	// All validator must implement below function.
	IsValid(input interface{}) bool
}
