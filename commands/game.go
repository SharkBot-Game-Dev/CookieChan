package commands

import (
	"errors"
	"math"
	"time"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var errInsufficientCookies = errors.New("insufficient cookies")
var errInvalidPurchase = errors.New("invalid purchase")
var errCookieOverflow = errors.New("cookie count overflow")
var errClickCooldown = errors.New("click cooldown active")

const clickCooldown = 10 * time.Minute

// Lock the user for every balance/inventory change so concurrent commands serialize.
func purchaseCookies(db *gorm.DB, userID string, item *consts.CookieItem, count int, user *models.CookieGameUser) error {
	if item == nil || count < 1 || count > 50 || item.Price < 0 || item.Price > math.MaxInt/count {
		return errInvalidPurchase
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		cost := item.Price * count
		if user.CookieCount < cost {
			return errInsufficientCookies
		}
		var owned models.CookieGameItem
		err := tx.First(&owned, "user_id = ? AND item_id = ?", userID, item.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			owned = models.CookieGameItem{UserId: userID, ItemId: item.ID, Count: count}
			if err := tx.Create(&owned).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			if owned.Count > math.MaxInt-count {
				return errCookieOverflow
			}
			if err := tx.Model(&owned).Update("count", owned.Count+count).Error; err != nil {
				return err
			}
		}
		return tx.Model(user).Update("cookie_count", user.CookieCount-cost).Error
	})
}

func clickCookies(db *gorm.DB, userID string) (user models.CookieGameUser, gain int, err error) {
	gain = 1
	err = db.Transaction(func(tx *gorm.DB) error {
		// A conflict-safe insert also serializes simultaneous first clicks.
		if err := tx.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "user_id"}}, DoNothing: true}).Create(&models.CookieGameUser{UserId: userID}).Error; err != nil {
			return err
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		now := time.Now()
		if now.Before(user.ClickCooldown) {
			return errClickCooldown
		}
		var items []models.CookieGameItem
		if err := tx.Where("user_id = ?", userID).Find(&items).Error; err != nil {
			return err
		}
		for _, owned := range items {
			item := consts.ItemIdToStruct(owned.ItemId)
			if item == nil || owned.Count <= 0 || item.ClickCount <= 0 {
				continue
			}
			if owned.Count > (math.MaxInt-gain)/item.ClickCount {
				return errCookieOverflow
			}
			gain += item.ClickCount * owned.Count
		}
		if user.CookieCount > math.MaxInt-gain {
			return errCookieOverflow
		}
		user.CookieCount += gain
		user.ClickCooldown = now.Add(clickCooldown)
		return tx.Model(&user).Updates(map[string]any{
			"cookie_count":   user.CookieCount,
			"click_cooldown": user.ClickCooldown,
		}).Error
	})
	return
}

func unlockAchievement(db *gorm.DB, userID, achievementID string) (bool, error) {
	unlocked := false
	err := db.Transaction(func(tx *gorm.DB) error {
		var user models.CookieGameUser
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "user_id = ?", userID).Error; err != nil {
			return err
		}
		var achievement models.CookieGameAchievement
		err := tx.First(&achievement, "user_id = ? AND achievement_id = ?", userID, achievementID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			achievement = models.CookieGameAchievement{UserId: userID, AchievementID: achievementID, Status: "achieved", Achieved: true}
			if err := tx.Create(&achievement).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else if achievement.Achieved {
			return nil
		} else if err := tx.Model(&achievement).Updates(map[string]any{"status": "achieved", "achieved": true}).Error; err != nil {
			return err
		}
		unlocked = true
		return nil
	})
	return unlocked, err
}
