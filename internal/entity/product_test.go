package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotNil(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100.0, product.Price)
	assert.NotNil(t, product.CreatedAt)
}

func Fuzz_NewProduct(f *testing.F) {
	// Add test cases with specific values for name and price
	f.Add("Product 1", 100.00)

	f.Fuzz(func(t *testing.T, name string, price float64) {
		product, err := NewProduct(name, price)

		if name == "" || price <= 0.0 {
			assert.NotNil(t, err)
			return
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, product)
			assert.NotNil(t, product.ID)
			assert.Equal(t, name, product.Name)
			assert.Equal(t, price, product.Price)
			assert.NotNil(t, product.CreatedAt)
		}

	})
}

func Test_ProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100.0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func Test_ProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func Test_ProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -100.0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func Test_ProductValidate(t *testing.T) {
	product, _ := NewProduct("Product 1", 100.0)
	err := product.Validate()
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
