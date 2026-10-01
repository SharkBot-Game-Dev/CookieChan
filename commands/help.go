package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var HelpCommand = discord.SlashCommandCreate{
	Name:        "help",
	Description: "このBotの使い方を表示します。",
}

func HelpCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	if err := client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	}); err != nil {
		return true
	}

	helpEmbed := discord.Embed{
		Title:       "クッキーちゃんの使い方",
		Description: "クッキーちゃんは主に\nスラッシュコマンドでやりとりします。",
		Color:       consts.CookieColor,
		Fields: []discord.EmbedField{
			{
				Name:   "基本",
				Value:  "`/help`",
				Inline: &consts.False,
			},
			{
				Name:   "ゲーム",
				Value:  "`/janken`",
				Inline: &consts.False,
			},
			{
				Name:   "クッキーゲーム",
				Value:  "`/cookie` `/click` `/buy` `/shop`",
				Inline: &consts.False,
			},
			{
				Name:   "挨拶",
				Value:  "`おはよう` `こんにちは` `こんばんは` `挨拶を無効化`",
				Inline: &consts.False,
			},
		},
		Footer: &discord.EmbedFooter{
			Text: "/から始まるコマンドはスラッシュコマンドです。",
		},
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
