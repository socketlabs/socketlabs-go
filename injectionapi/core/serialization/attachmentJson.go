package serialization

// AttachmentJson Represents a message attachment in the form of a byte array.
// To be serialized into JSON string before sending to the Injection Api.
type AttachmentJson struct {

	// Name of attachment (displayed in email clients)
	Name string `json:"Name"`

	// Attachment data
	Content *[]byte `json:"Content"`

	// When set, used to embed an image within the body of an email message.
	ContentID string `json:"ContentId"`

	// The ContentType (MIME type) of the attachment.
	ContentType string `json:"ContentType"`

	// A list of custom headers added to the attachment.
	CustomHeaders []CustomHeadersJson `json:"CustomHeaders"`
}
