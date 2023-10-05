package service

import (
	"errors"

	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/entity"
	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func(service CategoryService) Get (id string)(*entity.Category, error){ //business logic
	result := service.Repository.FindById(id)
	if result != nil{
		return result, nil
	}else {
		return nil, errors.New("Category not found")
	}
}