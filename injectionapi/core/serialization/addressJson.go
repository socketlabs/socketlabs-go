package serialization

// An AddressJson represents an individual email address for a message.
// To be serialized into JSON string before sending to the Injection Api.
type AddressJson struct {

	// EmaillAddress is a valid email address
	EmailAddress string `json:"EmailAddress"`

	//The friendly or display name for the recipient.
	FriendlyName string `json:"FriendlyName"`
}
