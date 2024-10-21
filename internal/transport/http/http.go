package http

import (
	"authorization_service/internal/application"
	"authorization_service/internal/transport/http/controller"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	server  *echo.Echo
	useCase application.UseCase
}

func New(useCase application.UseCase) *HttpServer {
	return &HttpServer{
		server:  echo.New(),
		useCase: useCase,
	}
}

func (h *HttpServer) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *HttpServer) Register() {
	h.server.Use(middleware.CORS())
	h.server.Use(middleware.BodyLimit("10M"))

	controller := controller.New(h.useCase)

	//
	api := h.server.Group("/api/v1/authorization")

	api.POST("/join", controller.Join)
	api.POST("/confirm_phone", controller.ConfirmPhone)
	api.POST("/register", controller.Register)

	//
}
