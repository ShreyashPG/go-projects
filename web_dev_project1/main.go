package main

import (
	// "fmt"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

type Employee struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Position string `json:"position"`

}

var employees =[]Employee{
	{ID : 1, Name: "John Doe", Email: "jhn@gmial.com" , Age: 25, Position: "Software Engineer"},
	{ID : 2, Name: "Jane Doe", Email: "jane@gmail.com" , Age: 30, Position: "Senior Software Engineer"},
	{ID: 3, Name: "Michael Doe", Email: "michel@gmail.com", Age: 35, Position: "Project Manager"},
	{ID : 4, Name: "Jonny Doe", Email: "doe@gmail.com" , Age: 40, Position: "Senior Project Manager"},
	{ID : 5, Name: "Jenny Doe", Email: "jenney@gmail.com" , Age: 45, Position: "CEO"},
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/employees", func(c *fiber.Ctx) error {
		return c.JSON(employees)
	})

	app.Post("/add-employee", func(c *fiber.Ctx) error {
		employee:= new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee.ID = len(employees)+1
		
		employees=append(employees,*employee)
		return  c.Status(201).JSON(fiber.Map{
			"message": "Employee added successfully",
			"employee": employee,
		})
	})

	app.Put("/update-employee/:id", func(c *fiber.Ctx) error{
		idStr :=c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}


		employee:= new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		for i, emp := range employees{
			if emp.ID == id {
				employees[i] = *employee
				return c.JSON(fiber.Map{
					"message": "Employee updated successfully",
					"employee": employee,
				})
			}
		}
		return c.Status(404).SendString("Employee not found")
	})

	app.Delete("/delete-employee/:id", func(c *fiber.Ctx) error{
		idStr:= c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).SendString(err.Error())	}

		for i, emp := range employees{
			if emp.ID == id {
				employees = append(employees[:i], employees[i+1:]...)
				return c.JSON(fiber.Map{
					"message": "Employee deleted successfully",
				})
			}
		}
		return c.Status(404).SendString("Employee not found")
	})

	app.Get("/", func(c *fiber.Ctx) error{
		return c.SendString("Hello, World!")
	})
	app.Post("/hello-world", func(c *fiber.Ctx) error {
		greet := c.FormValue("greet")
			// if greet == "" {
			// 	greet = "Hello, World!"
			// }

		return c.JSON(fiber.Map{
			"message": greet,
		})
	})

	log.Fatal(app.Listen(":3001"))

}
