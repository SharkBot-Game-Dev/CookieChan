package applicationCommandInteractionCreate

import (
	"github.com/SharkBot-Game-Dev/CookieChan/commands"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func CommandEvent(e *events.ApplicationCommandInteractionCreate) {
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

	execute, ok := commands.CommandExecutes[e.Data.CommandName()]
	if !ok {
		e.Client().Rest.CreateInteractionResponse(e.ID(), e.Token(), discord.InteractionResponse{Type: discord.InteractionResponseTypeCreateMessage, Data: discord.MessageCreate{Content: "このコマンドは現在利用できません。", Flags: discord.MessageFlagEphemeral}})
		return
	}
	execute(*e, *e.Client())
}
