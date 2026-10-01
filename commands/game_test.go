package commands

import (
	"errors"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"math"
	"testing"
)

func TestInvalidPurchasesDoNotTouchDatabase(t *testing.T) {
	for _, count := range []int{-1, 0, 51} {
		if err := purchaseCookies(nil, "user", &consts.CookieItems[0], count, &models.CookieGameUser{}); !errors.Is(err, errInvalidPurchase) {
			t.Fatalf("count %d: %v", count, err)
		}
	}
	if err := purchaseCookies(nil, "user", &consts.CookieItem{Price: math.MaxInt}, 2, &models.CookieGameUser{}); !errors.Is(err, errInvalidPurchase) {
		t.Fatalf("overflow: %v", err)
	}
}

func TestCommandInitializationIsIdempotent(t *testing.T) {
	InitCommand()
	count := len(Commands)
	InitCommand()
	if len(Commands) != count || len(CommandExecutes) != count {
		t.Fatal("duplicate commands after initialization")
	}
}
