package main

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/link"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/stat"
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	// Миграция включает дневную статистику и её связь со ссылками.
	if err := db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{}); err != nil {
		panic(err)
	}
}
