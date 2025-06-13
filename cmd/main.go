package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"creditninja/internal/handlers"
	"creditninja/internal/models"
)

func main() {
	// Load env
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Init DB
	db, err := models.ConnectDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	// Fiber with Go html/template engine
	engine := html.New("./internal/templates", ".html")
	// Add custom template functions
	engine.AddFunc("default", func(value interface{}, fallback interface{}) interface{} {
		switch v := value.(type) {
		case string:
			if v == "" {
				return fallback
			}
		case int, int64, float64:
			if v == 0 {
				return fallback
			}
		case nil:
			return fallback
		}
		return value
	})
	// Reload templates on each render (useful for development)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(logger.New())

	// Sessions
	store := session.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("store", store)
		c.Locals("db", db)
		return c.Next()
	})

	// Routes
	handlers.RegisterRoutes(app)

	// Static
	app.Static("/static", "./static")

	log.Printf("Listening on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
