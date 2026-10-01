package interactionCreate

import (
	"github.com/SharkBot-Game-Dev/CookieChan/commands"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func CommandEvent(e *events.InteractionCreate) {
	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		commands.CommandExecutes[i.Data.CommandName()](e, *e.Client())
	}

}
