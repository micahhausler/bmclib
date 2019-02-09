/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// This resource contains information about a specific Job scheduled  or being executed by a Redfish service's Job Service.
type Job2 struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The person or program that created this job entry.
	CreatedBy string `json:"CreatedBy,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The date-time stamp that the job was completed.
	EndTime time.Time `json:"EndTime,omitempty"`
	// Indicates that the contents of the Payload should be hidden from view after the Job has been created.  When set to True, the Payload object will not be returned on GET.
	HidePayload bool `json:"HidePayload,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	JobState JobState `json:"JobState,omitempty"`
	JobStatus Health `json:"JobStatus,omitempty"`
	// The maximum amount of time the job is allowed to execute.
	MaxExecutionTime string `json:"MaxExecutionTime,omitempty"`
	// This is an array of messages associated with the job.
	Messages []IdRef `json:"Messages,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	Payload Payload2 `json:"Payload,omitempty"`
	// The completion percentage of this job.
	PercentComplete int32 `json:"PercentComplete,omitempty"`
	Schedule IdRef `json:"Schedule,omitempty"`
	// The date-time stamp that the job was started or is scheduled to start.
	StartTime time.Time `json:"StartTime,omitempty"`
	// This represents the serialized execution order of the Job Steps.
	StepOrder []string `json:"StepOrder,omitempty"`
	Steps IdRef `json:"Steps,omitempty"`
}
