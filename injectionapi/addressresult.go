package injectionapi

//AddressResult The result of a single recipient in the Injection request.
type AddressResult struct {

	// The recipient's email address.
	EmailAddress string

	// Whether the recipient was accepted for delivery.
	Accepted bool

	// An error code detailing why the recipient was not accepted.
	ErrorCode string
}
