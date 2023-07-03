package validator

import (
	"errors"
	"regexp"
)

//this

type stockValidator struct {
	supportedLocation []string
	//regular expression of allowed form of stock input
	validPatten string
}

func (v *stockValidator) ValidateString(input string) error {

	if matched, _ := regexp.MatchString(v.validPatten, input); !matched {
		return errors.New("invalid pattern")
	}

	return nil
}

func NewStockValidator() *stockValidator {
	v := stockValidator{}
	v.supportedLocation = []string{"kor", "usa"}
	v.validPatten = "^[0-9, A-Z]+-[a-z]{3}$"
	return &v
}
