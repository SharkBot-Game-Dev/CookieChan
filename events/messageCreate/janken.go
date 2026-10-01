package messageCreate

import (
	"slices"

	"github.com/SharkBot-Game-Dev/CookieChan/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func determineWinner(userChoice, computerChoice string) string {
	if userChoice == computerChoice {
		return "あいこ"
	}

	switch userChoice {
	case "ぐー":
		if computerChoice == "ちょき" {
			return "勝ち"
		}
		return "負け"
	case "ちょき":
		if computerChoice == "ぱー" {
			return "勝ち"
		}
		return "負け"
	case "ぱー":
		if computerChoice == "ぐー" {
			return "勝ち"
		}
		return "負け"
	}

	return "不明"
}

func JankenMessageCreate(e *events.MessageCreate) {
	valid := slices.Contains([]string{"ぐー", "ちょき", "ぱー"}, e.Message.Content)
	answer, active := cache.PlayJanken(e.Message.Author.ID, e.ChannelID, valid)
	if !active {
		return
	}

	if !slices.Contains([]string{"ぐー", "ちょき", "ぱー"}, e.Message.Content) {
		e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: "ぐー、ちょき、ぱーのどれかを入力してね！"})
		return
	}

	e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: "結果は" + determineWinner(e.Message.Content, answer) + "だよ！\nまた遊んでね！"})

}
