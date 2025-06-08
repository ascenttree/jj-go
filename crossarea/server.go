package crossarea

import (
	"fmt"
	"net"

	"github.com/ascenttree/jj-go/common"
)

type CrossareaServer struct {
	Host   string
	Port   uint16
	Logger *common.Logger
}

func (server *CrossareaServer) Serve() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		server.Logger.Error(fmt.Sprintf("Failed to start crossarea server on %s:%d: %v", server.Host, server.Port, err))
		return
	}

	server.Logger.Info(fmt.Sprintf("Crossarea server started on %s:%d", server.Host, server.Port))

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			server.Logger.Error(fmt.Sprintf("Failed to accept connection: %v", err))
			continue
		}

		go server.HandleConnection(conn)
	}
}

func (server *CrossareaServer) HandleConnection(conn net.Conn) {
	defer conn.Close()

	server.Logger.Info(fmt.Sprintf("Accepted connection from %s", conn.RemoteAddr()))

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			server.Logger.Error(fmt.Sprintf("Failed to read from connection: %v", err))
			return
		}

		buffer = buffer[:n]
		server.Logger.Debug(fmt.Sprintf("Received data from %s: %s", conn.RemoteAddr(), common.FormatBytes(buffer)))

		// TODO: figure out how the packets work
		conn.Write(buffer)
	}
}

func NewCrossareaServer(host string, port uint16, logger *common.Logger) *CrossareaServer {
	return &CrossareaServer{
		Host:   host,
		Port:   port,
		Logger: logger,
	}
}
