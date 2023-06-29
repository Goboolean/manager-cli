package vaildator

import (
	"errors"
	"regexp"
)

//this

type stockInputVailator struct {
	supprtedLocation []string
	//regular expression of allowed form of stock input
	vaildPatten string
}

func (v *stockInputVailator) VaildateString(input string) error {

	if mached, _ := regexp.MatchString(v.vaildPatten, input); !mached {
		return errors.New("invaild pattern allowd pattern is [Stock Code]-[Location] example: 00001-kr")
	}

	return nil
}

func NewStockInputVailator() *stockInputVailator {
	v := stockInputVailator{}
	v.supprtedLocation = []string{"ko", "us"}
	// TODO: check requied code form
	v.vaildPatten = "[0-9, A-z]+-[a-z]{2}"
	return &v
}
