package validator

import (
	"regexp"
)

type ProductIdValidator struct {
	//Regular expression of valid product id
	validPatterns []string //pattern: {product type}.{product name}.{location}
}

func NewProductIdValidator() Validator {

	return &ProductIdValidator{
		validPatterns: []string{"^stock.[a-z]*.[a-z]{3}$", "^coin.[a-z]*.null$"},
	}
}

func (v ProductIdValidator) IsValid(input interface{}) bool {
	inStr := input.(string)

	for _, pattern := range v.validPatterns {
		match, _ := regexp.MatchString(pattern, inStr)
		if match {
			return true
		}
	}
	return false

}
