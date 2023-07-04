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
		return errors.New("invalid stock")
	}

	return nil
}

func NewStockValidator() Validator {
	v := &stockValidator{}
	v.supportedLocation = []string{"kor", "usa"}
	//valid stock code: Decimal numbers or uppercase letters of the alphabet.
	//valid stock origin: Standard 3-letter lowercase alphabetic country code fallowed ISO3166-1
	v.validPatten = "^[0-9, A-Z]+-[a-z]{3}$"
	return v
}
