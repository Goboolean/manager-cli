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
		// Valid date: 3 decimal numbers separated by "/" like yyyy/mm/dd
		// Valid range: year-[0000-9999], month-[00-12], day-[00-31]
		validPatten: "^[0-9]{4}\\/(0[1-9]|1[012])\\/(0[1-9]|[12][0-9]|3[01])$",
	}

	return v
}

func (v *dateValidator) ValidateString(input string) error {
	if matched, _ := regexp.MatchString(v.validPatten, input); !matched {
		return errors.New("invalid form of date")
	}

	return nil
}
