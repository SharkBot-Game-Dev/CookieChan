package events

import (
	"github.com/SharkBot-Game-Dev/CookieChan/events/interactionCreate"
	"github.com/disgoorg/disgo/events"
)

func InteractionCreate(e *events.InteractionCreate) {
	interactionCreate.CommandEvent(e)
}
