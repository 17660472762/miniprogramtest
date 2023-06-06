package server

import (
	"github.com/17660472762/miniprogramtest/pkg/server/message"
	"github.com/17660472762/miniprogramtest/pkg/server/user"
	"github.com/17660472762/miniprogramtest/pkg/wrap"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine     *gin.Engine
	messagesvc message.Service
	usersvc    user.Service
}

type Builder struct {
	svr *Server
}

func NewBuilder(appid, secret string) *Builder {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &Builder{
		svr: &Server{
			engine:  engine,
			usersvc: user.NewService(appid, secret),
		},
	}

}
func (builder *Builder) Router() *Builder {
	builder.svr.router()
	return builder
}

func (s *Server) router() {
	s.engine.GET("/api/v1/menssage/receive", s.messagesvc.Receive)
	s.engine.GET("/api/v1/user/get", wrap.Wrap(s.usersvc.GetUser))
	s.engine.GET("/api/v1/user/access_token", wrap.Wrap(s.usersvc.GetAccessToken))
}

func (builder *Builder) Build() *Server {
	return builder.svr
}

func (s *Server) Run() error {
	return s.engine.Run()
}
