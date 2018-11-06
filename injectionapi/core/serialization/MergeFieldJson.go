package serialization

// MergeFieldJson represents a merge field as a field and value pair.
// To be serialized into JSON string before sending to the Injection Api.
type MergeFieldJson struct {

	// Gets or sets the merge field
	Field string `json:"Field"`

	// Gets or sets the merge field value
	Value string `json:"Value"`
}
