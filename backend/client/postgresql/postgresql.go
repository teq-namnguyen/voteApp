package postgresql

import (
	"context"
	"fmt"
	"github.com/namnguyen/backend/utils"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/namnguyen/backend/config"
)

var db *gorm.DB

func init() {
	var (
		err error
		cfg = config.GetConfig()
	)

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Postgresql.Host,
		cfg.Postgresql.User,
		cfg.Postgresql.Pass,
		cfg.Postgresql.DBName,
		cfg.Postgresql.Port,
	)

	db, err = gorm.Open(postgres.New(postgres.Config{DSN: connectionString}), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db = db.Debug()

	rawDB, _ := db.DB()
	rawDB.SetMaxIdleConns(cfg.Postgresql.DBMaxIdleConns)
	rawDB.SetMaxOpenConns(cfg.Postgresql.DBMaxOpenConns)
	rawDB.SetConnMaxLifetime(time.Minute * 5)

	err = rawDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetClient(ctx context.Context) *gorm.DB {
	if utils.IsEnableTx(ctx) {
		return utils.GetTx(ctx)
	}
	return db.Session(&gorm.Session{})
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	if db != nil {
		sqlDB.Close()
	}
}
