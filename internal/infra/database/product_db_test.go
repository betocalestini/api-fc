package database

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/betocalestini/api-fc/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

}
func Test_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 0; i < 25; i++ {
		product, err := entity.NewProduct("Product "+strconv.Itoa(i), rand.Float64()*100)
		assert.NoError(t, err)

		productDB := NewProduct(db)
		err = productDB.Create(product)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.ID)
	}

}

func Test_FindByID(t *testing.T) {

}

func Test_Update(t *testing.T) {

}

func Test_Delete(t *testing.T) {

}
