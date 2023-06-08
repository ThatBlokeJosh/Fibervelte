package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var test string

func main() {
	//template render engine
	engine := html.New("./dist", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})

  app.Use(func(c *fiber.Ctx) error {
    fmt.Println(test)
    return c.Next()
  })
  
  app.Static("/", "./dist")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
    "Test": "test",
  }); })

  app.Post("/", func(c *fiber.Ctx) error {
    var body struct {
      Name string
    } 
    if err := c.BodyParser(&body); err != nil {
      return err
    }
    test = body.Name
		return c.Render("index", nil); 
  })

	app.Listen(":8080")
}

