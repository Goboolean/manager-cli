package status_test

import (
	"testing"

	statusMock "github.com/Goboolean/manager-cli/internal/adaptor/status/mock"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/manager-cli/internal/domain/service/status"
	"github.com/stretchr/testify/assert"
)

var targetId string

func TestMain(m *testing.M) {
	targetId = "stock.apple.usa"
	m.Run()
}

func TestAdd(t *testing.T) {

	t.Run("RemoveNormal", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)
		if err := statusAdpator.SetStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: false,
		}); err != nil {
			panic(err)
		}

		//act

		err := instance.AddStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      true,
			Transmitted: false,
		})

		//assert
		actual, _ := statusAdpator.GetStatus(targetId)
		assert.NoError(t, err)
		assert.Equal(t, entity.ProductStatus{
			Relayable:   true,
			Stored:      true,
			Transmitted: false,
		},
			actual)
	})

	t.Run("AddWithInvalid", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)
		if err := statusAdpator.SetStatus(targetId, entity.ProductStatus{
			Relayable:   false,
			Stored:      false,
			Transmitted: false,
		}); err != nil {
			panic(err)
		}

		//act
		err := instance.RemoveStatus(targetId, entity.ProductStatus{
			Relayable:   false,
			Stored:      true,
			Transmitted: false,
		})

		assert.Error(t, err)

	})

}

func TestRemove(t *testing.T) {

	t.Run("RemoveNormal", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)
		if err := statusAdpator.SetStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      true,
			Transmitted: false,
		}); err != nil {
			panic(err)
		}

		//act
		err := instance.RemoveStatus(targetId, entity.ProductStatus{
			Relayable:   false,
			Stored:      true,
			Transmitted: false,
		})

		//assert
		actual, _ := statusAdpator.GetStatus(targetId)
		assert.NoError(t, err)
		assert.Equal(t, entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: false,
		},
			actual)
	})

	t.Run("RemoveWithInvalid", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)
		if err := statusAdpator.SetStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      true,
			Transmitted: false,
		}); err != nil {
			panic(err)
		}

		//act
		err := instance.RemoveStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: false,
		})

		assert.Error(t, err)

	})

}

func TestSet(t *testing.T) {
	t.Run("SetNormal", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)

		//act
		err := instance.SetStatus(targetId, entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: true,
		})

		//assert
		actual, _ := statusAdpator.GetStatus(targetId)
		assert.NoError(t, err)
		assert.Equal(t, entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: true,
		},
			actual)
	})

	t.Run("SetWithExpectedError", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdpator, _ := statusMock.New()
		instance = status.New(
			statusAdpator,
		)

		//act
		err := instance.SetStatus(targetId, entity.ProductStatus{
			Relayable:   false,
			Stored:      false,
			Transmitted: true,
		})

		//assert
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	//arrange
	var instance *status.StatusService
	statusAdpator, _ := statusMock.New()
	instance = status.New(
		statusAdpator,
	)

	statusAdpator.SetStatus(
		targetId,
		entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: true,
		})

	//act
	res, err := instance.GetStatus(targetId)

	assert.NoError(t, err)
	assert.Equal(t, entity.ProductStatus{
		Relayable:   true,
		Stored:      false,
		Transmitted: true,
	},
		res)
}
