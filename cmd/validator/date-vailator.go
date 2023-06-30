package validator

import (
	"errors"
	"regexp"
)

type dateValidator struct {
	//regular expression of allowed form of stock input
	validPatten string
}

func NewDateValidator() *dateValidator {
	v := &dateValidator{
		validPatten: "^[0-9]{4}\\/(0[1-9]|1[012])\\/(0[1-9]|[12][0-9]|3[01])$",
	}
	// yyyy/mm/dd pattern

	return v
}

func (v *dateValidator) ValidateString(input string) error {
	if matched, _ := regexp.MatchString(v.validPatten, input); !matched {
		return errors.New("invalid form of date")
	}

	return nil
}
