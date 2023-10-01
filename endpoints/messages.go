package endpoints

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/ChristianMoesl/chat-server/chat"
)

type MessageTemplate struct {
	Message chat.Message
}

type MessagesEndpoint struct {
	Database *sqlx.DB
}

func (e *MessagesEndpoint) HandlePost(c echo.Context) error {
	message := chat.Message{
		Text: c.FormValue("message"),
	}
	_, err := e.Database.NamedExec("INSERT INTO `messages` (`message`) VALUES (:message)", message)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "message.gohtml", &MessageTemplate{Message: message})
}
