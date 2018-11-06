package serialization

// CustomHeadersJson represents a custom header as a name and value pair.
// To be serialized into JSON string before sending to the Injection Api.
type CustomHeadersJson struct {

	// Gets or sets the custom header name
	Name string `json:"Name"`

	// Gets or sets the custom header value
	Value string `json:"Value"`
}
