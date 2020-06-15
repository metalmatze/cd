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
	"errors"
)

// DeploymentApiService is a service that implents the logic for the DeploymentApiServicer
// This service should implement the business logic for every endpoint for the DeploymentApi API.
// Include any external packages or services that will be required by this service.
type DeploymentApiService struct {
}

// NewDeploymentApiService creates a default api service
func NewDeploymentApiService() DeploymentApiServicer {
	return &DeploymentApiService{}
}

// GetCurrentDeployment - Get the current Deployment
func (s *DeploymentApiService) GetCurrentDeployment() (interface{}, error) {
	// TODO - update GetCurrentDeployment with the required logic for this service method.
	// Add api_deployment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetCurrentDeployment' not implemented")
}

// ListDeployments - List Deployments
func (s *DeploymentApiService) ListDeployments() (interface{}, error) {
	// TODO - update ListDeployments with the required logic for this service method.
	// Add api_deployment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ListDeployments' not implemented")
}

// SetCurrentDeployment - Set the current Deployment
func (s *DeploymentApiService) SetCurrentDeployment(inlineObject InlineObject) (interface{}, error) {
	// TODO - update SetCurrentDeployment with the required logic for this service method.
	// Add api_deployment_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'SetCurrentDeployment' not implemented")
}
