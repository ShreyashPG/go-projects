package controllers

import (
	"github.com/ShreyashPG/MongoDB-web-dev-project/database"
	"github.com/ShreyashPG/MongoDB-web-dev-project/model"
	 "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

func CreateEmployee(c *fiber.Ctx) error {
	db := database.GetCollection("employees")

	var employee model.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate the employee data
	if err := employee.BeforeCreate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee data"})
	}

	// Check if employee already exists (optional)
	existingEmployee := model.Employee{}
	err := db.FindOne(context.Background(), bson.M{"email": employee.Email}).Decode(&existingEmployee)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Employee with this email already exists"})
	}

	// Insert employee into the database
	result, err := db.InsertOne(context.Background(), employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	// Optionally, return the created employee ID
	employee.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Employee created successfully",
		"data":    employee,
	})
}

func GetAllEmployees(c *fiber.Ctx) error {
	db := database.GetCollection("employees")

	cursor, err := db.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch employees"})
	}
	defer cursor.Close(context.Background()) // Close cursor after use

	var employees []model.Employee
	if err := cursor.All(context.Background(), &employees); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse employees"})
	}

	// If no employees found
	if len(employees) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No employees found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employees fetched successfully",
		"data":    employees,
	})
}

func GetEmployeeByID(c *fiber.Ctx) error {
	db := database.GetCollection("employees")

	// Get employee ID from the route parameter
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}

	var employee model.Employee
	err = db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&employee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch employee"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee fetched successfully",
		"data":    employee,
	})
}


func UpdateEmployee(c *fiber.Ctx) error{
	db:=database.GetCollection("employees")

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}

	//check if employee exists
	var existingEmployee model.Employee
	err = db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&existingEmployee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch employee"})
	}

	//parse and update employee
	var employee model.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := employee.BeforeUpdate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee data"})
	}
	
	//update employee in the databse
	_, err = db.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": employee})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update employee"})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee updated successfully",
		"data": employee,
	})
	
}

func DeleteEmployee(c *fiber.Ctx) error{
	db:=database.GetCollection("employees")

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}
	
	//check if employee exists
	var existingEmployee model.Employee
	err = db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&existingEmployee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch employee"})
	}

	//delete employee from the database
	_, err = db.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete employee"})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee deleted successfully",
	})
	
	
}
