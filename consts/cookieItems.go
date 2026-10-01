package consts

type CookieItem struct {
	Name        string
	Description string
	Price       int
	ID          string
	ClickCount  int
}

var CookieItems = []CookieItem{
	{
		Name:        "ダブルクリック",
		Description: "通常のクリックに加えて1回クリックできるようにします。",
		Price:       5,
		ID:          "double_click",
		ClickCount:  1,
	},
	{Name: "クッキーミキサー", Description: "生地をまとめて混ぜます。1個につきクリックで得るクッキーが10増えます。", Price: 40, ID: "cookie_mixer", ClickCount: 10},
	{Name: "小さなオーブン", Description: "焼きたてをお届け。1個につきクリックで得るクッキーが100増えます。", Price: 300, ID: "small_oven", ClickCount: 100},
	{Name: "クッキー工房", Description: "職人と一緒に製造。1個につきクリックで得るクッキーが1,000増えます。", Price: 2500, ID: "cookie_workshop", ClickCount: 1000},
	{Name: "クッキー工場", Description: "大量生産を始めます。1個につきクリックで得るクッキーが10,000増えます。", Price: 20000, ID: "cookie_factory", ClickCount: 10000},
	{Name: "宇宙クッキー基地", Description: "宇宙へおいしさを広げます。1個につきクリックで得るクッキーが100,000増えます。", Price: 175000, ID: "space_cookie_base", ClickCount: 100000},
}

func ItemNameToStruct(ItemName string) *CookieItem {
	for _, item := range CookieItems {
		if item.Name == ItemName {
			return &item
		}
	}
	return nil
}

func ItemIdToStruct(ItemID string) *CookieItem {
	for _, item := range CookieItems {
		if item.ID == ItemID {
			return &item
		}
	}
	return nil
}
