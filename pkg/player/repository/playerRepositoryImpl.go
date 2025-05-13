package repository

import (
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_playerException "github.com/bigstth/isekai-shop-api/pkg/player/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{db, logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		if custom.IsUniqueConstraintError(err, "player") {
			r.logger.Error("Player already exists:", err)
			return nil, &_playerException.PlayerCreating{}

		}

		r.logger.Error("Failed to create player:", err)
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindById(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	if err := r.db.Connect().Table("players").Where("id = ?", playerID).First(player).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warn("Player not found, ID:", playerID)
			return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
		}
		r.logger.Error("Failed to query player:", err)
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
