package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/rmarasigan/go-fiber-crm-basic/database"
	"github.com/rmarasigan/go-fiber-crm-basic/lead"
)

// Accepts fiber application by calling it by reference
func setupRoutes(app *fiber.App) {
	// `app` is an instance of your fiber. first param: it is where it is routed
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)

}

// initDatabase initializes our database
func initDatabase() {
	var err error

	// We use gorm to open a connection your sqlite3 database.
	// The database name is `leads.db`
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	// Check if there's an error
	if err != nil {
		panic("failed to connect to database")
	}

	// If everything went well
	fmt.Println("Connection opened to database")

	// Migrating your model, Lead. Auto migration for given models,
	// will only add missing fields, won't delete/change current data.
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	// Creates a new Fiber isntance
	app := fiber.New()
	// Start the connection
	initDatabase()
	setupRoutes(app)

	// Listen to port 3000 just like ListenAndServe
	app.Listen(3000)

	// Closing the connection. It will happen at the end when the function has completely run.
	defer database.DBConn.Close()
}
