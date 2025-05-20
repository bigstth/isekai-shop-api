package repository

import (
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/entities"
	_adminException "github.com/bigstth/isekai-shop-api/pkg/admin/exception"
	"github.com/bigstth/isekai-shop-api/pkg/custom"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{db, logger}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err != nil {
		if custom.IsUniqueConstraintError(err, "admin") {
			r.logger.Error("Admin already exists:", err)
			return nil, &_adminException.AdminCreating{}

		}

		r.logger.Error("Failed to create admin:", err)
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}
	return admin, nil
}

func (r *adminRepositoryImpl) FindById(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	if err := r.db.Connect().Table("admins").Where("id = ?", adminID).First(admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warn("Admin not found, ID:", adminID)
			return nil, &_adminException.AdminNotFound{AdminID: adminID}
		}
		r.logger.Error("Failed to query admin:", err)
		return nil, &_adminException.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
