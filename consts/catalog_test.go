package consts

import "testing"

func TestCatalogProgression(t *testing.T) {
	ids, names := map[string]bool{}, map[string]bool{}
	for i, item := range CookieItems {
		if item.ID == "" || item.Name == "" || item.Description == "" || ids[item.ID] || names[item.Name] || item.Price <= 0 || item.ClickCount <= 0 {
			t.Fatalf("invalid item: %+v", item)
		}
		ids[item.ID], names[item.Name] = true, true
		if i > 0 {
			prev := CookieItems[i-1]
			if item.Price <= prev.Price || item.ClickCount <= prev.ClickCount || item.Price*prev.ClickCount >= prev.Price*item.ClickCount {
				t.Fatalf("tier %s should cost more and offer better efficiency", item.ID)
			}
		}
	}
	ids = map[string]bool{}
	for i, achievement := range CookieAchievements {
		if achievement.ID == "" || achievement.Name == "" || achievement.Description == "" || ids[achievement.ID] || achievement.CookieCount <= 0 {
			t.Fatalf("invalid achievement: %+v", achievement)
		}
		ids[achievement.ID] = true
		if i > 0 && achievement.CookieCount <= CookieAchievements[i-1].CookieCount {
			t.Fatal("thresholds must increase")
		}
	}
	if CookieItems[0].ID != "double_click" || CookieItems[0].Price != 5 || CookieAchievements[0].ID != "beginner" || CookieAchievements[0].CookieCount != 3 {
		t.Fatal("existing progression changed")
	}
}
