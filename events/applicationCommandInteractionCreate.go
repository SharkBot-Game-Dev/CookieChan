package events

import (
	"github.com/SharkBot-Game-Dev/CookieChan/events/applicationCommandInteractionCreate"
	"github.com/disgoorg/disgo/events"
)

func ApplicationCommandInteractionCreate(e *events.ApplicationCommandInteractionCreate) {
	applicationCommandInteractionCreate.CommandEvent(e)
}
