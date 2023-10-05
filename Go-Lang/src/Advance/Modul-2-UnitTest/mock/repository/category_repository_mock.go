package repository

import (
	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-2-UnitTest/mock/entity"
	"github.com/stretchr/testify/mock"
)

// reposityory -> berperan sebagai jembatan ke database

type CategoryRepositoryMock struct {
	Mock mock.Mock 
	// jadi mock ini adalah pengganti dari 3rd party sistem
	// dimana kita tentunya harus memprogram dulu mocknya
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category{
	argument := repository.Mock.Called(id) 
	// memanggil mock dan tanda bahwa methodnya terpanggil lalu meneruskan value yang ada di parameter 
	// return array dari mock
	if argument.Get(0) == nil { // get -> return index ke-0 dari array mock
		return nil
	}
	category := argument.Get(0).(entity.Category) // convert ke category
	return &category
}

