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
	"time"
)

type Deployment struct {
	Number int64 `json:"number"`

	Created time.Time `json:"created,omitempty"`
}
