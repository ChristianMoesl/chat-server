package main

import (
	"html/template"
	"io"
	"log/slog"
	"os"

	"github.com/ChristianMoesl/chat-server/database"
	"github.com/ChristianMoesl/chat-server/endpoints"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	logger := createLogger()
	logger.Info("starting server...")

	e := echo.New()
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}
	e.Renderer = renderer

	database, err := database.Connect()
	if err != nil {
		e.Logger.Fatal(err)
	}

	indexEndpoint := &endpoints.IndexEndpoint{Database: database}
	messagesEndpoint := &endpoints.MessagesEndpoint{Database: database}
	e.GET("/", indexEndpoint.HandleIndex)
	e.POST("/messages", messagesEndpoint.HandlePost)

	port, is_port_defined := os.LookupEnv("PORT")
	if !is_port_defined {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func createLogger() *slog.Logger {
	options := slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				return slog.Attr{Key: "severity", Value: a.Value}
			case slog.MessageKey:
				return slog.Attr{Key: "message", Value: a.Value}
			default:
				return a
			}
		},
	}

	var handler slog.Handler
	if os.Getenv("MODE") == "development" {
		handler = slog.NewTextHandler(os.Stdout, &options)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &options)
	}
	return slog.New(handler)
}
