/*
 * Vendor API V1
 *
 * Apps documentation
 *
 * API version: 1.0.0
 * Contact: info@replicated.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// An app belongs to a team. It contains channels onto which releases can be promoted.
type App struct {
	// The ID of the app
	Id string `json:"Id"`
	// The name of the app
	Name string `json:"Name"`
	// The scheduler associated with the app
	Scheduler string `json:"Scheduler,omitempty"`
	// A unique slug for the app
	Slug string `json:"Slug"`
}