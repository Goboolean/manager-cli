package status_test

import (
	"context"
	"testing"

	statusMock "github.com/Goboolean/manager-cli/internal/adaptor/status/mock"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/manager-cli/internal/domain/service/status"
	"github.com/stretchr/testify/assert"
)

var targetId string
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	targetId = "stock.apple.usa"
	m.Run()
}

func TestAdd(t *testing.T) {

	t.Run("RemoveNormal", func(t *testing.T) {
		//arrange
		var instance *status.StatusService
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)
		if err := statusAdaptor.SetStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   true,
				Stored:      false,
				Transmitted: false,
			}); err != nil {
			panic(err)
		}

		//act

		err := instance.AddStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   true,
				Stored:      true,
				Transmitted: false,
			})

		//assert
		actual, _ := statusAdaptor.GetStatus(ctx, targetId)
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
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)
		if err := statusAdaptor.SetStatus(ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   false,
				Stored:      false,
				Transmitted: false,
			}); err != nil {
			panic(err)
		}

		//act
		err := instance.AddStatus(
			ctx,
			targetId,
			entity.ProductStatus{
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
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)
		if err := statusAdaptor.SetStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   true,
				Stored:      true,
				Transmitted: false,
			}); err != nil {
			panic(err)
		}

		//act
		err := instance.RemoveStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   false,
				Stored:      true,
				Transmitted: false,
			})

		//assert
		actual, _ := statusAdaptor.GetStatus(ctx, targetId)
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
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)
		if err := statusAdaptor.SetStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   true,
				Stored:      true,
				Transmitted: false,
			}); err != nil {
			panic(err)
		}

		//act
		err := instance.RemoveStatus(
			ctx,
			targetId,
			entity.ProductStatus{
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
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)

		//act
		err := instance.SetStatus(
			ctx,
			targetId,
			entity.ProductStatus{
				Relayable:   true,
				Stored:      false,
				Transmitted: true,
			})

		//assert
		actual, _ := statusAdaptor.GetStatus(ctx, targetId)
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
		statusAdaptor, _ := statusMock.New()
		instance = status.New(
			statusAdaptor,
		)

		//act
		err := instance.SetStatus(
			ctx,
			targetId,
			entity.ProductStatus{
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
	statusAdaptor, _ := statusMock.New()
	instance = status.New(
		statusAdaptor,
	)

	statusAdaptor.SetStatus(
		ctx,
		targetId,
		entity.ProductStatus{
			Relayable:   true,
			Stored:      false,
			Transmitted: true,
		})

	//act
	res, err := instance.GetStatus(ctx, targetId)

	assert.NoError(t, err)
	assert.Equal(t, entity.ProductStatus{
		Relayable:   true,
		Stored:      false,
		Transmitted: true,
	},
		res)
}
