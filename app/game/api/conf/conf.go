package conf

import (
	"flag"
	"fmt"
	"os"
	"solo/solo-go-common/database/sql/mysql/orm"
	//"solo/solo-go-common/database/sql/mysql/sql"
	"solo/solo-go-common/net/http"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf *Config
)

// Config config.
type Config struct {
	// db
	//DB *sql.Config
	// orm
	ORM *orm.Config
	// http
	HTTP *http.ServerConfig
}

func init() {
	fmt.Println(os.Getwd())
	flag.StringVar(&confPath, "conf", "conf.toml", "config path")
}

// Init init.
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	fmt.Println(err)
	fmt.Println(Conf)
	return
}
