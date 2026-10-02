package router

import (
	"github.com/PrathmeshAdhav2006/go-boilerplate/internal/handler"

	"github.com/labstack/echo/v4"
)

// registerSystemRoutes registers the system routes for the application.
func registerSystemRoutes(r *echo.Echo, h *handler.Handlers) {
	r.GET("/status", h.Health.CheckHealth)

	r.Static("/static", "static")

	r.GET("/docs", h.OpenAPI.ServeOpenAPIUI)
}
