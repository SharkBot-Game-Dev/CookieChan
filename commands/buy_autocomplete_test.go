package commands

import (
	"fmt"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/disgoorg/disgo/discord"
	"strings"
	"testing"
)

func TestBuyItemChoices(t *testing.T) {
	for _, query := range []string{"ミキサー", " COOKIE_MIXER "} {
		choices := BuyItemChoices(query)
		if len(choices) != 1 {
			t.Fatalf("%q: got %d choices", query, len(choices))
		}
		choice := choices[0].(discord.AutocompleteChoiceString)
		item := consts.ItemNameToStruct(choice.Value)
		if item == nil || item.ID != "cookie_mixer" || !strings.Contains(choice.Name, "40🍪") || !strings.Contains(choice.Name, "+10") {
			t.Fatalf("invalid choice: %+v", choice)
		}
	}
	if len(BuyItemChoices("存在しないアイテム")) != 0 {
		t.Fatal("unexpected match")
	}
	if len(BuyItemChoices("")) != len(consts.CookieItems) {
		t.Fatal("empty query should list catalog")
	}
}

func TestBuyItemChoicesLimit(t *testing.T) {
	original := consts.CookieItems
	defer func() { consts.CookieItems = original }()
	consts.CookieItems = nil
	for i := range 30 {
		consts.CookieItems = append(consts.CookieItems, consts.CookieItem{Name: fmt.Sprintf("item%d", i), ID: fmt.Sprintf("id%d", i), Price: 5, ClickCount: 1})
	}
	if len(BuyItemChoices("")) != 25 {
		t.Fatal("autocomplete limit exceeded")
	}
}
