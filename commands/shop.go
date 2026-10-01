package commands

import (
	"strconv"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var ShopCommand = discord.SlashCommandCreate{
	Name:        "shop",
	Description: "クッキーで買えるアイテムを表示します。",
}

func ShopCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	})

	Fields := []discord.EmbedField{}

	for _, item := range consts.CookieItems {
		Fields = append(Fields, discord.EmbedField{
			Name:  item.Name + " (" + strconv.Itoa(item.Price) + "クッキー)",
			Value: item.Description,
		})
	}

	helpEmbed := discord.Embed{
		Title:  "クッキーショップ",
		Color:  consts.CookieColor,
		Fields: Fields,
		Footer: &discord.EmbedFooter{
			Text: "/buy コマンドで買えます。",
		},
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
