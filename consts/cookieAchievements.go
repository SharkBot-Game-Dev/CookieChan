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
	{ID: "cookie_vault", Name: "おやつの大金庫", Description: "所持クッキー1,500,000枚以上。", CookieCount: 1500000},
	{ID: "cookie_double_million", Name: "ダブルミリオン", Description: "所持クッキー2,000,000枚以上。", CookieCount: 2000000},
	{ID: "cookie_moon_collector", Name: "月まで届く山", Description: "所持クッキー3,000,000枚以上。", CookieCount: 3000000},
	{ID: "cookie_planet_owner", Name: "お菓子の惑星", Description: "所持クッキー5,000,000枚以上。", CookieCount: 5000000},
	{ID: "cookie_comet", Name: "クッキー彗星", Description: "所持クッキー7,500,000枚以上。", CookieCount: 7500000},
	{ID: "cookie_ten_million", Name: "千万枚の焼きたて", Description: "所持クッキー10,000,000枚以上。", CookieCount: 10000000},
	{ID: "cookie_stellar_baker", Name: "恒星の焼き菓子職人", Description: "所持クッキー15,000,000枚以上。", CookieCount: 15000000},
	{ID: "cookie_space_fleet", Name: "おやつの宇宙艦隊", Description: "所持クッキー20,000,000枚以上。", CookieCount: 20000000},
	{ID: "cookie_galaxy_collector", Name: "銀河の収集家", Description: "所持クッキー30,000,000枚以上。", CookieCount: 30000000},
	{ID: "cookie_galaxy_ruler", Name: "銀河クッキー王", Description: "所持クッキー50,000,000枚以上。", CookieCount: 50000000},
	{ID: "cookie_black_hole", Name: "底なしのおやつ箱", Description: "所持クッキー75,000,000枚以上。", CookieCount: 75000000},
	{ID: "cookie_hundred_million", Name: "一億枚の祝宴", Description: "所持クッキー100,000,000枚以上。", CookieCount: 100000000},
	{ID: "cookie_dimension_traveler", Name: "次元を渡る収集家", Description: "所持クッキー150,000,000枚以上。", CookieCount: 150000000},
	{ID: "cookie_time_keeper", Name: "時を越えるおやつ", Description: "所持クッキー200,000,000枚以上。", CookieCount: 200000000},
	{ID: "cookie_parallel_collector", Name: "平行世界の大富豪", Description: "所持クッキー250,000,000枚以上。", CookieCount: 250000000},
	{ID: "cookie_quantum_master", Name: "量子クッキーマスター", Description: "所持クッキー300,000,000枚以上。", CookieCount: 300000000},
	{ID: "cookie_dream_collector", Name: "夢いっぱいのおやつ箱", Description: "所持クッキー400,000,000枚以上。", CookieCount: 400000000},
	{ID: "cookie_half_billion", Name: "五億枚の大収穫", Description: "所持クッキー500,000,000枚以上。", CookieCount: 500000000},
	{ID: "cookie_nebula_collector", Name: "星雲を包む甘さ", Description: "所持クッキー600,000,000枚以上。", CookieCount: 600000000},
	{ID: "cookie_universe_collector", Name: "宇宙いっぱいのクッキー", Description: "所持クッキー750,000,000枚以上。", CookieCount: 750000000},
	{ID: "cookie_billionaire", Name: "ビリオンクッキー", Description: "所持クッキー1,000,000,000枚以上。", CookieCount: 1000000000},
	{ID: "cookie_world_tree_keeper", Name: "世界樹のおやつ守り", Description: "所持クッキー1,250,000,000枚以上。", CookieCount: 1250000000},
	{ID: "cookie_legendary_collector", Name: "伝説のクッキー収集家", Description: "所持クッキー1,500,000,000枚以上。", CookieCount: 1500000000},
	{ID: "cookie_eternal_collector", Name: "永遠のおやつ時間", Description: "所持クッキー2,000,000,000枚以上。", CookieCount: 2000000000},
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
