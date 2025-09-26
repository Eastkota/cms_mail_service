package main

import (
	schema "cms_mail_service/graph"
	"cms_mail_service/handlers"
	"cms_mail_service/helpers"
	"cms_mail_service/repositories"
	"cms_mail_service/resolvers"
	"cms_mail_service/services"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := helpers.GetGormDB()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	EmailLogRepository := repositories.NewEmailLogRepository(db)
	EmailLogService := services.NewEmailLogService(EmailLogRepository)
	resolver := resolvers.NewEmailLogResolver(EmailLogService)

	queryType := schema.NewQueryType(resolver)
	mutationType := schema.NewMutationType(resolver)

	schema.InitSchema(queryType, mutationType)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, 
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	e.POST("/graphql", handlers.Handler)
	e.Logger.Fatal(e.Start(":8095"))
}
