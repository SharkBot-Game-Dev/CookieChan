package commands

import (
	"math/rand"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"

	"github.com/SharkBot-Game-Dev/CookieChan/cache"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var JankenCommand = discord.SlashCommandCreate{
	Name:        "janken",
	Description: "私とじゃんけんするよ！",
}

func JankenCommandExecute(interaction discord.Interaction, client bot.Client) bool {
	client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	})

	if cache.JankenSessions[interaction.User().ID] == interaction.Channel().ID() {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{
			Content: "ゲームを始めたばっかりだよ！\nぐー、ちょき、ぱーのどれかを入力してね！",
		})
		return true
	}

	cache.JankenSessions[interaction.User().ID] = interaction.Channel().ID()

	answers := []string{"ぐー", "ちょき", "ぱー"}

	cache.JankenAnswer[interaction.User().ID] = answers[rand.Intn(len(answers))]

	helpEmbed := discord.Embed{
		Title:       "じゃんけん！",
		Description: "ぐー、ちょき、ぱーのどれかを入力してね！",
		Color:       consts.CookieColor,
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
