package commands

import (
	"errors"
	"strconv"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"gorm.io/gorm"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"
)

var ClickCommand = discord.SlashCommandCreate{
	Name:        "click",
	Description: "クッキーをクリックします。",
}

func ClickCommandExecute(
	interaction events.ApplicationCommandInteractionCreate,
	client bot.Client,
) bool {
	client.Rest.CreateInteractionResponse(
		interaction.ID(),
		interaction.Token(),
		discord.InteractionResponse{
			Type: discord.InteractionResponseTypeDeferredCreateMessage,
		},
	)

	userID := interaction.User().ID.String()

	var cookieUser models.CookieGameUser

	result := consts.DB.
		Preload("Achievements").
		Preload("Items").
		First(&cookieUser, "user_id = ?", userID)

	clickCount := 1
	cookieCount := 0

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		cookieUser = models.CookieGameUser{
			UserId:      userID,
			CookieCount: 1,
		}

		if err := consts.DB.Create(&cookieUser).Error; err != nil {
			client.Rest.CreateFollowupMessage(
				client.ApplicationID,
				interaction.Token(),
				discord.MessageCreate{
					Content: "内部DBエラーが発生しました。",
				},
			)
			return true
		}

		cookieCount = 1

	} else if result.Error != nil {
		client.Rest.CreateFollowupMessage(
			client.ApplicationID,
			interaction.Token(),
			discord.MessageCreate{
				Content: "内部DBエラーが発生しました。",
			},
		)
		return true

	} else {
		for _, item := range cookieUser.Items {
			itemData := consts.ItemIdToStruct(item.ItemId)

			clickCount += itemData.ClickCount * item.Count
		}

		cookieCount = cookieUser.CookieCount + clickCount

		if err := consts.DB.
			Model(&cookieUser).
			Update("cookie_count", cookieCount).
			Error; err != nil {

			client.Rest.CreateFollowupMessage(
				client.ApplicationID,
				interaction.Token(),
				discord.MessageCreate{
					Content: "内部DBエラーが発生しました。",
				},
			)

			return true
		}

		cookieUser.CookieCount = cookieCount
	}

	clickEmbed := discord.Embed{
		Title: "🖱️クッキーをクリックしました！",
		Description: strconv.Itoa(clickCount) +
			"クッキーを入手しました。\n現在は" +
			strconv.Itoa(cookieCount) +
			"クッキー持っています。",
		Color: consts.CookieColor,
	}

	client.Rest.CreateFollowupMessage(
		client.ApplicationID,
		interaction.Token(),
		discord.MessageCreate{
			Embeds: []discord.Embed{clickEmbed},
		},
	)

	for _, achievement := range consts.CookieAchievements {
		if achievement.CookieCount > cookieCount {
			continue
		}

		alreadyAchieved := false

		for _, userAchievement := range cookieUser.Achievements {
			if userAchievement.AchievementID == achievement.ID &&
				userAchievement.Achieved {

				alreadyAchieved = true
				break
			}
		}

		if alreadyAchieved {
			continue
		}

		newAchievement := models.CookieGameAchievement{
			UserId:        userID,
			AchievementID: achievement.ID,
			Status:        "achieved",
			Achieved:      true,
		}

		if err := consts.DB.Create(&newAchievement).Error; err != nil {
			continue
		}

		cookieUser.Achievements = append(
			cookieUser.Achievements,
			newAchievement,
		)

		client.Rest.CreateMessage(
			interaction.Channel().ID(),
			discord.MessageCreate{
				Content: "やったー！新しい実績を達成したよ！\n「" +
					achievement.Name +
					"」\n/cookieで確認してみてね！",
			},
		)
	}

	return true
}
