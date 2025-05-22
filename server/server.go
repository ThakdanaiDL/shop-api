package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	echoApp := echo.New()              // echo.New() คือการสร้าง instance ของ echo
	echoApp.Logger.SetLevel(log.DEBUG) // echoApp.Logger.SetLevel(log.DEBUG) คือการตั้งค่า log level ของ echo ให้เป็น DEBUG

	once.Do(func() {
		server = &echoServer{ // server คือ instance ของ echoServer
			app:  echoApp,
			db:   db,
			conf: conf,
		}

	})

	return server

}

func (s *echoServer) Start() {

	s.app.GET("/v1/health", s.healthCheck)

	//***********grace full shutdown***********//
	quiteCH := make(chan os.Signal, 1)                      //core-->  quiteCH คือ channel ที่ใช้ในการรับ signal ที่ส่งมาจาก os
	signal.Notify(quiteCH, syscall.SIGINT, syscall.SIGTERM) //core-->  signal.Notify(quiteCH, os.Interrupt) คือการรอรับ signal ที่ส่งมาจาก os
	go s.gracefulshutdown(quiteCH)                          //core-->  go routine
	//***********grace full shutdown***********//

	s.httplistening() // core-->  s.httplistening() คือการเริ่มต้น server โดยใช้ url ที่กำหนด

}

func (s *echoServer) httplistening() {

	url := fmt.Sprintf("0.0.0.0:%d", s.conf.Server.Port)                    //buid url string
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed { //core-->  app.Start(url) คือการเริ่มต้น server โดยใช้ url ที่กำหนด
		s.app.Logger.Fatal("Shutting down the server")
	}

}

func (s *echoServer) gracefulshutdown(quiteCH chan os.Signal) {

	ctx := context.Background() //core-->  ctx คือการสร้าง context เพื่อใช้ในการรอรับ signal ที่ส่งมาจาก channel quite

	<-quiteCH //core-->  <- quite คือการรอรับ signal ที่ส่งมาจาก channel quite
	s.app.Logger.Info("gracefull shut down the server")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal("gracefull shut down the server")
	}

}

func (s *echoServer) healthCheck(e echo.Context) error { //endpoint health check
	return e.String(http.StatusOK, "OK")
}
