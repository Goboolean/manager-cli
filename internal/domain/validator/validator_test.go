package validator

import (
	"os"
	"testing"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
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

func TestProductStatusValidator(t *testing.T) {
	validator := NewProductStatusValidator()

	if validator == nil {
		t.Error("Expected validator to be created")
	}

	var testCases []*pair.Pair[entity.ProductStatus, bool]

	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: false, Stored: false, Transmitted: false}, true))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: false, Stored: false, Transmitted: true}, false))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: false, Stored: true, Transmitted: false}, false))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: false, Stored: true, Transmitted: true}, false))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: true, Stored: false, Transmitted: false}, true))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: true, Stored: false, Transmitted: true}, true))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: true, Stored: true, Transmitted: false}, true))
	testCases = append(testCases, pair.New(entity.ProductStatus{Relayable: true, Stored: true, Transmitted: true}, true))

	for _, testCase := range testCases {
		if validator.IsValid(testCase.First) != testCase.Second {
			t.Errorf("Expected %v to be %v", testCase.First, testCase.Second)
		}
	}

}
