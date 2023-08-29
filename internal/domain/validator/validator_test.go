package validator

import (
	"os"
	"testing"

	"github.com/notEpsilon/go-pair"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNewProductIdValidator(t *testing.T) {
	validator := NewProductIdValidator()
	//First is value to test second is expected result
	var testCases []*pair.Pair[string, bool]

	//test first field validation
	testCases = append(testCases, pair.New("coin.bitcoin.null", true))
	testCases = append(testCases, pair.New("stock.apple.usd.", false))
	testCases = append(testCases, pair.New("bark.bitcoin.usd.", false))
	testCases = append(testCases, pair.New(".bitcoin.usd", false))

	//test second field validation
	testCases = append(testCases, pair.New("stock.apple.usd", true))
	testCases = append(testCases, pair.New("stock.AAPL.usd", false))
	testCases = append(testCases, pair.New("stock..USD", false))

	//test third field validation
	testCases = append(testCases, pair.New("stock.samsung.usa.", false))
	testCases = append(testCases, pair.New("stock.samsung.korea", false))
	testCases = append(testCases, pair.New("stock.samsung.", false))

	testCases = append(testCases, pair.New("...", false))

	for _, testCase := range testCases {
		if validator.IsValid(testCase.First) != testCase.Second {
			t.Errorf("Expected %v to be %v", testCase.First, testCase.Second)
		}
	}

}
