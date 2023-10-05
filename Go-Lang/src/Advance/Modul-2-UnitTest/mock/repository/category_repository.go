package repository

import "github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}