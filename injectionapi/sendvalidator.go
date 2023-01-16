package injectionapi

import (
	"strconv"
	"strings"

	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

//SendValidator is used by the SocketLabsClient to conduct basic validation on the message before sending to the Injection API.
type sendValidator struct{}

//ValidateCredentials Validate the ServerId and Api Key pair prior before sending to the Injection API.
func (sendValidator) ValidateCredentials(serverID int, apiKey string) (response SendResponse) {
	response = SendResponse{}
	if len(strings.TrimSpace(apiKey)) == 0 || serverID < 0 {
		response.Result = SendResultAUTHENTICATIONVALIDATIONFAILED
	} else {
		response.Result = SendResultSUCCESS
	}
	return
}

//ValidateBasicMessage Validate a basic email message before sending to the Injection API.
func (sendValidator) ValidateBasicMessage(basicMessage message.BasicMessage) (response SendResponse) {
	response = SendResponse{Result: SendResultUNKNOWNERROR}

	// ReplyTo is an optional property, so check to see if it is in use first
	if len(basicMessage.ReplyTo.EmailAddress) > 0 { // ReplyTo is an optional property, so check to see if it is in use first
		if !basicMessage.ReplyTo.IsValid() {
			response.Result = SendResultRECIPIENTVALIDATIONINVALIDREPLYTO
			return response
		}
	}

	if !isValidSubject(basicMessage.Subject) {
		response.Result = SendResultMESSAGEVALIDATIONEMPTYSUBJECT

	} else if !basicMessage.From.IsValid() {
		response.Result = SendResultEMAILADDRESSVALIDATIONINVALIDFROM

	} else if !isValidBasicMessage(basicMessage) {
		response.Result = SendResultEMPTYMESSAGE

	} else if !isValidBasicMessageCustomHeaders(basicMessage) {
		response.Result = SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS

	} else if !isValidBasicMessageMetadata(basicMessage) {
		response.Result = SendResultMESSAGEVALIDATIONINVALIDMETADATA

	} else {
		response.Result = SendResultSUCCESS
		validRecipientCount, allRecipients := isValidBasicRecipientCount(basicMessage)
		if validRecipientCount {
			for _, recipient := range allRecipients {
				if !recipient.IsValid() {
					response.Result = SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS
					addressResult := AddressResult{
						Accepted:     false,
						EmailAddress: recipient.EmailAddress,
						ErrorCode:    "InvalidAddress"}
					response.AddressResults = append(response.AddressResults, addressResult)
				}
			}
		} else {
			response.Result = SendResultNOVALIDRECIPIENTS
		}
	}

	return response
}

//ValidateBulkMessage Validate a basic email message before sending to the Injection API.
func (sendValidator) ValidateBulkMessage(bulkMessage message.BulkMessage) (response SendResponse) {
	response = SendResponse{Result: SendResultUNKNOWNERROR}

	// ReplyTo is an optional property, so check to see if it is in use first
	if len(bulkMessage.ReplyTo.EmailAddress) > 0 {
		if !bulkMessage.ReplyTo.IsValid() {
			response.Result = SendResultRECIPIENTVALIDATIONINVALIDREPLYTO
			return response
		}
	}

	if !isValidSubject(bulkMessage.Subject) {
		response.Result = SendResultMESSAGEVALIDATIONEMPTYSUBJECT
	} else if !bulkMessage.From.IsValid() {
		response.Result = SendResultEMAILADDRESSVALIDATIONINVALIDFROM
	} else if !isValidBulkMessage(bulkMessage) {
		response.Result = SendResultRECIPIENTVALIDATIONINVALIDREPLYTO
	} else if !isValidBulkMessageCustomHeaders(bulkMessage) {
		response.Result = SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS
	} else if !isValidBulkMessageMetadata(bulkMessage) {
		response.Result = SendResultMESSAGEVALIDATIONINVALIDMETADATA
	} else {
		response.Result = SendResultSUCCESS
		validRecipientCount, allRecipients := isValidBulkRecipientCount(bulkMessage)
		if validRecipientCount {
			for _, recipient := range allRecipients {
				if !recipient.IsValid() {
					response.Result = SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS
					addressResult := AddressResult{
						Accepted:     false,
						EmailAddress: recipient.Email,
						ErrorCode:    "InvalidAddress"}
					response.AddressResults = append(response.AddressResults, addressResult)
				}
			}
		} else {
			response.Result = SendResultNOVALIDRECIPIENTS
		}
	}
	return response
}

func isValidSubject(val string) bool {
	return !(len(strings.TrimSpace(val)) == 0)
}

func isValidFromAddress(val string) bool {
	return !(len(strings.TrimSpace(val)) == 0)
}

func isValidBasicMessage(message message.BasicMessage) bool {

	return isValidAPITemplate(message.ApiTemplate) ||
		isValidHTMLBody(message.HtmlBody) ||
		isValidTextBody(message.PlainTextBody)
}

func isValidBulkMessage(message message.BulkMessage) bool {

	return isValidAPITemplate(message.ApiTemplate) ||
		isValidHTMLBody(message.HtmlBody) ||
		isValidTextBody(message.PlainTextBody)
}

func isValidAPITemplate(val string) bool {
	// make sure its actually a number
	templateID, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	return templateID != 0
}

func isValidHTMLBody(val string) bool {
	return !(len(strings.TrimSpace(val)) == 0)
}

func isValidTextBody(val string) bool {
	return !(len(strings.TrimSpace(val)) == 0)
}

func isValidBasicMessageCustomHeaders(message message.BasicMessage) bool {
	if message.CustomHeaders == nil {
		return true
	}
	if len(message.CustomHeaders) == 0 {
		return true
	}
	for _, header := range message.CustomHeaders {
		if !header.IsValid() {
			return false
		}
	}
	return true
}

func isValidBulkMessageCustomHeaders(message message.BulkMessage) bool {
	if message.CustomHeaders == nil {
		return true
	}
	if len(message.CustomHeaders) == 0 {
		return true
	}
	for _, header := range message.CustomHeaders {
		if !header.IsValid() {
			return false
		}
	}
	return true
}

func isValidBasicMessageMetadata(message message.BasicMessage) bool {
	if message.Metadatas == nil {
		return true
	}
	if len(message.Metadatas) == 0 {
		return true
	}
	for _, metadata := range message.Metadatas {
		if !metadata.IsValid() {
			return false
		}
	}
	return true
}

func isValidBulkMessageMetadata(message message.BulkMessage) bool {
	if message.Metadatas == nil {
		return true
	}
	if len(message.Metadatas) == 0 {
		return true
	}
	for _, metadata := range message.Metadatas {
		if !metadata.IsValid() {
			return false
		}
	}
	return true
}

func isValidBasicRecipientCount(message message.BasicMessage) (isValid bool, allRecipients []message.EmailAddress) {
	allRecipients = append(message.To, message.Bcc...)
	allRecipients = append(allRecipients, message.Cc...)
	isValid = len(allRecipients) > 0
	return
}

func isValidBulkRecipientCount(message message.BulkMessage) (isValid bool, allRecipients []message.BulkRecipient) {
	allRecipients = message.To
	isValid = len(allRecipients) > 0
	return
}
