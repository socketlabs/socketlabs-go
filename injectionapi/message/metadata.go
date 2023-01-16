package message

// CustomHeader  info
type Metadata struct {
	Name  string
	Value string
}

// NewMetadata  Factory Constructor
func NewMetadata(name string, value string) Metadata {
	var c Metadata
	c.Name = name
	c.Value = value
	return c
}

// IsValid  Determines if the Metadata is valid
func (c Metadata) IsValid() bool {

	if len(c.Name) == 0 {
		return false
	}

	if len(c.Value) == 0 {
		return false
	}

	return true
}