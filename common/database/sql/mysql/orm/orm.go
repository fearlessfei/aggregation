package orm

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"solo/solo-go-common/log"
)

// Config mysql orm config.
type Config struct {
	DSN         string         // data source name.
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout time.Duration  // connect max life time.
}

//type ormLog struct{}
//
//func (l ormLog) Print(v ...interface{}) {
//	log.Log.Infoln(v...)
//}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	db.LogMode(true)
	//db.SetLogger(ormLog{})

	return
}
