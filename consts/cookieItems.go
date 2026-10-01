package consts

type CookieItem struct {
	Name        string
	Description string
	Price       int // 売値は2分の1
	ID          string
	ClickCount  int
}

var CookieItems = []CookieItem{
	{
		Name:        "ダブルクリック",
		Description: "通常のクリックに加えて1回クリックできるようにします。",
		Price:       10,
		ID:          "double_click",
		ClickCount:  1,
	},
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
