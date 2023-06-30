package validator

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestValidatorStockIn(t *testing.T) {

	//init interface
	var v Validator
	stockV := NewStockInputValidator()
	v = stockV

	//valid case
	var validTestCase = []string{"20214-ko", "AAPL-us", "aapl-us", "251422-us"}

	for _, i := range validTestCase {
		if res := v.ValidateString(i); res != nil {
			t.Errorf("Fail test case:" + i)
		}
	}

	var inValidTestCase = []string{"", "abecdefd", "12345", "12345a"}

	for _, i := range inValidTestCase {
		if res := v.ValidateString(i); res == nil {
			t.Errorf("Fail test case:" + i)
		}
	}
}

func TestDateValidator(t *testing.T) {
	var v Validator
	dateV := NewDateValidator()
	v = dateV

	//valid case
	var validTestCase = []string{"2023/05/01", "2023/12/31", "2023/02/21", "1998/03/12", "2023/10/20", "2023/10/12"}

	for _, i := range validTestCase {
		if res := v.ValidateString(i); res != nil {
			t.Errorf("Fail test case:" + i)
		}
	}

	var inValidTestCase = []string{"2023/13/01", "2023/12/33/", "2023/2/2", "23/03/12", "2023/2/13"}

	for _, i := range inValidTestCase {
		if res := v.ValidateString(i); res == nil {
			t.Errorf("Fail test case:" + i)
		}
	}
}
