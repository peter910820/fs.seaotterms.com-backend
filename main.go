package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"fs.seaotterms.com-backend/internal/router"
)

var (
	frontendFolder string = "./dist"
)

func init() {
	os.MkdirAll("./resource", os.ModePerm)
	os.MkdirAll("./resource/image", os.ModePerm)
	os.MkdirAll("./resource/test", os.ModePerm)
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	// logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	// Create a new engine
	engine := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/public", "./public")
	app.Static("/resource", "./resource")

	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:8080",
		AllowMethods: "POST"}))

	// static folder
	app.Static("/", frontendFolder)

	// route group
	apiGroup := app.Group("/api") // main api route group
	// register router
	router.ApiRouter(apiGroup) // check identity for front-end routes

	app.Get("/folder", func(c *fiber.Ctx) error {
		dir := "./resource"
		fileName := []string{}
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logrus.Error(err)
				return err
			}
			if !info.IsDir() {
				fileName = append(fileName, path)
			}
			return nil
		})
		if err != nil {
			logrus.Error(err)
			return c.Status(fiber.StatusInternalServerError).SendString("server has an error")
		}
		return c.Render("folder", fiber.Map{
			"FileName": fileName,
		}, "layouts/base")
	})

	app.Get("/text-editor", func(c *fiber.Ctx) error {
		return c.Render("textEditor", nil, "layouts/base")
	})

	/* --------------------------------- */
	// match all routes
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile(frontendFolder + "/index.html")
	})

	logrus.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
