package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"

	"math/rand"
)

var OmikujiCommand = discord.SlashCommandCreate{
	Name:        "omikuji",
	Description: "おみくじを引きます。",
}

func OmikujiCommandExecute(
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

	omikuji := []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶"}

	index := rand.Intn(len(omikuji))

	omikujiEmbed := discord.Embed{
		Title:       "🥠おみくじを引きました。",
		Description: omikuji[index] + "です。\nまた明日挑戦してね！",
		Color:       consts.CookieColor,
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
