package lead

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/rmarasigan/go-fiber-crm-basic/database"
)

// Lead contains the basic information of our lead.
type Lead struct {
	gorm.Model
	Name    string `json:"name"` // Wer'e telling golang that this is what our JSON should look like
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

// GetLeads returns all the leads that is existing.
// Context pass request-scoped values, cancelation signals
// and deadlines across API boundaries.
func GetLeads(ctx *fiber.Ctx) {
	// Database connection
	db := database.DBConn

	// Slice of leads, that means it could be multiple leads which is like an array
	var leads []Lead

	// Get all the leads
	db.Find(&leads)

	// Converts into JSON format and sending response
	ctx.JSON(leads)
}

// GetLead returns a single row data based on the `id` parameter you set
func GetLead(ctx *fiber.Ctx) {
	// `Lead` is the struct that you defined, so `lead` is the name of the variable.
	// `Lead` is the data type
	var lead Lead

	// Getting the route parameter
	id := ctx.Params("id")

	// Database connection
	db := database.DBConn
	// Find the Lead using id parameter supplied
	db.Find(&lead, id)

	// Converts into JSON format and sending response
	ctx.JSON(lead)
}

// NewLead creates or insert a new data with the values supplied.
func NewLead(ctx *fiber.Ctx) {
	db := database.DBConn

	// `new` initialize a struct and returns a pointer to an instance of Lead struct
	lead := new(Lead)

	// Parses the body of lead data that the user is sending. User might use
	// Postman/Curl and it is sending you a new data of Lead. So we need to check
	// first the data by parsing the body. So if there's an issue parsing the data,
	// it will send an error.
	if err := ctx.BodyParser(lead); err != nil {
		// Status Code 503 = StatusServiceUnavailable
		ctx.Status(503).Send(err)
		return
	}

	// Create or inserting new data to our database
	db.Create(&lead)
	ctx.JSON(lead)
}

// DeleteLead deletes the specific lead based on the `id` paramater supplied.
func DeleteLead(ctx *fiber.Ctx) {
	var lead Lead
	id := ctx.Params("id")
	db := database.DBConn

	// Gets the first record that matches the given condition
	db.First(&lead, id)

	if lead.Name == "" {
		// Status Code 500 = StatusInternalServerError
		ctx.Status(http.StatusInternalServerError).Send("No lead found with ID")
		return
	}

	// Delete the record
	db.Delete(&lead)
	ctx.Send("Lead successfully deleted")
}
