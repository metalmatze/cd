// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/signalcd/signalcd/api/v1/models"
)

// DeploymentsOKCode is the HTTP code returned for type DeploymentsOK
const DeploymentsOKCode int = 200

/*DeploymentsOK OK

swagger:response deploymentsOK
*/
type DeploymentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Deployment `json:"body,omitempty"`
}

// NewDeploymentsOK creates DeploymentsOK with default headers values
func NewDeploymentsOK() *DeploymentsOK {

	return &DeploymentsOK{}
}

// WithPayload adds the payload to the deployments o k response
func (o *DeploymentsOK) WithPayload(payload []*models.Deployment) *DeploymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deployments o k response
func (o *DeploymentsOK) SetPayload(payload []*models.Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeploymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Deployment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeploymentsInternalServerErrorCode is the HTTP code returned for type DeploymentsInternalServerError
const DeploymentsInternalServerErrorCode int = 500

/*DeploymentsInternalServerError internal server error

swagger:response deploymentsInternalServerError
*/
type DeploymentsInternalServerError struct {
}

// NewDeploymentsInternalServerError creates DeploymentsInternalServerError with default headers values
func NewDeploymentsInternalServerError() *DeploymentsInternalServerError {

	return &DeploymentsInternalServerError{}
}

// WriteResponse to the client
func (o *DeploymentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
