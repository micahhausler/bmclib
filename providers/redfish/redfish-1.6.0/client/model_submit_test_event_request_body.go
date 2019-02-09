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

// This action is used to generate a test event.
type SubmitTestEventRequestBody struct {
	// This is the ID of event to be added.
	EventId string `json:"EventId"`
	// This is the time stamp of event to be added.
	EventTimestamp time.Time `json:"EventTimestamp"`
	EventType EventType `json:"EventType"`
	// This is the human readable message of event to be added.
	Message string `json:"Message"`
	// This is the array of message arguments of the event to be added.
	MessageArgs []string `json:"MessageArgs"`
	// This is the message ID of event to be added.
	MessageId string `json:"MessageId"`
	// This is the string of the URL within the OriginOfCondition property of the event to be added.  It is not a reference object.
	OriginOfCondition string `json:"OriginOfCondition"`
	// This is the Severity of event to be added.
	Severity string `json:"Severity"`
}
