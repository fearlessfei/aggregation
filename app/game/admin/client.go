package admin

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"

	"github.com/leesper/holmes"
	"solo/solo-go-common/net/tcp"
)

const (
	// ChatMessage is the message number of chat message.
	ChatMessage int32 = 1
)

// Message defines the chat message.
type Message struct {
	Content string
}

// MessageNumber returns the message number.
func (cm Message) MessageNumber() int32 {
	return ChatMessage
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tcp.Message, err error) {
	if data == nil {
		return nil, tcp.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn tcp.WriteCloser) {
	holmes.Infof("ProcessMessage")
	s, ok := tcp.ServerFromContext(ctx)
	if ok {
		msg := tcp.MessageFromContext(ctx)
		s.Broadcast(msg)
	}
}

func main() {
	defer holmes.Start().Stop()

	tcp.Register(ChatMessage, DeserializeMessage, nil)

	c, err := net.Dial("tcp", "127.0.0.1:12346")
	if err != nil {
		holmes.Fatalln(err)
	}

	onConnect := tcp.OnConnectOption(func(c tcp.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := tcp.OnErrorOption(func(c tcp.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := tcp.OnCloseOption(func(c tcp.WriteCloser) {
		holmes.Infoln("on close")
	})

	onMessage := tcp.OnMessageOption(func(msg tcp.Message, c tcp.WriteCloser) {
		fmt.Print(msg.(Message).Content)
	})

	options := []tcp.ServerOption{
		onConnect,
		onError,
		onClose,
		onMessage,
		tcp.ReconnectOption(),
	}

	conn := tcp.NewClientConn(0, c, options...)
	defer conn.Close()

	conn.Start()
	for {
		reader := bufio.NewReader(os.Stdin)
		talk, _ := reader.ReadString('\n')
		if talk == "bye\n" {
			break
		} else {
			msg := Message{
				Content: talk,
			}
			if err := conn.Write(msg); err != nil {
				holmes.Infoln("error", err)
			}
		}
	}
	fmt.Println("goodbye")
}
