package controllers

import (
	"example.com/go-api-starter/internal/db"
	"example.com/go-api-starter/internal/models"

	"github.com/gofiber/fiber/v2"
)

type AddUserRequestBody struct {
    Name  string `json:"name"`
    Email string  `json:"email"`
}

// Add user 
// @Description Add user using given info.
// @Summary Add user
// @Tags User
// @Accept json
// @Produce json
// @Param body body AddUserRequestBody true "Name"
// @Success 200 {object} models.User
// @Router /api/v1/users/ [post]
func AddUserController(c *fiber.Ctx) error {
    body := AddUserRequestBody{}

    // parse body, attach to AddUserRequestBody struct
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    var user models.User

    user.Name = body.Name
    user.Email = body.Email

    // insert new db entry

	DB := db.GetDB()

    if result := DB.Create(&user); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(&user)
}

// Get user by Id or 404 error
// @Description Get user using Id.
// @Summary Get user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /api/v1/users/{id} [get]
func GetUserController(c *fiber.Ctx) error {
    id := c.Params("id")

    var product models.User

	DB := db.GetDB()

    if result := DB.First(&product, id); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusOK).JSON(&product)
}

// Get all users
// @Description Get all users.
// @Summary Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Router /api/v1/users/ [get]
func GetUsersController(c *fiber.Ctx) error {
    var products []models.User

	DB := db.GetDB()
    if result := DB.Find(&products); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusOK).JSON(&products)
}