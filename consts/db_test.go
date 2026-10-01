package consts

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLegacyForeignKeyMigration(t *testing.T) {
	dsn := os.Getenv("COOKIECHAN_TEST_DSN")
	if dsn == "" {
		t.Skip("COOKIECHAN_TEST_DSN is required for PostgreSQL migration test")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()
	// Roll back the isolated schema, including on test failure.
	rollback := fmt.Errorf("rollback test schema")
	err = db.Transaction(func(tx *gorm.DB) error {
		name := fmt.Sprintf("cookiechan_test_%d", time.Now().UnixNano())
		if err := tx.Exec(`CREATE SCHEMA "` + name + `"`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`SET LOCAL search_path TO "` + name + `"`).Error; err != nil {
			return err
		}
		statements := []string{
			`CREATE TABLE cookie_game_users (id bigserial PRIMARY KEY, created_at timestamptz, user_id text NOT NULL UNIQUE, cookie_count bigint)`,
			`CREATE TABLE cookie_game_achievements (id bigserial PRIMARY KEY, created_at timestamptz, user_id bigint NOT NULL, achievement_id text NOT NULL, status text, achieved boolean NOT NULL, CONSTRAINT fk_cookie_game_users_achievements FOREIGN KEY(user_id) REFERENCES cookie_game_users(id))`,
			`CREATE TABLE cookie_game_items (id bigserial PRIMARY KEY, created_at timestamptz, user_id bigint NOT NULL, item_id text NOT NULL, count bigint NOT NULL, CONSTRAINT fk_cookie_game_users_items FOREIGN KEY(user_id) REFERENCES cookie_game_users(id))`,
			`INSERT INTO cookie_game_users(id,user_id,cookie_count) VALUES(1,'123456789012345678',20)`,
			`INSERT INTO cookie_game_achievements(user_id,achievement_id,status,achieved) VALUES(1,'beginner','achieved',true)`,
			`INSERT INTO cookie_game_items(user_id,item_id,count) VALUES(1,'double_click',2)`,
		}
		for _, statement := range statements {
			if err := tx.Exec(statement).Error; err != nil {
				return err
			}
		}
		for range 2 {
			if err := migrateDB(tx); err != nil {
				return err
			}
		}
		var user models.CookieGameUser
		if err := tx.Preload("Items").Preload("Achievements").First(&user, "user_id = ?", "123456789012345678").Error; err != nil {
			return err
		}
		if user.CookieCount != 20 || len(user.Items) != 1 || user.Items[0].Count != 2 || len(user.Achievements) != 1 || !user.Achievements[0].Achieved {
			return fmt.Errorf("migration lost data: %+v", user)
		}
		return rollback
	})
	if err != rollback {
		t.Fatal(err)
	}
}
