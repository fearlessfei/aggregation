package service

import (
	"solo/app/game/api/conf"
	"solo/app/game/api/dao"
)

// Service struct
type Service struct {
	c   *conf.Config
	dao *dao.Dao
}


// New service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}

	return
}
