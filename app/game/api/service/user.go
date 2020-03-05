package service

import (
	"context"
	"solo/app/game/api/models"
)

func (s *Service) User(c context.Context) []models.User {
	users := s.dao.User(c)
	return users
}
