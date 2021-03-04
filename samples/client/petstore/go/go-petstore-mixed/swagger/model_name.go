/*
 * Swagger Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Model for testing model name same as property name
type Name struct {
	Name int32 `json:"name"`
	SnakeCase int32 `json:"snake_case,omitempty"`
	Property string `json:"property,omitempty"`
	Var123Number int32 `json:"123Number,omitempty"`
}