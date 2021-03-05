/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ContactInformation struct {

	Email string `json:"email,omitempty"`

	OnCall Link `json:"onCall,omitempty"`

	Chat Link `json:"chat,omitempty"`

	IssueTracker Link `json:"issueTracker,omitempty"`
}
