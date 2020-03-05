package dao

import (
	"github.com/jinzhu/gorm"
	"solo/app/game/api/conf"
	"solo/solo-go-common/database/sql/mysql/orm"
)

// Dao dao
type Dao struct {
	c          *conf.Config
	db         *gorm.DB
	//redis      *redis.Pool
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
		db: orm.NewMySQL(c.ORM),
	}

	return
}
