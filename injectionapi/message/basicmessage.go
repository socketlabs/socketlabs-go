package message

// BasicMessage is a message similar to one created in a personal email client such as Outlook.
// This message can have many recipients of different types, such as To, CC, and BCC.  This
// message does not support merge fields.
type BasicMessage struct {

	//The message subject.
	Subject string

	// Plain text portion of the message body.
	// (Optional) Either TextBody or HtmlBody must be used with the AmpBody or use a ApiTemplate
	PlainTextBody string

	// HTML portion of the message body.
	// (Optional) Either TextBody or HtmlBody must be used with the AmpBody or use a ApiTemplate
	HtmlBody string

	// AMP portion of the message body.
	// (Optional) Either TextBody or HtmlBody must be used with the AmpBody or use a ApiTemplate
	AmpBody string

	// Api Template for the message body.
	// (Optional) Either TextBody or HtmlBody must be used with the AmpBody or use a ApiTemplate
	ApiTemplate string

	//Custom MailingId for the message.
	//See https://www.socketlabs.com/blog/best-practices-for-using-custom-mailingids-and-messageids/ for more information.
	//(Optional)
	MailingId string

	//Custom MessageId for the message.
	//(Optional)
	MessageId string

	//From Address for the message.
	From EmailAddress

	//ReplyTo Address for the message.
	//(Optional)
	ReplyTo EmailAddress

	//The optional character set for the message.
	//(Optional)
	CharSet string

	//Optional collection of message attachments.
	//(Optional)
	Attachments []Attachment

	//Optional collection of custom headers for the message.
	//(Optional)
	CustomHeaders []CustomHeader

	//Optional collection of meta data for the message.
	//(Optional)
	Metadatas []Metadata

	//Optional collection of tags for the message.
	//(Optional)
	Tags []string

	//Collection of To Recipients for the message.
	To []EmailAddress

	//Collection of CC Recipients for the message.
	Cc []EmailAddress

	//Collection of BCC Recipients for the message.
	Bcc []EmailAddress
}

// AddToEmailAddress adds an email address to the To Recipients collection
func (basic *BasicMessage) AddToEmailAddress(email string) {
	to := NewEmailAddress(email)
	basic.To = append(basic.To, to)
}

// AddToFriendlyEmailAddress adds an email address paired with a friendlyname to the To Recipients collection
func (basic *BasicMessage) AddToFriendlyEmailAddress(email string, friendly string) {
	to := NewFriendlyEmailAddress(email, friendly)
	basic.To = append(basic.To, to)
}

// AddCcEmailAddress adds an email address to the Cc Recipients collection
func (basic *BasicMessage) AddCcEmailAddress(email string) {
	cc := NewEmailAddress(email)

	basic.Cc = append(basic.Cc, cc)
}

// AddCcFriendlyEmailAddress adds an email address paired with a friendlyname to the Cc Recipients collection
func (basic *BasicMessage) AddCcFriendlyEmailAddress(email string, friendly string) {
	cc := NewFriendlyEmailAddress(email, friendly)
	basic.Cc = append(basic.Cc, cc)
}

// AddBccEmailAddress adds an email address to the Bcc Recipients collection
func (basic *BasicMessage) AddBccEmailAddress(email string) {
	bcc := NewEmailAddress(email)
	basic.Bcc = append(basic.Bcc, bcc)
}

// AddBccFriendlyEmailAddress adds an email address paired with a friendlyname to the Bcc Recipients collection
func (basic *BasicMessage) AddBccFriendlyEmailAddress(email string, friendly string) {
	bcc := NewFriendlyEmailAddress(email, friendly)
	basic.Bcc = append(basic.Bcc, bcc)
}

// AddCustomHeader adds a custom header to the message
func (basic *BasicMessage) AddCustomHeader(name string, value string) {
	customHeader := NewCustomHeader(name, value)
	basic.CustomHeaders = append(basic.CustomHeaders, customHeader)
}

// AddMetadata adds meta data to the message
func (basic *BasicMessage) AddMetadata(name string, value string) {
	metadata := NewMetadata(name, value)
	basic.Metadatas = append(basic.Metadatas, metadata)
}

// AddTag adds a tag to the message
func (basic *BasicMessage) AddTag(value string) {
	basic.Tags = append(basic.Tags, value)
}
