package serialization

// MetadataJson represents a piece of meta data as a key and value pair.
// To be serialized into JSON string before sending to the Injection Api.
type MetadataJson struct {

	// Gets or sets the meta data key
	Key string `json:"Key"`

	// Gets or sets the meta data value
	Value string `json:"Value"`
}