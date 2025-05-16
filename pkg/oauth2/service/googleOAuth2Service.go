package service

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_adminModel "github.com/bigstth/isekai-shop-api/pkg/admin/model"
	_adminRepository "github.com/bigstth/isekai-shop-api/pkg/admin/repository"
	_playerModel "github.com/bigstth/isekai-shop-api/pkg/player/model"
	_playerRepository "github.com/bigstth/isekai-shop-api/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) *googleOAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingRequest *_playerModel.PlayerCreatingReq) error {
	if !s.IsPlayerExist(playerCreatingRequest.ID) {

		playerEntity := &entities.Player{
			ID:     playerCreatingRequest.ID,
			Name:   playerCreatingRequest.Name,
			Email:  playerCreatingRequest.Email,
			Avatar: playerCreatingRequest.Avatar,
		}
		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingRequest *_adminModel.AdminCreatingReq) error {
	if !s.IsAdminExist(adminCreatingRequest.ID) {

		adminEntity := &entities.Admin{
			ID:     adminCreatingRequest.ID,
			Name:   adminCreatingRequest.Name,
			Email:  adminCreatingRequest.Email,
			Avatar: adminCreatingRequest.Avatar,
		}
		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) IsPlayerExist(playerID string) bool {
	playerEntity, err := s.playerRepository.FindById(playerID)
	if err != nil {
		return false
	}

	return playerEntity != nil
}

func (s *googleOAuth2Service) IsAdminExist(adminID string) bool {
	adminEntity, err := s.adminRepository.FindById(adminID)
	if err != nil {
		return false
	}

	return adminEntity != nil
}
