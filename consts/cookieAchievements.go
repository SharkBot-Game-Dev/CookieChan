package consts

type CookieAchievement struct {
	ID          string
	Name        string
	Description string
	CookieCount int
}

var CookieAchievements = []CookieAchievement{
	{
		ID:          "beginner",
		Name:        "初心者",
		Description: "所持クッキーが3枚以上になると達成できます。",
		CookieCount: 3,
	},
	{ID: "cookie_snack", Name: "おやつの時間", Description: "所持クッキー25枚以上。", CookieCount: 25},
	{ID: "cookie_baker", Name: "見習いパティシエ", Description: "所持クッキー100枚以上。", CookieCount: 100},
	{ID: "cookie_basket", Name: "クッキーの山", Description: "所持クッキー500枚以上。", CookieCount: 500},
	{ID: "cookie_artisan", Name: "一人前の職人", Description: "所持クッキー2,500枚以上。", CookieCount: 2500},
	{ID: "cookie_shopkeeper", Name: "人気店のオーナー", Description: "所持クッキー10,000枚以上。", CookieCount: 10000},
	{ID: "cookie_tycoon", Name: "クッキー王", Description: "所持クッキー50,000枚以上。", CookieCount: 50000},
	{ID: "cookie_astronaut", Name: "宇宙のパティシエ", Description: "所持クッキー250,000枚以上。", CookieCount: 250000},
	{ID: "cookie_millionaire", Name: "ミリオンクッキー", Description: "所持クッキー1,000,000枚以上。", CookieCount: 1000000},
}

func AchievementNameToStruct(ItemName string) *CookieAchievement {
	for _, item := range CookieAchievements {
		if item.Name == ItemName {
			return &item
		}
	}
	return nil
}

func AchievementIdToStruct(ItemID string) *CookieAchievement {
	for _, item := range CookieAchievements {
		if item.ID == ItemID {
			return &item
		}
	}
	return nil
}
