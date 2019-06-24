package main

import (
	"os"

	database "gin-example/db"
	middleware "gin-example/lib/middlewares"
	router "gin-example/routes"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default()

	//c.Set("db")
	app.Use(database.Inject(db))
	app.Use(middleware.CORS())

	router.RootRoutes(app)

	app.Run(":" + port)
}
