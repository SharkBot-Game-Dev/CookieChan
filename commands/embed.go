package commands

import (
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/omit"
)

var EmbedCommand = discord.SlashCommandCreate{
	Name:                     "embed",
	Description:              "埋め込みメッセージを作成します。",
	DefaultMemberPermissions: omit.NewPtr(discord.PermissionManageChannels),
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "title",
			Description: "埋め込みのタイトル",
			Required:    true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "description",
			Description: "埋め込みの説明",
			Required:    true,
		},
	},
}

func EmbedCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	if err := client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	}); err != nil {
		return true
	}

	title := interaction.SlashCommandInteractionData().String("title")
	description := interaction.SlashCommandInteractionData().String("description")

	Embed := discord.Embed{
		Title:       title,
		Color:       consts.CookieColor,
		Description: description,
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{Embed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	return true
}
