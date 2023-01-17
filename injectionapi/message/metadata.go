package message

// CustomHeader  info
type Metadata struct {
	Key  string
	Value string
}

// NewMetadata  Factory Constructor
func NewMetadata(key string, value string) Metadata {
	var c Metadata
	c.Key = key
	c.Value = value
	return c
}

// IsValid  Determines if the Metadata is valid
func (c Metadata) IsValid() bool {

	if len(c.Key) == 0 {
		return false
	}

	if len(c.Value) == 0 {
		return false
	}

	return true
}