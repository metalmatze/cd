// Code generated by go-swagger; DO NOT EDIT.

package pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/signalcd/signalcd/api/v1/models"
)

// CreatePipelineReader is a Reader for the CreatePipeline structure.
type CreatePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePipelineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePipelineOK creates a CreatePipelineOK with default headers values
func NewCreatePipelineOK() *CreatePipelineOK {
	return &CreatePipelineOK{}
}

/*CreatePipelineOK handles this case with default header values.

OK
*/
type CreatePipelineOK struct {
	Payload *models.Pipeline
}

func (o *CreatePipelineOK) Error() string {
	return fmt.Sprintf("[POST /pipelines][%d] createPipelineOK  %+v", 200, o.Payload)
}

func (o *CreatePipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineBadRequest creates a CreatePipelineBadRequest with default headers values
func NewCreatePipelineBadRequest() *CreatePipelineBadRequest {
	return &CreatePipelineBadRequest{}
}

/*CreatePipelineBadRequest handles this case with default header values.

CreatePipelineBadRequest create pipeline bad request
*/
type CreatePipelineBadRequest struct {
}

func (o *CreatePipelineBadRequest) Error() string {
	return fmt.Sprintf("[POST /pipelines][%d] createPipelineBadRequest ", 400)
}

func (o *CreatePipelineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
