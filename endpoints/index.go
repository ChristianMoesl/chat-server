package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexTemplate struct {
	Messages []string
}

func HandleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.gohtml", &IndexTemplate{Messages: []string{"Hello", "World"}})
}
