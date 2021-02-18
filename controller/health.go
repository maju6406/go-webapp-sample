package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maju6406/go-webapp-sample/mycontext"
)

// HealthController is a controller returns the current status of this application.
type HealthController struct {
	context mycontext.Context
}

// NewHealthController is constructor.
func NewHealthController(context mycontext.Context) *HealthController {
	return &HealthController{context: context}
}

// GetHealthCheck returns whether this application is alive or not.
func (controller *HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
