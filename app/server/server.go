package server

import (
	"fmt"
	"net/http"

	"github.com/blockfint/di-example-go/app/handler"
	"github.com/blockfint/di-example-go/app/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	e      *echo.Echo
	h      *handler.Handler
	logger *zap.Logger
}

func (s Server) Serve() error {
	HTTP_SERVER_PORT := viper.GetInt("HTTP_SERVER.PORT")
	servePath := fmt.Sprintf(":%d", HTTP_SERVER_PORT)

	s.logger.Sugar().Fatal(s.e.Start(servePath))

	return nil
}

func New(h *handler.Handler, lg *zap.Logger) *Server {
	e := echo.New()

	e.Validator = handler.NewRequestValidator()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut},
	}))

	e.Use(logger.EchoMiddleware(lg))

	api := e.Group("/api")
	v1 := api.Group("/v1")
	h.Register(v1)

	return &Server{e, h, lg}
}
