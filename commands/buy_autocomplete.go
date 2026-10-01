package commands

import (
	"fmt"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/disgoorg/disgo/discord"
	"strings"
)

func BuyItemChoices(query string) []discord.AutocompleteChoice {
	query = strings.ToLower(strings.TrimSpace(query))
	choices := make([]discord.AutocompleteChoice, 0, 25)
	for _, item := range consts.CookieItems {
		if !strings.Contains(strings.ToLower(item.Name), query) && !strings.Contains(strings.ToLower(item.ID), query) {
			continue
		}
		choices = append(choices, discord.AutocompleteChoiceString{
			Name:  fmt.Sprintf("%s｜%d🍪｜クリック +%d", item.Name, item.Price, item.ClickCount),
			Value: item.Name,
		})
		if len(choices) == 25 {
			break
		}
	}
	return choices
}
