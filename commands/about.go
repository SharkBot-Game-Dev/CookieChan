package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var AboutCommand = discord.SlashCommandCreate{
	Name:        "about",
	Description: "Botの情報を表示します。",
}

func AboutCommandExecute(
	interaction events.ApplicationCommandInteractionCreate,
	client bot.Client,
) bool {
	if err := client.Rest.CreateInteractionResponse(
		interaction.ID(),
		interaction.Token(),
		discord.InteractionResponse{
			Type: discord.InteractionResponseTypeDeferredCreateMessage,
		},
	); err != nil {
		return true
	}

	omikujiEmbed := discord.Embed{
		Title:       "🤖クッキーちゃんについて",
		Description: "クッキーちゃんは、便利な多機能Botです。",
		Color:       consts.CookieColor,
		Fields: []discord.EmbedField{
			{
				Name:  "バージョン",
				Value: "v1.0.1",
			},
		},
	}

	client.Rest.CreateFollowupMessage(
		client.ApplicationID,
		interaction.Token(),
		discord.MessageCreate{
			Embeds: []discord.Embed{omikujiEmbed},
		},
	)
	return true
}
