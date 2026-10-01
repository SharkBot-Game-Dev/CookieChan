package commands

import (
	"log"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var HelpCommand = discord.SlashCommandCreate{
	Name:        "help",
	Description: "このBotの使い方を表示します。",
}

func HelpCommandExecute(interaction discord.Interaction, client bot.Client) bool {
	client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	})

	helpEmbed := discord.Embed{
		Title:       "クッキーちゃんの使い方",
		Description: "クッキーちゃんは主に\nスラッシュコマンドでやりとりします。",
		Color:       consts.CookieColor,
		Fields: []discord.EmbedField{
			{
				Name:   "基本コマンド",
				Value:  "`/help`",
				Inline: &consts.False,
			},
			{
				Name:   "挨拶コマンド",
				Value:  "`おはよう` `こんにちは`",
				Inline: &consts.False,
			},
		},
		Footer: &discord.EmbedFooter{
			Text: "/から始まるコマンドはスラッシュコマンドです。",
		},
	}

	helpEmbed.AddField("挨拶コマンド", "`おはよう` `こんにちは`", false)

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	_, err := client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	log.Print(err)

	return true
}
