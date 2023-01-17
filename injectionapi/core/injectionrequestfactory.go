package core

import (
	"github.com/socketlabs/socketlabs-go/injectionapi/core/serialization"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

// InjectionRequestFactory is used to convert user-friendly message types to a structure that is ready for json-serialization
// when making injection api requests
type InjectionRequestFactory struct {
	ServerID int
	APIKey   string
}

// GenerateBasicRequest converts a basic message to an injection request in a json-serialization ready format
func (factory InjectionRequestFactory) GenerateBasicRequest(message *message.BasicMessage) serialization.InjectionRequestJson {

	r := serialization.InjectionRequestJson{
		ServerID: factory.ServerID,
		APIKey:   factory.APIKey,
	}

	messageJson := serialization.MessageJson{
		Subject:     message.Subject,
		TextBody:    message.PlainTextBody,
		HTMLBody:    message.HtmlBody,
		AMPBody:     message.AmpBody,
		MailingID:   message.MailingId,
		MessageID:   message.MessageId,
		CharSet:     message.CharSet,
		APITemplate: message.ApiTemplate,
	}

	//set attachments
	messageJson.Attachments = getAttachments(&message.Attachments)

	//set custom headers
	messageJson.CustomHeaders = getCustomHeaders(&message.CustomHeaders)

	//set from address
	messageJson.From = serialization.AddressJson{
		EmailAddress: message.From.EmailAddress,
		FriendlyName: message.From.FriendlyName,
	}

	//reply-to is optional, if provided, set it
	if message.ReplyTo.IsValid() {
		messageJson.ReplyTo = &serialization.AddressJson{
			EmailAddress: message.ReplyTo.EmailAddress,
			FriendlyName: message.ReplyTo.FriendlyName,
		}
	}

	messageJson.To = getAddressList(&message.To)
	messageJson.Cc = getAddressList(&message.Cc)
	messageJson.Bcc = getAddressList(&message.Bcc)

	//add the message
	r.Messages = append(r.Messages, messageJson)

	return r
}

// GenerateBulkRequest converts a bulk message to an injection request in a json-serialization ready format
func (factory InjectionRequestFactory) GenerateBulkRequest(message *message.BulkMessage) serialization.InjectionRequestJson {

	r := serialization.InjectionRequestJson{
		ServerID: factory.ServerID,
		APIKey:   factory.APIKey,
	}

	messageJson := serialization.MessageJson{
		Subject:     message.Subject,
		TextBody:    message.PlainTextBody,
		HTMLBody:    message.HtmlBody,
		AMPBody:     message.AmpBody,
		MailingID:   message.MailingId,
		MessageID:   message.MessageId,
		CharSet:     message.CharSet,
		APITemplate: message.ApiTemplate,
		Tags:		 message.Tags,
	}

	//set attachments
	messageJson.Attachments = getAttachments(&message.Attachments)

	//set custom headers
	messageJson.CustomHeaders = getCustomHeaders(&message.CustomHeaders)

	//set meta data
	messageJson.Metadata = getMetadata(&message.Metadata)

	//set from address
	messageJson.From = serialization.AddressJson{
		EmailAddress: message.From.EmailAddress,
		FriendlyName: message.From.FriendlyName,
	}

	//reply-to is optional, if provided, set it
	if message.ReplyTo.IsValid() {
		messageJson.ReplyTo = &serialization.AddressJson{
			EmailAddress: message.ReplyTo.EmailAddress,
			FriendlyName: message.ReplyTo.FriendlyName,
		}
	}

	placeholder := serialization.AddressJson{
		EmailAddress: "%%DeliveryAddress%%",
		FriendlyName: "%%RecipientName%%",
	}

	messageJson.To = append(messageJson.To, placeholder)
	messageJson.MergeData = &serialization.MergeDataJson{
		Global:     getGlobalMergeFields(&message.Global),
		PerMessage: getBulkMergeFields(&message.To),
	}

	//add the message
	r.Messages = append(r.Messages, messageJson)

	return r
}

//getGlobalMergeFields converts string map to merge fields ready for json-serialization
func getGlobalMergeFields(global *map[string]string) []serialization.MergeFieldJson {
	var globalMergeData = *global
	if globalMergeData == nil || len(globalMergeData) == 0 {
		return nil
	}

	mergeFields := []serialization.MergeFieldJson{}

	//iterate through merge data and add to results
	for field, value := range globalMergeData {
		mergeFields = append(mergeFields, serialization.MergeFieldJson{
			Field: field,
			Value: value,
		})
	}

	return mergeFields
}

// getAddressList converts an address slice to json-serialization ready format
func getAddressList(emailAddresses *[]message.EmailAddress) []serialization.AddressJson {

	if len(*emailAddresses) == 0 {
		return nil
	}

	results := []serialization.AddressJson{}

	for _, sourceAddress := range *emailAddresses {
		results = append(results, serialization.AddressJson{
			EmailAddress: sourceAddress.EmailAddress,
			FriendlyName: sourceAddress.FriendlyName,
		})
	}

	return results
}

// getAttachments converts attachments to json-serialization ready format
func getAttachments(attachments *[]message.Attachment) []serialization.AttachmentJson {

	if len(*attachments) == 0 {
		return nil
	}

	results := []serialization.AttachmentJson{}

	for _, sourceAttachment := range *attachments {
		attachment := serialization.AttachmentJson{
			Name:        sourceAttachment.Name,
			ContentType: sourceAttachment.MimeType,
			ContentID:   sourceAttachment.ContentID,
			Content:     &sourceAttachment.Content,
		}
		attachment.CustomHeaders = getCustomHeaders(&sourceAttachment.CustomHeaders)
		results = append(results, attachment)
	}
	return results
}

// getCustomHeaders converts custom headers to json-serialization ready format
func getCustomHeaders(customHeaders *[]message.CustomHeader) []serialization.CustomHeadersJson {

	if len(*customHeaders) == 0 {
		return nil
	}

	var results = []serialization.CustomHeadersJson{}

	for _, sourceHeader := range *customHeaders {
		results = append(results, serialization.CustomHeadersJson{
			Name:  sourceHeader.Name,
			Value: sourceHeader.Value,
		})
	}

	return results
}

// getMetadata converts meta data to json-serialization ready format
func getMetadata(metadata *[]message.Metadata) []serialization.MetadataJson {

	if len(*metadata) == 0 {
		return nil
	}

	var results = []serialization.MetadataJson{}

	for _, sourceData := range *metadata {
		results = append(results, serialization.MetadataJson{
			Key:  sourceData.Key,
			Value: sourceData.Value,
		})
	}

	return results
}

// getBulkMergeFields extracts merge fields from slice of bulk recipients in a json-serialization ready format
func getBulkMergeFields(bulkRecipients *[]message.BulkRecipient) [][]serialization.MergeFieldJson {

	if len(*bulkRecipients) == 0 {
		return nil
	}

	mergeFields := [][]serialization.MergeFieldJson{}

	for _, recipient := range *bulkRecipients {
		mergeFields = append(mergeFields, getRecipientMergeFields(&recipient))
	}

	return mergeFields
}

// getRecipientMergeFields extracts merge fields from a bulk recipient in a json-serialization ready format
func getRecipientMergeFields(bulkRecipient *message.BulkRecipient) []serialization.MergeFieldJson {

	recipientMergeFields := []serialization.MergeFieldJson{}

	//add address
	recipientMergeFields = append(recipientMergeFields, serialization.MergeFieldJson{
		Field: "DeliveryAddress",
		Value: bulkRecipient.Email,
	})

	//add friendly name if provided
	recipientMergeFields = append(recipientMergeFields, serialization.MergeFieldJson{
		Field: "RecipientName",
		Value: bulkRecipient.FriendlyName,
	})

	//iterate through merge data and add to results
	if len(bulkRecipient.MergeData) > 0 {
		for field, value := range bulkRecipient.MergeData {
			recipientMergeFields = append(recipientMergeFields, serialization.MergeFieldJson{
				Field: field,
				Value: value,
			})
		}
	}

	return recipientMergeFields
}
