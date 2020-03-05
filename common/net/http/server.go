package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type ServerConfig struct {
	Addr            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxHeaderBytes  int
}

type Engine struct {
	Server *http.Server
	Conf *ServerConfig
}

func NewServer(c *ServerConfig, router *gin.Engine) (engine *Engine) {

	s := &http.Server{
		Addr:           c.Addr,
		Handler:        router,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxHeaderBytes: c.MaxHeaderBytes,
	}

	engine = &Engine{
		Server: s,
		Conf: c,
	}
	return
}

func (e *Engine) Start()  {
	fmt.Printf("Runing at %s\n", e.Conf.Addr)
	go func() {
		if err := e.Server.ListenAndServe(); err != nil {
			err = errors.Wrapf(err, "addrs: %v", e.Conf.Addr)
		}
	}()
}
