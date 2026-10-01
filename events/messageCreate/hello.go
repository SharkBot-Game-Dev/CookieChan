package messageCreate

import (
	"math/rand"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func HelloMessageCreate(e *events.MessageCreate) {
	if strings.Contains(e.Message.Content, "おは") {
		randomGoodMorning := []string{"おはよ～\n起きてる～？", "おはよー！", "おはようございます", "おっはよー！"}
		e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: randomGoodMorning[rand.Intn(len(randomGoodMorning))]})
		return
	}

	if strings.Contains(e.Message.Content, "こんに") {
		randomGoodMorning := []string{"こんにちは！", "こんにちは～！", "こんちは～", "こんにちは！！"}
		e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: randomGoodMorning[rand.Intn(len(randomGoodMorning))]})
		return
	}

	if strings.Contains(e.Message.Content, "こんば") {
		randomGoodMorning := []string{"こんばんは！", "こんばんは～！", "こんばんは～", "こんばんは～\nまだ起きてるの？"}
		e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: randomGoodMorning[rand.Intn(len(randomGoodMorning))]})
		return
	}

	if strings.Contains(e.Message.Content, "挨拶を無効化") {
		e.Client().Rest.CreateMessage(e.ChannelID, discord.MessageCreate{Content: "挨拶を無効化したいチャンネルを、\nクッキーちゃんが表示できなくすると、\n無効化できるよ！"})
		return
	}
}
