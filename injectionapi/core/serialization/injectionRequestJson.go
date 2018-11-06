package serialization

// InjectionRequestJson Represents a injection request for sending to the Injection Api.
type InjectionRequestJson struct {

	// Your SocketLabs ServerId number.
	ServerID int `json:"serverId"`

	// Your SocketLabs Injection API key.
	APIKey string `json:"APIKey"`

	// Slice of messages to be sent.
	Messages []MessageJson `json:"Messages"`
}
