package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg MongoInstance

const (
	dbName   = "fiber-hrms"
	mongoURI = "mongodb://localhost:27017/" + dbName
)

// MongoInstance contains the instance of the mongo client and db.
type MongoInstance struct {
	Client *mongo.Client   // A reference to mongo.Client
	Db     *mongo.Database // A reference to mongo.Database
}

// Employee contains the basic information of an employee.
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// Connect is used to connect to your Mongo DB.
func Connect() error {
	// Create a new client to connect to the mongo URI
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	// In case your MongoDB is not responding or something goes wrong or insert
	// is taking too long you have a timeout, you don't want to block the entire program.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Initializes the client by starting background monitoring goroutines.
	err = client.Connect(ctx)

	// Returns a handle for a database
	db := client.Database(dbName)

	// Handling error
	if err != nil {
		return err
	}

	// Setting our mongo instance
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {
	// Connecting to Mongo DB
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Defining your application by creating a new fiber instance.
	app := fiber.New()

	// Defining your routes with method type.
	// Getting all the employees from our database.
	app.Get("/employee", func(ctx *fiber.Ctx) error {
		// Slice of employee, like an array. It contains the actual objects. We're
		// initializing it using the `make` command. `make`  creates an empty slice.
		var employees []Employee = make([]Employee, 0)

		// Defining query. Because we want to get all the employees, we just leave it an empty {}.
		query := bson.D{{}}

		// Getting all the data under the employees table. `Collection` is the table. Find is used to
		// get or execute our query to get all the data.
		cursor, err := mg.Db.Collection("employees").Find(ctx.Context(), query)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// All the result from the `Find` function it will pass the values to the employees slice.
		if err := cursor.All(ctx.Context(), &employees); err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Converts into JSON format and sending response
		return ctx.JSON(employees)
	})

	// Creating or inserting new employee information
	app.Post("/employee", func(ctx *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		// `new` initialize a struct and returns a pointer to an instance of Employee struct
		employee := new(Employee)

		// Getting data from the actual request. It parses the body of employee data that the user is sending.
		if err := ctx.BodyParser(employee); err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}

		employee.ID = ""

		// Inserts a single row of data
		insertResult, err := collection.InsertOne(ctx.Context(), employee)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		// Building a query. The `InsertedID` is the ID of inserted data.
		filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}

		// Getting the created record
		createdRecord := collection.FindOne(ctx.Context(), filter)

		createdEmployee := &Employee{}
		// Unmarshaling data to the createdEmployee
		createdRecord.Decode(createdEmployee)

		return ctx.Status(http.StatusOK).JSON(createdEmployee)
	})

	// Updates the information of the specific employee
	app.Put("/employee/:id", func(ctx *fiber.Ctx) error {
		// Getting the route parameter id
		id := ctx.Params("id")

		// Creates a new ID from a hex string
		employeeID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}

		// `new` initialize a struct and returns a pointer to an instance of Employee struct
		employee := new(Employee)

		// Parses the body of employee data that the user is sending. User might use
		// Postman/Curl and it is sending you a new data of Employee. So we need to check
		// first the data by parsing the body. So if there's an issue parsing the data,
		// it will send an error.
		if err := ctx.BodyParser(employee); err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}

		// Find the exact record in order to update the specific employee by building the query.
		query := bson.D{{Key: "_id", Value: employeeID}}
		// Build the update query.
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}

		// Updates the information of specific employee
		err = mg.Db.Collection("employees").FindOneAndUpdate(ctx.Context(), query, update).Err()
		if err != nil {
			// If there's no matched documents or record
			if err == mongo.ErrNoDocuments {
				return ctx.SendStatus(http.StatusNotFound)
			}

			return ctx.SendStatus(http.StatusInternalServerError)
		}

		employee.ID = id
		return ctx.Status(http.StatusOK).JSON(employee)
	})

	// Deletes the specific employee
	app.Delete("/employee/:id", func(ctx *fiber.Ctx) error {
		// Creates a new ID from a hex string and getting the route param
		employeeID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
		if err != nil {
			return ctx.SendStatus(http.StatusBadRequest)
		}

		// Build the query
		query := bson.D{{Key: "_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(ctx.Context(), query)
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		// No record found and deleted
		if result.DeletedCount < 0 {
			return ctx.SendStatus(http.StatusNotFound)
		}

		return ctx.Status(http.StatusOK).JSON("record deleted")
	})

	// Starts the server. If it does not server, it will log the error.
	log.Fatal(app.Listen(":3000"))
}
