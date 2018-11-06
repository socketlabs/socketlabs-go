package serialization

// MergeDataJson represents MergeData for a single message.
// To be serialized into JSON string before sending to the Injection Api.
type MergeDataJson struct {

	// Defines merge field data for all messages in the request
	Global []MergeFieldJson `json:"Global,omitempty"`

	// Defines merge field data for each message
	PerMessage [][]MergeFieldJson `json:"PerMessage,omitempty"`
}
