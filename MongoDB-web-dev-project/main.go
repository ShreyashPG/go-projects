package main

import (
	"log"
	"os"
		"fmt"
		"github.com/ShreyashPG/MongoDB-web-dev-project/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ShreyashPG/MongoDB-web-dev-project/controllers"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	
	app := fiber.New()
	app.Use(cors.New())
	
	database.ConnectToDB()
	app.Get("/", controllers.GetAllEmployees)
	app.Get("/:id", controllers.GetEmployeeByID)
	app.Post("/", controllers.CreateEmployee)
	app.Put("/:id", controllers.UpdateEmployee)
	app.Delete("/:id", controllers.DeleteEmployee)


	port := os.Getenv("PORT")
	fmt.Println("Server is running on port", port)
	fmt.Println("MongoDB URI:", os.Getenv("MONGO_URI"))
	log.Fatal(app.Listen(":" + port))


}


