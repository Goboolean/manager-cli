package status_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Goboolean/manager-cli/internal/adaptor/status"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

var a *status.StatusAdaptor

func TestMain(m *testing.M) {
	var err error
	fmt.Println("Test main is called")

	a, err = status.New()
	fmt.Println("Status Adaptor is created")

	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestGetStatus(t *testing.T) {
	status, err := a.GetStatus("AAPL")

	if err != nil {
		t.Errorf(err.Error())
	}

	if status.Relayable != true {
		t.Errorf("Expected response of relayable to be true, got %s", strconv.FormatBool(status.Relayable))
	}

	if status.Stored != false {
		t.Errorf("Expected response stored to be false, got %s", strconv.FormatBool(status.Relayable))
	}

	if status.Transmitted != true {
		t.Errorf("Expected response transmitted to be true AAPL, got %s", strconv.FormatBool(status.Relayable))
	}

}

func TestSetStatus(t *testing.T) {

	desired := entity.ProductStatus{
		Relayable:   true,
		Stored:      true,
		Transmitted: true,
	}

	if err := a.SetStatus("AAPL", desired); err != nil {
		t.Errorf(err.Error())
	}
}
