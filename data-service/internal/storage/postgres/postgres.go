package postgres

import (
	"data-service/internal/entity"
	extraClausePlugin "github.com/WinterYukky/gorm-extra-clause-plugin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	cnf := &gorm.Config{}

	cnf.Logger = logger.Default

	db, err := gorm.Open(postgres.Open(dsn), cnf)
	if err != nil {
		return nil, err
	}

	err = db.Use(extraClausePlugin.New())
	if err != nil {
		return nil, err
	}

	models := []any{
		&entity.Film{},
		&entity.Review{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

	return db, nil
}
