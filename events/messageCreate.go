package events

import (
	"github.com/SharkBot-Game-Dev/CookieChan/events/messageCreate"
	"github.com/disgoorg/disgo/events"
)

func MessageCreate(e *events.MessageCreate) {
	if e.Message.Author.Bot {
		return
	}

	messageCreate.HelloMessageCreate(e)
	messageCreate.JankenMessageCreate(e)
}
