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
	// TODO: check required code form
	v.validPatten = "[0-9, A-z]+-[a-z]{2}"
	return &v
}
