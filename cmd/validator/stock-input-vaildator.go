package validator

import (
	"errors"
	"regexp"
)

//this

type stockInputValidator struct {
	supportedLocation []string
	//regular expression of allowed form of stock input
	validPatten string
}

func (v *stockInputValidator) ValidateString(input string) error {

	if matched, _ := regexp.MatchString(v.validPatten, input); !matched {
		return errors.New("invalid pattern")
	}

	return nil
}

func NewStockInputValidator() *stockInputValidator {
	v := stockInputValidator{}
	v.supportedLocation = []string{"ko", "us"}
	v.validPatten = "[0-9, A-Z]+-[a-z]{2}"
	return &v
}
