package serialization

// MessageJson Represents a message for sending to the Injection Api.
// To be serialized into JSON string before sending to the Injection Api.
type MessageJson struct {

	// List of To recipients.
	To []AddressJson `json:"To"`

	// List of Cc recipients.
	Cc []AddressJson `json:"Cc,omitempty"`

	// List of Bcc recipients.
	Bcc []AddressJson `json:"Bcc,omitempty"`

	// From Address
	From AddressJson `json:"From"`

	// Message Subject
	Subject string `json:"Subject"`

	// Container for both global and per recipient merge data (only applicable to bulk style messages)
	MergeData *MergeDataJson `json:"MergeData,omitempty"`

	// Plain text portion of the message body.
	TextBody string `json:"TextBody,omitempty"`

	// HTML portion of the message body.
	HTMLBody string `json:"HtmlBody,omitempty"`

	// Api Template for the message.
	APITemplate string `json:"ApiTemplate,omitempty"`

	// Custom MailingId for the message.
	MailingID string `json:"MailingId,omitempty"`

	// Custom MessageId for the message.
	MessageID string `json:"MessageId,omitempty"`

	// Optional character set for your message.
	CharSet string `json:"CharSet,omitempty"`

	// Reply To address.
	ReplyTo *AddressJson `json:"ReplyTo,omitempty"`

	// List of attachments.
	Attachments []AttachmentJson `json:"Attachments,omitempty"`

	// List of custom message headers added to the message.
	CustomHeaders []CustomHeadersJson `json:"CustomHeaders,omitempty"`
}
