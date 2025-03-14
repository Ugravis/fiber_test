package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}


var articles = []Article{
	{ Id: 1, Title: "Titre test premier", Body: "Body test premier" },
	{ Id: 2, Title: "Titre test deuxième", Body: "Body test deuxième" },
	{ Id: 3, Title: "Titre test troisième", Body: "Body test troisième" },
}


func main() {
	app := fiber.New()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`Hello world from Fiber`)
	})


	app.Get("/articles", func(c *fiber.Ctx) error {
		return c.JSON(articles)
	})


	app.Get("/articles/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id")) /* Conversion en entier */
		if err != nil {
			return c.Status(400).SendString(`Invalid id format`)
		}

		for _, article := range articles {
			if article.Id == id {
				return c.JSON(article)
			}
		}
		return c.Status(404).SendString(`Article not found`)
	})


	app.Post("/articles", func(c *fiber.Ctx) error {
		var newArticle Article
		if err := c.BodyParser(&newArticle); err != nil {
			return c.Status(404).SendString(`Invalid form data for post article`)
		}

		newArticle.Id = len(articles) + 1
		articles = append(articles, newArticle)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Article %d crée avec succès", newArticle.Id),
			"article": newArticle,
		})
	})


	app.Put("/articles/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("ID invalide")
		}
		var updatedArticle Article
		if err := c.BodyParser(&updatedArticle); err != nil {
			return c.Status(400).SendString((`Invalid form data for put articles`))
		}

		for i, article := range articles {
			if article.Id == id {
				articles[i].Title = updatedArticle.Title
				articles[i].Body = updatedArticle.Body
				return c.JSON(fiber.Map{
					"message": fmt.Sprintf("Article %d modifié avec succès", id),
					"article": articles[i],
				})
			}
		}
		return c.Status(404).SendString(fmt.Sprintf(`Article %d not found`, id))
	})


	app.Delete("/articles/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("ID invalide")
		}

		for i, article := range articles {
			if article.Id == id {
				articles = append(articles[:i], articles[i+1:]...)
				return c.SendString(fmt.Sprintf(`Article %d supprimé avec succès`, id))
			}
		}
		return c.Status(404).SendString(`Article not found`)
	})

	
	app.Listen(":3000")
}