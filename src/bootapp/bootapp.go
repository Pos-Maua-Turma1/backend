package bootapp

import (
	"log"
	"os"

	"backend/src/domain"
	"backend/src/repository"
	"backend/src/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func InitApplication () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	configuration := storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(&configuration)
	if err != nil {
		log.Fatal("could noot load the database")
	}

	err = domain.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate the database")
	}

	rep := repository.Repository{
		DB: db,
	}

	app := fiber.New()
	rep.SetupRoutes(app)
	app.Listen(":8080")

}