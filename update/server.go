package update

import (
	"fmt"
	"io"

	"github.com/ascenttree/jj-go/common"
	"github.com/gin-gonic/gin"
)

type UpdateServer struct {
	Host   string
	Port   uint16
	Logger *common.Logger
}

type UpdateContext struct {
	*gin.Context
	Server *UpdateServer
}

func (server *UpdateServer) Serve() {
	server.Logger.Info(fmt.Sprintf("Starting update server on %s:%d", server.Host, server.Port))

	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
	router := gin.New()
	jjGroup := router.Group("/others/rockrock/")

	jjGroup.GET("/version_info/latest-version.xml", server.WithContext(LatestVersionHandler))

	if err := router.Run(fmt.Sprintf("%s:%d", server.Host, server.Port)); err != nil {
		server.Logger.Error(fmt.Sprintf("Failed to start update server: %v", err))
	}
}

func (server *UpdateServer) WithContext(handler func(*UpdateContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.Logger.Debug(fmt.Sprintf("Handling request: %s %s", ctx.Request.Method, ctx.Request.URL.Path))

		context := &UpdateContext{
			Context: ctx,
			Server:  server,
		}

		handler(context)
	}
}

func NewUpdateServer(host string, port uint16, logger *common.Logger) *UpdateServer {
	return &UpdateServer{
		Host:   host,
		Port:   port,
		Logger: logger,
	}
}
