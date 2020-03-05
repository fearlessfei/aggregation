package http

import (
	"github.com/sirupsen/logrus"
	"os"

	"solo/app/game/api/conf"
	"solo/app/game/api/service"
	"solo/solo-go-common/log"
	"solo/solo-go-common/net/http"
)

var srv *service.Service

func Init(c *conf.Config, s *service.Service)  {
	logrus.Infoln()
	logger := logrus.New()

	logger.AddHook(&log.RedirectHook{})

	//logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",//时间格式化

	})

	logger.SetOutput(os.Stdout)

	srv = s

	router := InitRouter()
	engine := http.NewServer(c.HTTP, router)
	engine.Start()
}
