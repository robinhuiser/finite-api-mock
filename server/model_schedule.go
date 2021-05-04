/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abstract API
 *
 * API version: v1.5.1
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

// Schedule - Schedule object
type Schedule struct {

	// The execution strategy of the schedule
	Strategy string `json:"strategy,omitempty"`

	// Indicates how frequencly this schedule should execute
	Frequency string `json:"frequency,omitempty"`

	// The day of the month on which a execution should take
	DayOn string `json:"dayOn,omitempty"`

	// ISO 6801 formatted date for the schedule to start
	StartDateTime string `json:"startDateTime,omitempty"`

	// ISO 6801 formatted date for the schedule to complete
	EndDateTime string `json:"endDateTime,omitempty"`

	// ISO 6801 formatted date for the next scheduled execution
	NextDateTime string `json:"nextDateTime,omitempty"`

	// The number of times this schedule should execute
	RepeatCount int32 `json:"repeatCount,omitempty"`

	// Indicates if the execution is even or odd
	IsEveryTime bool `json:"isEveryTime,omitempty"`
}
