package validator

import (
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// It Validates whether the product can have the given status
type ProductStatusValidator struct {
}

// Validate action is run by calling Validate function.
// All validator must implement below function.

func NewProductStatusValidator() Validator {
	return &ProductStatusValidator{}
}

func (v ProductStatusValidator) IsValid(input interface{}) bool {
	inStatus := input.(entity.ProductStatus)

	var StatusInInt int8
	if inStatus.Relayable {
		StatusInInt = StatusInInt | 1<<2
	} else if inStatus.Stored {
		StatusInInt = StatusInInt | 1<<1
	} else if inStatus.Transmitted {
		StatusInInt = StatusInInt | 1<<0
	}

	println(StatusInInt)

	if 1 <= StatusInInt && StatusInInt <= 3 {
		return false
	} else {
		return true
	}

}
