package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": fileName,
	})
}

func UploadFile(c *fiber.Ctx) error {
	directory := c.FormValue("directory")
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).SendString("uploaded failed")
	}
	logrus.Info(fmt.Sprintf("%s 上傳成功，大小為 %d Bytes", file.Filename, file.Size))
	logrus.Debug(file.Header["Content-Type"])
	// 20MB upper limit
	if file.Size > 20971520 {
		return c.Status(fiber.StatusBadRequest).SendString("uploaded failed")
	}
	directory = strings.ReplaceAll(directory, "\\", "/")
	err = c.SaveFile(file, fmt.Sprintf("./%s/%s", directory, file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("save file failed")
	}
	return c.Status(fiber.StatusOK).SendString("save file successful!")
}
