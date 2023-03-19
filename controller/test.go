package controller

import (
	"fiber/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func GetTest(c *fiber.Ctx) error {

	maps := &models.Test{}
	c.QueryParser(maps)

	res := models.GetTest(0, 10, maps)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "SUCCESS",
		"data":    res,
		"query":   maps,
	})
}

func AddTest(c *fiber.Ctx) error {
	test := &models.Test{}
	c.BodyParser(test)

	res := models.AddTest(test)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "SUCCESS",
		"data":    res,
	})
}

func EditTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal(err)
	}
	test := &models.Test{}
	c.BodyParser(test)
	code := 200
	res := false
	message := "SUCCESS"

	if models.ExistTestById(id) {
		res = models.EditTest(id, test)
	} else {
		code = 500
	}

	if res {
		message = "SUCCESS"
	} else {
		message = "ERROR"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}

func DelTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal(err)
	}

	code := 200
	res := false
	message := "SUCCESS"

	if models.ExistTestById(id) {
		res = models.DeleteTest(id)
	} else {
		code = 500
	}

	if res {
		message = "SUCCESS"
	} else {
		message = "ERROR"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}
