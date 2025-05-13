package repository

import "github.com/bigstth/isekai-shop-api/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindById(adminId string) (*entities.Admin, error)
}
