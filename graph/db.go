package graph

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/tabed23/movie-graphql/graph/model"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"localhost", "5432", "admin", "admin123", "movies_graphql", "disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil
	}

	db.AutoMigrate(&model.Movie{})
	return db
}
