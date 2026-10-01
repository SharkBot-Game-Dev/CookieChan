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
		Description: "3回クリックすると達成できます。",
		CookieCount: 3,
	},
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
