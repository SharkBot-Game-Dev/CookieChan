package consts

import (
	"log"

	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Print("DBに接続しました。")

	initDB()
}

func initDB() {
	if err := migrateDB(DB); err != nil {
		log.Fatal(err)
	}
}

// Migrate the old numeric parent IDs before GORM changes the column types.
// PostgreSQL rolls back both the schema and data if any step fails.
func migrateDB(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("SELECT pg_advisory_xact_lock(731940128)").Error; err != nil {
			return err
		}
		for _, relation := range []struct{ table, constraint string }{
			{"cookie_game_achievements", "fk_cookie_game_users_achievements"},
			{"cookie_game_items", "fk_cookie_game_users_items"},
		} {
			var legacy bool
			if err := tx.Raw(`SELECT EXISTS (
				SELECT 1 FROM pg_constraint c
				JOIN pg_attribute a ON a.attrelid = c.confrelid AND a.attnum = ANY(c.confkey)
				WHERE c.conrelid = to_regclass(?) AND c.conname = ?
				AND c.contype = 'f' AND c.confrelid = to_regclass('cookie_game_users') AND a.attname = 'id'
			)`, relation.table, relation.constraint).Scan(&legacy).Error; err != nil {
				return err
			}
			if !legacy {
				continue
			}
			// Identifiers below come only from the fixed list above.
			if err := tx.Exec(`ALTER TABLE "` + relation.table + `" DROP CONSTRAINT "` + relation.constraint + `"`).Error; err != nil {
				return err
			}
			if err := tx.Exec(`ALTER TABLE "` + relation.table + `" ALTER COLUMN user_id TYPE text USING user_id::text`).Error; err != nil {
				return err
			}
			if err := tx.Exec(`UPDATE "` + relation.table + `" child SET user_id = parent.user_id
				FROM cookie_game_users parent WHERE child.user_id = parent.id::text`).Error; err != nil {
				return err
			}
		}
		return tx.AutoMigrate(
			&models.CookieGameUser{},
			&models.CookieGameAchievement{},
			&models.CookieGameItem{},
		)
	})
}
