package vaildator

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestVaildatorStockIn(t *testing.T) {

	//init interface
	var v Vaildator
	stockV := NewStockInputVailator()
	v = stockV

	//vaild case
	if err := v.VaildateString("20214-ko"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("AAPL-us"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("aapl-us"); err != nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("251422-us"); err != nil {
		t.Errorf("Fail")
	}

	//invaid case
	if err := v.VaildateString(""); err == nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("abecdefd"); err == nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("12345"); err == nil {
		t.Errorf("Fail")
	}

	if err := v.VaildateString("12345a"); err == nil {
		t.Errorf("Fail")
	}
}
