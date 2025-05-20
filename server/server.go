package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bigstth/isekai-shop-api/config"
	"github.com/bigstth/isekai-shop-api/databases"
	_adminRepository "github.com/bigstth/isekai-shop-api/pkg/admin/repository"
	_oauth2Controller "github.com/bigstth/isekai-shop-api/pkg/oauth2/controller"
	_oauth2Service "github.com/bigstth/isekai-shop-api/pkg/oauth2/service"
	_playerRepository "github.com/bigstth/isekai-shop-api/pkg/player/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
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
	timeOutMiddleware := getTimeoutMiddleWare(s.conf.Server.Timeout)
	corsMiddleware := getCORSMiddleWare(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleWare(s.conf.Server.BodyLimit)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeOutMiddleware)

	authorizingMiddleware := s.getAuthorizingMiddleware()

	s.app.GET("/V1/health", s.healthCheck)
	s.initOAuth2Router()
	s.initItemShopRouter()
	s.initItemManagingRouter(authorizingMiddleware)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefullyShutdown(quitCh)
	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal) {
	ctx := context.Background()
	<-quitCh
	s.app.Logger.Info("Shutting down server...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getTimeoutMiddleWare(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request timed out",
		Timeout:      timeout * time.Second,
	})
}

func getCORSMiddleWare(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleWare(limit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(limit)
}

func (s *echoServer) getAuthorizingMiddleware() *authorizingMiddleware {

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(oauth2Service, s.conf.OAuth2, s.app.Logger)

	return &authorizingMiddleware{
		oauth2Controller: oauth2Controller,
		oauth2Config:     s.conf.OAuth2,
		logger:           s.app.Logger,
	}
}
