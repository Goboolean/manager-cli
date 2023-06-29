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
	if err := v.ValidateString("20214-ko"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("AAPL-us"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("aapl-us"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("251422-us"); err != nil {
		t.Errorf("Fail")
	}

	//invalid case
	if err := v.ValidateString(""); err == nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("abecdefd"); err == nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("12345"); err == nil {
		t.Errorf("Fail")
	}

	if err := v.ValidateString("12345a"); err == nil {
		t.Errorf("Fail")
	}
}
