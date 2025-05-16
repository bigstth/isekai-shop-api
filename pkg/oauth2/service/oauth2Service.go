package service

import (
	_adminModel "github.com/bigstth/isekai-shop-api/pkg/admin/model"
	_playerModel "github.com/bigstth/isekai-shop-api/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingRequest *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingRequest *_adminModel.AdminCreatingReq) error
	IsPlayerExist(playerID string) bool
	IsAdminExist(adminID string) bool
}
