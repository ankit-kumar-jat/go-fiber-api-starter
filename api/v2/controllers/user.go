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

func GetUserController(c *fiber.Ctx) error {
    param := struct {ID uint `params:"id"`}{}

    if err := c.ParamsParser(&param); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    // get user from db

	DB := db.GetDB()

    result := DB.Where("Id = ?", param.ID)
    if result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.JSON(result)
}