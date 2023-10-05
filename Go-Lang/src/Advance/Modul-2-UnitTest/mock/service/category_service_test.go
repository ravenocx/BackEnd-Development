package service

import (
	"testing"

	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/entity"
	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock : mock.Mock{}} // creating template 3rd party system
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindById","1").Return(nil) 
	// testing, ketika call FindById with parameter id =1 harus mengembalikan return apa (expected)
	// sama seperti test case (expected outputnya apa)

	category,err := categoryService.Get("1")
	assert.Nil(t, category) // expected category -> nil
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// result, err := categoryService.Get("2")
	// assert.Nil(t, err)
	// assert.NotNil(t, result)
	// assert.Equal(t, category.Id, result.Id)
	// assert.Equal(t, category.Name, result.Name)
}