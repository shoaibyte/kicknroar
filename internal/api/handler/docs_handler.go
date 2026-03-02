package handler

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/labstack/echo/v4"
)

// DocsHandler handles documentation routes
type DocsHandler struct{}

// NewDocsHandler creates a new docs handler
func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

// ScalarUI serves the Scalar API reference UI
// @Summary      API Documentation
// @Description  Interactive API documentation using Scalar
// @Tags         docs
// @Produce      text/html
// @Success      200  {string}  string  "HTML content"
// @Router       /docs [get]
func (h *DocsHandler) ScalarUI(c echo.Context) error {
	// Construct the full URL for the OpenAPI spec
	scheme := "http"
	if c.Request().TLS != nil {
		scheme = "https"
	}
	host := c.Request().Host
	specURL := scheme + "://" + host + "/api/v1/swagger/doc.json"

	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: specURL,
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Kick&Roar API Documentation",
		},
		DarkMode: true,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate documentation: "+err.Error())
	}

	return c.HTML(http.StatusOK, htmlContent)
}
