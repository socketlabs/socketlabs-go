package message

// CustomHeader  info
type CustomHeader struct {
	Name  string
	Value string
}

// NewCustomHeader  Factory Constructor
func NewCustomHeader(name string, value string) CustomHeader {
	var c CustomHeader
	c.Name = name
	c.Value = value
	return c
}

// IsValid  Determines if the CustomHeader is valid
func (c CustomHeader) IsValid() bool {

	if len(c.Name) == 0 {
		return false
	}

	if len(c.Value) == 0 {
		return false
	}

	return true
}
