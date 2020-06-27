/*
 * SignalCD
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DeploymentApiController binds http requests to an api service and writes the service results to the http response
type DeploymentApiController struct {
	service DeploymentApiServicer
}

// NewDeploymentApiController creates a default api controller
func NewDeploymentApiController(s DeploymentApiServicer) Router {
	return &DeploymentApiController{service: s}
}

// Routes returns all of the api route for the DeploymentApiController
func (c *DeploymentApiController) Routes() Routes {
	return Routes{
		{
			"GetCurrentDeployment",
			strings.ToUpper("Get"),
			"/api/v1/deployments/current",
			c.GetCurrentDeployment,
		},
		{
			"ListDeployments",
			strings.ToUpper("Get"),
			"/api/v1/deployments",
			c.ListDeployments,
		},
		{
			"SetCurrentDeployment",
			strings.ToUpper("Post"),
			"/api/v1/deployments/current",
			c.SetCurrentDeployment,
		},
		{
			"UpdateDeploymentStatus",
			strings.ToUpper("Patch"),
			"/api/v1/deployments/{id}/status",
			c.UpdateDeploymentStatus,
		},
	}
}

// GetCurrentDeployment - Get the current Deployment
func (c *DeploymentApiController) GetCurrentDeployment(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetCurrentDeployment()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// ListDeployments - List Deployments
func (c *DeploymentApiController) ListDeployments(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListDeployments()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// SetCurrentDeployment - Set the current Deployment
func (c *DeploymentApiController) SetCurrentDeployment(w http.ResponseWriter, r *http.Request) {
	setCurrentDeployment := &SetCurrentDeployment{}
	if err := json.NewDecoder(r.Body).Decode(&setCurrentDeployment); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.SetCurrentDeployment(*setCurrentDeployment)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// UpdateDeploymentStatus - Update parts of the Status of a Deployment
func (c *DeploymentApiController) UpdateDeploymentStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := parseIntParameter(params["id"])
	if err != nil {
		w.WriteHeader(500)
		return
	}

	deploymentStatusUpdate := &DeploymentStatusUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&deploymentStatusUpdate); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.UpdateDeploymentStatus(id, *deploymentStatusUpdate)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
