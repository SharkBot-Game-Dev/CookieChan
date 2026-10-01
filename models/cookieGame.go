package models

import "time"

type CookieGameUser struct {
	ID           uint                    `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time               `json:"created_at"`
	UserId       string                  `gorm:"not null;uniqueIndex" json:"user_id"`
	CookieCount  int                     `json:"cookie_count"`
	Achievements []CookieGameAchievement `gorm:"foreignKey:UserId;references:UserId" json:"achievements"`
	Items        []CookieGameItem        `gorm:"foreignKey:UserId;references:UserId" json:"items"`
}

type CookieGameAchievement struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UserId        string    `gorm:"not null;uniqueIndex:idx_user_achievement" json:"user_id"`
	AchievementID string    `gorm:"not null;uniqueIndex:idx_user_achievement" json:"achievement_id"`
	Status        string    `json:"status"`
	Achieved      bool      `gorm:"not null" json:"achieved"`
}

type CookieGameItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `gorm:"not null" json:"user_id"`
	ItemId    string    `gorm:"not null" json:"item_id"`
	Count     int       `gorm:"not null" json:"count"`
}
