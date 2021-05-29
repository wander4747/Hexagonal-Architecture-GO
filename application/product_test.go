package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/wander4747/go-hexagonal/application"
	"reflect"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Notebook"
	product.Status = application.DISABLED
	product.Price = 3000

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Notebook"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 3000
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Notebook"
	product.Status = application.DISABLED
	product.Price = 3000

	product.Status = ""
	require.EqualValues(t, product.Status, "")

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	require.NotEmpty(t, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Notebook"

	require.NotEmpty(t, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.NotEmpty(t, product.GetStatus())

	require.EqualValues(t, product.Status, "enabled")

	product.Status = application.DISABLED
	require.EqualValues(t, product.Status, "disabled")
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 3000

	require.NotEmpty(t, product.GetPrice())

	typeOf := reflect.TypeOf(product.Price).Kind()
	require.True(t, typeOf == reflect.Float64)
}