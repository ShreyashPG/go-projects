//fiber is a expres like framework

package main

import (
    "fmt"
    "log"
    "os"

    "github.com/ShreyashPG/URLShortner/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"


)

func setupRoutes(app * fiber.App){
    app.Get("/:url", routes.ResolveURL)
    app.Post("api/v1", routes.ShortenURL)

}



func main(){
    err :=godotenv.Load()
    if err!= nil{
        fmt.Println(err);
    }

    app:= fiber.New()


//logger is used as log middleware in fiber 
//fiber is faster than gorilla mux 



    app.Use(logger.New());

        setupRoutes(app);

        log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}