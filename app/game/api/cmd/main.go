package main

import (
	"flag"
	"os"
	"os/signal"
	"solo/app/game/api/conf"
	"solo/app/game/api/server/http"
	"solo/app/game/api/server/tcp"
	"solo/app/game/api/service"
	"syscall"
)

func main()  {
	flag.Parse()

	// Init config
	err := conf.Init()
	if err != nil {
		panic("Init config error")
	}

	// New service
	s := service.New(conf.Conf)

	// Init http server
	http.Init(conf.Conf, s)

	// Init tcp server
	tcp.Init(conf.Conf, s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}