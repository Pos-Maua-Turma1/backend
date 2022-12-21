package repository

import (
	"backend/src/domain"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *Repository) CreateUser(c *fiber.Ctx) error{
	user := domain.User{}

	err := c.BodyParser(&user)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failded"})
		return err
	}

	err = r.DB.Create(&user).Error
	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not Create User"})
		return err
	}

	c.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book has been created"})

	return nil
}

func (r *Repository) DeleteUser(context *fiber.Ctx) error{
	user := domain.User{}
	id := context.Params("id")

	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id cannot be empty"})
		return nil
	}

	err := r.DB.Delete(user, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not Delete User"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "User delete suceffuly"})

	return nil
}

func (r *Repository) GetUserById(context *fiber.Ctx) error{
	user := domain.User{}
	id := context.Params("id")

	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Id cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not search the User"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "User Id fetched suceffuly",
				"data": user,
	})
		
	return nil
}

// func (r *Repository) UpdateUser(context *fiber.Ctx) {}

func (r *Repository) GetUsers(context *fiber.Ctx) error{
	users := &[]domain.User{}

	err := r.DB.Find(users).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not Get List of Users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User Fetched Successfully",
		"data": users,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/user", r.CreateUser)
	api.Delete("/user/:id", r.DeleteUser)
	api.Get("/user/:id", r.GetUserById)
	api.Get("/users", r.GetUsers)
}