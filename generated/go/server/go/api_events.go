/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A EventsApiController binds http requests to an api service and writes the service results to the http response
type EventsApiController struct {
	service EventsApiServicer
}

// NewEventsApiController creates a default api controller
func NewEventsApiController(s EventsApiServicer) Router {
	return &EventsApiController{ service: s }
}

// Routes returns all of the api route for the EventsApiController
func (c *EventsApiController) Routes() Routes {
	return Routes{ 
		{
			"EventsPost",
			strings.ToUpper("Post"),
			"/v2/events",
			c.EventsPost,
		},
	}
}

// EventsPost - 
func (c *EventsApiController) EventsPost(w http.ResponseWriter, r *http.Request) { 
	xEffxApiKey := r.Header.Get("X-Effx-Api-Key")
	createEventPayload := &CreateEventPayload{}
	if err := json.NewDecoder(r.Body).Decode(&createEventPayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.EventsPost(r.Context(), xEffxApiKey, *createEventPayload)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
