package commands

import (
	"errors"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"gorm.io/gorm"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"

	"strconv"
)

var CookieCommand = discord.SlashCommandCreate{
	Name:        "cookie",
	Description: "現在のステータスを取得します。",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionUser{
			Name:        "user",
			Description: "ステータスを表示するユーザー",
			Required:    false,
		},
	},
}

func CookieCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	})

	var cookieUser models.CookieGameUser

	user, userOk := interaction.SlashCommandInteractionData().OptUser("user")
	if !userOk {
		user = interaction.User()
	}

	cookieUserResult := consts.DB.First(&cookieUser, "user_id = ?", user.ID.String())
	if errors.Is(cookieUserResult.Error, gorm.ErrRecordNotFound) {
		if cookieUserResult.Error != nil {
			client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "そのユーザーはクッキーを持っていません。"})
			return true
		}
	} else {
		if cookieUserResult.Error != nil {
			client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "内部DBエラーが発生しました。"})
			return true
		}
	}

	achievementText := ""
	for _, achievement := range cookieUser.Achievements {
		if achievement.Achieved {
			achievemntSt := consts.AchievementIdToStruct(achievement.AchievementID)
			achievementText += (achievemntSt.Name + "\n")
		}
	}
	if achievementText == "" {
		achievementText = "まだ入手していません。"
	}

	itemText := ""
	for _, item := range cookieUser.Items {
		if item.Count <= 0 {
			itemSt := consts.ItemIdToStruct(item.ItemId)
			itemText += (itemSt.Name + " (" + strconv.Itoa(item.Count) + "個)" + "\n")
		}
	}
	if itemText == "" {
		itemText = "まだ入手していません。"
	}

	helpEmbed := discord.Embed{
		Title: "😆" + *user.GlobalName + "のステータス",
		Color: consts.CookieColor,
		Fields: []discord.EmbedField{
			{
				Name:   "🍪クッキー数",
				Value:  strconv.Itoa(cookieUser.CookieCount) + "クッキー",
				Inline: &consts.False,
			},
			{
				Name:   "🍮所持アイテム",
				Value:  itemText,
				Inline: &consts.False,
			},
			{
				Name:   "🏅達成済み実績",
				Value:  achievementText,
				Inline: &consts.False,
			},
		},
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
