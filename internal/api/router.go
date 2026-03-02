package api

import (
	"io/fs"
	"net/http"
	"strings"

	_ "kicknroar/docs" // swagger docs
	"kicknroar/internal/api/handler"
	"kicknroar/internal/api/middleware"
	"kicknroar/internal/config"
	"kicknroar/internal/pkg/jwt"
	"kicknroar/web"

	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
)

// SetupRouter configures all routes
func SetupRouter(
	cfg *config.Config,
	jwtManager *jwt.Manager,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	matchHandler *handler.MatchHandler,
	venueHandler *handler.VenueHandler,
	uploadHandler *handler.UploadHandler,
	docsHandler *handler.DocsHandler,
) *echo.Echo {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS(&cfg.Server))

	// Embedded SPA filesystem (web/dist); serve at /
	contentFS, _ := fs.Sub(web.DistFS, "dist")

	// Health check
	// @Summary      Health check
	// @Description  Check if the API is running
	// @Tags         health
	// @Produce      json
	// @Success      200  {object}  map[string]string
	// @Router       /api/v1/health [get]
	e.GET("/api/v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Documentation routes
	e.GET("/docs", docsHandler.ScalarUI)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Swagger documentation (must be after v1 group is created)
	v1.GET("/swagger/*", echoswagger.WrapHandler)

	// Auth routes (with auth rate limiting)
	auth := v1.Group("/auth")
	auth.Use(middleware.AuthRateLimit(&cfg.RateLimit))
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
	}

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.Auth(jwtManager))
	protected.Use(middleware.RateLimit(&cfg.RateLimit))

	// User routes
	users := protected.Group("/users")
	{
		users.GET("/me", userHandler.GetMe)
		users.PUT("/me", userHandler.UpdateMe)
		users.GET("/:id", userHandler.GetUser)
		users.GET("/:id/stats", userHandler.GetUserStats)
	}

	// Match routes
	matches := protected.Group("/matches")
	{
		matches.GET("", matchHandler.ListMatches)
		matches.POST("", matchHandler.CreateMatch)
		matches.GET("/:id", matchHandler.GetMatch)
		matches.PUT("/:id", matchHandler.UpdateMatch)
		matches.DELETE("/:id", matchHandler.DeleteMatch)
		matches.POST("/:id/join", matchHandler.JoinMatch)
		matches.POST("/:id/leave", matchHandler.LeaveMatch)
		matches.GET("/:id/participants", matchHandler.GetParticipants)
	}

	// Venue routes
	venues := protected.Group("/venues")
	{
		venues.GET("", venueHandler.ListVenues)
		venues.POST("", venueHandler.CreateVenue)
		venues.GET("/nearby", venueHandler.FindNearby)
		venues.GET("/:id", venueHandler.GetVenue)
		venues.PUT("/:id", venueHandler.UpdateVenue)
	}

	// Upload routes (with upload rate limiting)
	upload := protected.Group("/upload")
	upload.Use(middleware.UploadRateLimit(&cfg.RateLimit))
	{
		upload.POST("/avatar", uploadHandler.UploadAvatar)
	}

	// Venue upload route
	venues.POST("/:id/upload", venueHandler.UploadVenueImage)

	// Serve embedded SPA: static assets and SPA fallback (index.html for client-side routes)
	e.GET("/*", spaHandler(contentFS))

	return e
}

// spaHandler serves files from the embedded FS and falls back to index.html for SPA client-side routing.
func spaHandler(contentFS fs.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := strings.TrimPrefix(c.Request().URL.Path, "/")
		if path == "" {
			path = "index.html"
		}
		f, err := contentFS.Open(path)
		if err == nil {
			defer f.Close()
			info, _ := f.Stat()
			if info.IsDir() {
				path = strings.TrimSuffix(path, "/") + "/index.html"
				f.Close()
				f, err = contentFS.Open(path)
				if err != nil {
					return serveIndex(c, contentFS)
				}
				defer f.Close()
			}
			return c.Stream(http.StatusOK, getContentType(path), f)
		}
		return serveIndex(c, contentFS)
	}
}

func serveIndex(c echo.Context, contentFS fs.FS) error {
	index, err := contentFS.Open("index.html")
	if err != nil {
		return c.String(http.StatusNotFound, "index.html not found")
	}
	defer index.Close()
	return c.Stream(http.StatusOK, "text/html; charset=utf-8", index)
}

func getContentType(path string) string {
	switch {
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".ico"):
		return "image/x-icon"
	case strings.HasSuffix(path, ".png"):
		return "image/png"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	case strings.HasSuffix(path, ".json"):
		return "application/json"
	case strings.HasSuffix(path, ".woff2"):
		return "font/woff2"
	case strings.HasSuffix(path, ".woff"):
		return "font/woff"
	default:
		return "application/octet-stream"
	}
}
