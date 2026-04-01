package initialize

import (
	"fmt"
	"go-cover-parroto/global"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres() {
	p := global.Config.Postgres
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Ho_Chi_Minh",
		p.Host, p.Port, p.Username, p.Password, p.DBName, p.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect to PostgreSQL: %w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get sql.DB: %w", err))
	}

	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetime) * time.Second)

	global.DB = db
	global.Logger.Info("PostgreSQL connected")
}