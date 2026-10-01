package interactionCreate

import (
	"github.com/SharkBot-Game-Dev/CookieChan/commands"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func CommandEvent(e *events.InteractionCreate) {
	switch i := e.Interaction.(type) {
	case discord.ApplicationCommandInteraction:
		if e.GuildID() == nil {
			e.Client().Rest.CreateInteractionResponse(e.ID(), e.Token(), discord.InteractionResponse{
				Type: discord.InteractionResponseTypeCreateMessage,
				Data: discord.MessageCreate{
					Content: "スラッシュコマンドはDMで使用できません。",
					Flags:   discord.MessageFlagEphemeral,
				},
			})
			return
		}

		commands.CommandExecutes[i.Data.CommandName()](e, *e.Client())
	}

}
