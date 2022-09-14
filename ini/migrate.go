package ini

import "github.com/tarcea/go-auth-jwt/models"

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
