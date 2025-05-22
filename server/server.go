package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/ThakdanaiDL.git/shop-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {

	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}

	})

	return server

}

func (s *echoServer) Start() {

	s.app.GET("/v1/health", s.healthCheck)

	s.httplistening()

}

func (s *echoServer) httplistening() {

	url := fmt.Sprintf("0.0.0.0:%d", s.conf.Server.Port)
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal("Shutting down the server")
	}

}

func (s *echoServer) healthCheck(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}
