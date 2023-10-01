package endpoints

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/ChristianMoesl/chat-server/chat"
)

type IndexTemplate struct {
	Messages []chat.Message
}

type IndexEndpoint struct {
	Database *sqlx.DB
}

func (e *IndexEndpoint) HandleIndex(c echo.Context) error {
	messages := []chat.Message{}
	err := e.Database.Select(&messages, "SELECT * FROM `messages`")
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "index.gohtml", &IndexTemplate{Messages: messages})
}
