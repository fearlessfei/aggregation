package tcp

import (
	"github.com/leesper/holmes"
	"net"
	"solo/app/game/api/conf"
	"solo/app/game/api/service"
	"solo/solo-go-common/net/tcp"
)

var srv *service.Service

type TCPServer struct {
	*tcp.Server
	conf *conf.Config
}

// NewTCPServer returns NewTCPServer.
func NewTCPServer(c *conf.Config) *TCPServer {
	onConnect := tcp.OnConnectOption(func(conn tcp.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := tcp.OnErrorOption(func(conn tcp.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := tcp.OnCloseOption(func(conn tcp.WriteCloser) {
		holmes.Infoln("closing pingpong client")
	})

	return &TCPServer{
		tcp.NewServer(onConnect, onError, onClose),
		c,
	}
}

func Init(c *conf.Config, s *service.Service) {
	defer holmes.Start().Stop()

	tcp.Register(ChatMessage, DeserializeMessage, ProcessMessage)

	go func() {
		srv = s

		server := NewTCPServer(c)

		l, err := net.Listen("tcp", ":12346")
		if err != nil {
			holmes.Fatalln("listen error", err)
		}

		err = server.Start(l)
		if err != nil {
			holmes.Fatalln("start tcp server error", err)
		}
	}()
}
