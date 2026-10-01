package events

import (
	"github.com/SharkBot-Game-Dev/CookieChan/commands"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"log"
)

func AutocompleteInteractionCreate(e *events.AutocompleteInteractionCreate) {
	choices := []discord.AutocompleteChoice{}
	if e.GuildID() != nil && e.Data.CommandName == "buy" && e.Data.Focused().Name == "item" {
		choices = commands.BuyItemChoices(e.Data.String("item"))
	}
	if err := e.AutocompleteResult(choices); err != nil {
		log.Printf("autocomplete response failed: %v", err)
	}
}
