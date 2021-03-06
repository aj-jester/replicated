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

type Body7 struct {
	// Optional default value for this licensefield.
	Default_ string `json:"default"`
	// Indicates if this field will be visible from the on-premise license screen.
	Hidden bool `json:"hidden"`
	// Title of custom license field to display.
	Title string `json:"title"`
}
