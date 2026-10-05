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
	{Name: "月面クッキー農園", Description: "月の大地で材料を育てます。1個につきクリックで得るクッキーが250,000増えます。", Price: 425000, ID: "moon_cookie_farm", ClickCount: 250000},
	{Name: "惑星ベーカリー", Description: "惑星まるごと焼き上げます。1個につきクリックで得るクッキーが500,000増えます。", Price: 825000, ID: "planet_bakery", ClickCount: 500000},
	{Name: "恒星オーブン", Description: "星の熱で香ばしく焼きます。1個につきクリックで得るクッキーが1,000,000増えます。", Price: 1600000, ID: "stellar_oven", ClickCount: 1000000},
	{Name: "銀河クッキー便", Description: "銀河中の焼きたてを集めます。1個につきクリックで得るクッキーが2,500,000増えます。", Price: 3875000, ID: "galaxy_cookie_delivery", ClickCount: 2500000},
	{Name: "ブラックホール倉庫", Description: "光もクッキーも逃しません。1個につきクリックで得るクッキーが5,000,000増えます。", Price: 7500000, ID: "black_hole_storage", ClickCount: 5000000},
	{Name: "次元クッキーゲート", Description: "別の次元から焼きたてが届きます。1個につきクリックで得るクッキーが10,000,000増えます。", Price: 14500000, ID: "dimension_cookie_gate", ClickCount: 10000000},
	{Name: "時間逆行こね機", Description: "昨日の生地をもう一度焼きます。1個につきクリックで得るクッキーが25,000,000増えます。", Price: 35000000, ID: "time_reversal_kneader", ClickCount: 25000000},
	{Name: "平行世界ベーカリー", Description: "いくつもの世界で同時に焼きます。1個につきクリックで得るクッキーが50,000,000増えます。", Price: 67500000, ID: "parallel_world_bakery", ClickCount: 50000000},
	{Name: "量子クッキー複製機", Description: "量子のゆらぎでおやつを増やします。1個につきクリックで得るクッキーが100,000,000増えます。", Price: 130000000, ID: "quantum_cookie_copier", ClickCount: 100000000},
	{Name: "夢のクッキー工場", Description: "夢で見たお菓子を焼き上げます。1個につきクリックで得るクッキーが200,000,000増えます。", Price: 250000000, ID: "dream_cookie_factory", ClickCount: 200000000},
	{Name: "星雲シュガーミル", Description: "星雲を甘い砂糖に変えます。1個につきクリックで得るクッキーが300,000,000増えます。", Price: 360000000, ID: "nebula_sugar_mill", ClickCount: 300000000},
	{Name: "宇宙創造オーブン", Description: "新しい宇宙とおやつを焼きます。1個につきクリックで得るクッキーが500,000,000増えます。", Price: 575000000, ID: "universe_creation_oven", ClickCount: 500000000},
	{Name: "無限レシピ図書館", Description: "尽きないレシピを焼き続けます。1個につきクリックで得るクッキーが750,000,000増えます。", Price: 825000000, ID: "infinite_recipe_library", ClickCount: 750000000},
	{Name: "クッキーの世界樹", Description: "枝いっぱいにクッキーが実ります。1個につきクリックで得るクッキーが1,000,000,000増えます。", Price: 1050000000, ID: "cookie_world_tree", ClickCount: 1000000000},
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
