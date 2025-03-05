package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/common/utils"
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/database"
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/models"
	post "github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/post/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	defer database.Close()

	database.DATABASE.AutoMigrate(
		&models.User{},
		&post.Post{},
	)

	// Reading sample data
	var users []models.User
	var jsonPath string = "./03-postgres-db/fake/fake-data.json"
	err = utils.ReadJsonFile(jsonPath, &users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		result := database.DATABASE.Create(&user)
		fmt.Println("User ID: ", user.ID)
		fmt.Println("Posts: ", user.Posts)
		fmt.Println("Error: ", result.Error)
		fmt.Println("Rows affected: ", result.RowsAffected)
	}
}
