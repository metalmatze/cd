// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/metalmatze/cd/api/v1/restapi/operations/pipeline"
)

// NewCdAPI creates a new Cd instance
func NewCdAPI(spec *loads.Document) *CdAPI {
	return &CdAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		PipelinePipelineHandler: pipeline.PipelineHandlerFunc(func(params pipeline.PipelineParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelinePipeline has not yet been implemented")
		}),
		PipelinePipelineAgentsHandler: pipeline.PipelineAgentsHandlerFunc(func(params pipeline.PipelineAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelinePipelineAgents has not yet been implemented")
		}),
		PipelinePipelineCurrentHandler: pipeline.PipelineCurrentHandlerFunc(func(params pipeline.PipelineCurrentParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelinePipelineCurrent has not yet been implemented")
		}),
		PipelinePipelinesHandler: pipeline.PipelinesHandlerFunc(func(params pipeline.PipelinesParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelinePipelines has not yet been implemented")
		}),
		PipelineUpdateCurrentPipelineHandler: pipeline.UpdateCurrentPipelineHandlerFunc(func(params pipeline.UpdateCurrentPipelineParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelineUpdateCurrentPipeline has not yet been implemented")
		}),
		PipelineUpdatePipelineAgentsHandler: pipeline.UpdatePipelineAgentsHandlerFunc(func(params pipeline.UpdatePipelineAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation PipelineUpdatePipelineAgents has not yet been implemented")
		}),
	}
}

/*CdAPI Swagger 2.0 specification for CD */
type CdAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PipelinePipelineHandler sets the operation handler for the pipeline operation
	PipelinePipelineHandler pipeline.PipelineHandler
	// PipelinePipelineAgentsHandler sets the operation handler for the pipeline agents operation
	PipelinePipelineAgentsHandler pipeline.PipelineAgentsHandler
	// PipelinePipelineCurrentHandler sets the operation handler for the pipeline current operation
	PipelinePipelineCurrentHandler pipeline.PipelineCurrentHandler
	// PipelinePipelinesHandler sets the operation handler for the pipelines operation
	PipelinePipelinesHandler pipeline.PipelinesHandler
	// PipelineUpdateCurrentPipelineHandler sets the operation handler for the update current pipeline operation
	PipelineUpdateCurrentPipelineHandler pipeline.UpdateCurrentPipelineHandler
	// PipelineUpdatePipelineAgentsHandler sets the operation handler for the update pipeline agents operation
	PipelineUpdatePipelineAgentsHandler pipeline.UpdatePipelineAgentsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CdAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CdAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CdAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CdAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CdAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CdAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CdAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CdAPI
func (o *CdAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PipelinePipelineHandler == nil {
		unregistered = append(unregistered, "pipeline.PipelineHandler")
	}

	if o.PipelinePipelineAgentsHandler == nil {
		unregistered = append(unregistered, "pipeline.PipelineAgentsHandler")
	}

	if o.PipelinePipelineCurrentHandler == nil {
		unregistered = append(unregistered, "pipeline.PipelineCurrentHandler")
	}

	if o.PipelinePipelinesHandler == nil {
		unregistered = append(unregistered, "pipeline.PipelinesHandler")
	}

	if o.PipelineUpdateCurrentPipelineHandler == nil {
		unregistered = append(unregistered, "pipeline.UpdateCurrentPipelineHandler")
	}

	if o.PipelineUpdatePipelineAgentsHandler == nil {
		unregistered = append(unregistered, "pipeline.UpdatePipelineAgentsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CdAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CdAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *CdAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CdAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CdAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CdAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the cd API
func (o *CdAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CdAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pipelines/{id}"] = pipeline.NewPipeline(o.context, o.PipelinePipelineHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pipelines/status"] = pipeline.NewPipelineAgents(o.context, o.PipelinePipelineAgentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pipeline"] = pipeline.NewPipelineCurrent(o.context, o.PipelinePipelineCurrentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pipelines"] = pipeline.NewPipelines(o.context, o.PipelinePipelinesHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/pipeline"] = pipeline.NewUpdateCurrentPipeline(o.context, o.PipelineUpdateCurrentPipelineHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pipelines/status"] = pipeline.NewUpdatePipelineAgents(o.context, o.PipelineUpdatePipelineAgentsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CdAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CdAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CdAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CdAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
