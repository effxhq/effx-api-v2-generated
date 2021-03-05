/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Team struct {

	// uuid of the team object
	Id string `json:"id,omitempty"`

	// url safe string representation of the team's name
	Slug string `json:"slug,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Contact ContactInformation `json:"contact,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`

	// Attach metadata to resources like teams and services
	Annotations map[string]string `json:"annotations,omitempty"`
}
