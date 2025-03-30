package api

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetDirectory(c *fiber.Ctx) error {
	dir := "./resource"
	var fileName []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Error(err)
			return err
		}
		if info.IsDir() {
			fileName = append(fileName, path)
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"msg": fileName,
	})
}
