package injectionapi

//SendResult SendResult enum
type SendResult int

//Parse Find the enumvalue of a string
func (sendResult SendResult) Parse(arg string) SendResult {
	for k, v := range m {
		if v == arg {
			return k
		}
	}
	return -1
}

//ToString Get the string version of the enum
func (sendResult SendResult) ToString() string {
	return m[sendResult]
}

//ToResponseMessage Get the appropraite descriptive ResponseMessage for a given SendResult
func (sendResult SendResult) ToResponseMessage() string {
	switch sendResult {
	case SendResultUNKNOWNERROR:
		return "An error has occured that was unforeseen"
	case SendResultTIMEOUT:
		return "A timeout occurred sending the message"
	case SendResultSUCCESS:
		return "Successful send of message"
	case SendResultWARNING:
		return "Warnings were found while sending the message"
	case SendResultINTERNALERROR:
		return "Internal server error"
	case SendResultMESSAGETOOLARGE:
		return "Message size has exceeded the size limit"
	case SendResultTOOMANYRECIPIENTS:
		return "Message exceeded maximum recipient count in the message"
	case SendResultINVALIDDATA:
		return "Invalid data was found on the message"
	case SendResultOVERQUOTA:
		return "The account is over the send quota, rate limit exceeded"
	case SendResultTOOMANYERRORS:
		return "Too many errors occurred sending the message"
	case SendResultINVALIDAUTHENTICATION:
		return "The ServerId/ApiKey combination is invalid"
	case SendResultACCOUNTDISABLED:
		return "The account has been disabled"
	case SendResultTOOMANYMESSAGES:
		return "Too many messages were found in the request"
	case SendResultNOVALIDRECIPIENTS:
		return "No valid recipients were found in the message"
	case SendResultINVALIDADDRESS:
		return "An invalid recipient were found on the message"
	case SendResultINVALIDATTACHMENT:
		return "An invalid attachment were found on the message"
	case SendResultNOMESSAGES:
		return "No message body was found in the message"
	case SendResultEMPTYMESSAGE:
		return "No message body was found in the message"
	case SendResultEMPTYSUBJECT:
		return "No message body was found in the message"
	case SendResultINVALIDFROM:
		return "An invalid from address was found on the message"
	case SendResultEMPTYTOADDRESS:
		return "No To addresses were found in the message"
	case SendResultNOVALIDBODYPARTS:
		return "No valid message body was found in the message"
	case SendResultINVALIDTEMPLATEID:
		return "An invalid TemplateId was found in the message"
	case SendResultTEMPLATEHASNOCONTENT:
		return "The specified TemplateId has no content for the message"
	case SendResultMESSAGEBODYCONFLICT:
		return "A conflict occurred on the message body of the message"
	case SendResultINVALIDMERGEDATA:
		return "Invalid MergeData was found on the message"
	case SendResultAUTHENTICATIONVALIDATIONFAILED:
		return "SDK Validation Error : Authentication Validation Failed, Missing or invalid ServerId or ApiKey"
	case SendResultRECIPIENTVALIDATIONMAXEXCEEDED:
		return "SDK Validation Error : Message exceeded maximum recipient count in the message"
	case SendResultRECIPIENTVALIDATIONNONEINMESSAGE:
		return "SDK Validation Error : No Recipients were found in the message"
	case SendResultEMAILADDRESSVALIDATIONMISSINGFROM:
		return "SDK Validation Error : From email address is missing in the message"
	case SendResultRECIPIENTVALIDATIONMISSINGTO:
		return "SDK Validation Error : To addresses are missing in the message"
	case SendResultEMAILADDRESSVALIDATIONINVALIDFROM:
		return "SDK Validation Error : From email address in the message is invalid"
	case SendResultMESSAGEVALIDATIONEMPTYSUBJECT:
		return "SDK Validation Error : No Subject was found in the message"
	case SendResultMESSAGEVALIDATIONEMPTYMESSAGE:
		return "SDK Validation Error : No message body was found in the message"
	case SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS:
		return "SDK Validation Error : Invalid Custom Headers were found in the message"
	case SendResultMESSAGEVALIDATIONINVALIDMETADATA:
		return "SDK Validation Error : Invalid Meta Data was found in the message"
	case SendResultRECIPIENTVALIDATIONINVALIDREPLYTO:
		return "SDK Validation Error : Invalid ReplyTo Address was found in the message"
	case SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS:
		return "SDK Validation Error : Invalid recipients were found in the message"
	default:
		return ""
	}
}

var m = map[SendResult]string{
	SendResultUNKNOWNERROR:                          "UnknownError",
	SendResultTIMEOUT:                               "Timeout",
	SendResultSUCCESS:                               "Success",
	SendResultWARNING:                               "Warning",
	SendResultINTERNALERROR:                         "InternalError",
	SendResultMESSAGETOOLARGE:                       "MessageTooLarge",
	SendResultTOOMANYRECIPIENTS:                     "TooManyRecipients",
	SendResultINVALIDDATA:                           "InvalidData",
	SendResultOVERQUOTA:                             "OverQuota",
	SendResultTOOMANYERRORS:                         "TooManyErrors",
	SendResultINVALIDAUTHENTICATION:                 "InvalidAuthentication",
	SendResultACCOUNTDISABLED:                       "AccountDisabled",
	SendResultTOOMANYMESSAGES:                       "TooManyMessages",
	SendResultNOVALIDRECIPIENTS:                     "NoValidRecipients",
	SendResultINVALIDADDRESS:                        "InvalidAddress",
	SendResultINVALIDATTACHMENT:                     "InvalidAttachment",
	SendResultNOMESSAGES:                            "NoMessages",
	SendResultEMPTYMESSAGE:                          "EmptyMessage",
	SendResultEMPTYSUBJECT:                          "EmptySubject",
	SendResultINVALIDFROM:                           "InvalidFrom",
	SendResultEMPTYTOADDRESS:                        "EmptyToAddress",
	SendResultNOVALIDBODYPARTS:                      "NoValidBodyParts",
	SendResultINVALIDTEMPLATEID:                     "InvalidTemplateID",
	SendResultTEMPLATEHASNOCONTENT:                  "TemplateHasNoContent",
	SendResultMESSAGEBODYCONFLICT:                   "MessageBodyConflict",
	SendResultINVALIDMERGEDATA:                      "InvalidMergeData",
	SendResultAUTHENTICATIONVALIDATIONFAILED:        "AuthenticationValidationFailed",
	SendResultEMAILADDRESSVALIDATIONMISSINGFROM:     "EmailAddressValidationMissingFrom",
	SendResultEMAILADDRESSVALIDATIONINVALIDFROM:     "EmailAddressValidationInvalidFrom",
	SendResultRECIPIENTVALIDATIONMAXEXCEEDED:        "RecipientValidationMaxExceeded",
	SendResultRECIPIENTVALIDATIONNONEINMESSAGE:      "RecipientValidationNoneInMessage",
	SendResultRECIPIENTVALIDATIONMISSINGTO:          "RecipientValidationMissingTo",
	SendResultRECIPIENTVALIDATIONINVALIDREPLYTO:     "RecipientValidationInvalidReplyTo",
	SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS:  "RecipientValidationInvalidRecipients",
	SendResultMESSAGEVALIDATIONEMPTYSUBJECT:         "MessageValidationEmptySubject",
	SendResultMESSAGEVALIDATIONEMPTYMESSAGE:         "MessageValidationEmptyMessage",
	SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS: "MessageValidationInvalidCustomHeaders",
	SendResultMESSAGEVALIDATIONINVALIDMETADATA:		 "MessageValidationInvalidMetadata"}

const (
	//SendResultUNKNOWNERROR - An error has occured that was unforeseen
	SendResultUNKNOWNERROR SendResult = iota

	//SendResultTIMEOUT - A timeout occurred sending the message
	SendResultTIMEOUT

	//SendResultSUCCESS - Successful send of message
	SendResultSUCCESS

	//SendResultWARNING - Warnings were found while sending the message
	SendResultWARNING

	//SendResultINTERNALERROR - Internal server error
	SendResultINTERNALERROR

	//SendResultMESSAGETOOLARGE - Message size has exceeded the size limit
	SendResultMESSAGETOOLARGE

	//SendResultTOOMANYRECIPIENTS - Message exceeded maximum recipient count in the message
	SendResultTOOMANYRECIPIENTS

	//SendResultINVALIDDATA - Invalid data was found on the message
	SendResultINVALIDDATA

	//SendResultOVERQUOTA - The account is over the send quota rate limit exceeded
	SendResultOVERQUOTA

	//SendResultTOOMANYERRORS - Too many errors occurred sending the message
	SendResultTOOMANYERRORS

	//SendResultINVALIDAUTHENTICATION - The ServerId/ApiKey combination is invalid
	SendResultINVALIDAUTHENTICATION

	//SendResultACCOUNTDISABLED - The account has been disabled
	SendResultACCOUNTDISABLED

	//SendResultTOOMANYMESSAGES - Too many messages were found in the request
	SendResultTOOMANYMESSAGES

	//SendResultNOVALIDRECIPIENTS - No valid recipients were found in the message
	SendResultNOVALIDRECIPIENTS

	//SendResultINVALIDADDRESS - An invalid recipient were found on the message
	SendResultINVALIDADDRESS

	//SendResultINVALIDATTACHMENT - An invalid attachment were found on the message
	SendResultINVALIDATTACHMENT

	//SendResultNOMESSAGES - No message body was found in the message
	SendResultNOMESSAGES

	//SendResultEMPTYMESSAGE - No message body was found in the message
	SendResultEMPTYMESSAGE

	//SendResultEMPTYSUBJECT - No subject was found in the message
	SendResultEMPTYSUBJECT

	//SendResultINVALIDFROM - An invalid from address was found on the message
	SendResultINVALIDFROM

	//SendResultEMPTYTOADDRESS - No To addresses were found in the message
	SendResultEMPTYTOADDRESS

	//SendResultNOVALIDBODYPARTS - No valid message body was found in the message
	SendResultNOVALIDBODYPARTS

	//SendResultINVALIDTEMPLATEID - An invalid TemplateId was found in the message
	SendResultINVALIDTEMPLATEID

	//SendResultTEMPLATEHASNOCONTENT - The specified TemplateId has no content for the message
	SendResultTEMPLATEHASNOCONTENT

	//SendResultMESSAGEBODYCONFLICT - A conflict occurred on the message body of the message
	SendResultMESSAGEBODYCONFLICT

	//SendResultINVALIDMERGEDATA - Invalid MergeData was found on the message
	SendResultINVALIDMERGEDATA

	//SendResultAUTHENTICATIONVALIDATIONFAILED - Authentication Error Missing or invalid ServerId or ApiKey
	SendResultAUTHENTICATIONVALIDATIONFAILED

	//SendResultEMAILADDRESSVALIDATIONMISSINGFROM - From email address is missing in the message
	SendResultEMAILADDRESSVALIDATIONMISSINGFROM

	//SendResultEMAILADDRESSVALIDATIONINVALIDFROM - From email address in the message in invalid
	SendResultEMAILADDRESSVALIDATIONINVALIDFROM

	//SendResultRECIPIENTVALIDATIONMAXEXCEEDED - Message exceeded maximum recipient count in the message
	SendResultRECIPIENTVALIDATIONMAXEXCEEDED

	//SendResultRECIPIENTVALIDATIONNONEINMESSAGE - No recipients were found in the message
	SendResultRECIPIENTVALIDATIONNONEINMESSAGE

	//SendResultRECIPIENTVALIDATIONMISSINGTO - To addresses are missing in the message
	SendResultRECIPIENTVALIDATIONMISSINGTO

	//SendResultRECIPIENTVALIDATIONINVALIDREPLYTO - Invalid ReplyTo address found
	SendResultRECIPIENTVALIDATIONINVALIDREPLYTO

	//SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS - Invalid recipients found
	SendResultRECIPIENTVALIDATIONINVALIDRECIPIENTS

	//SendResultMESSAGEVALIDATIONEMPTYSUBJECT - No subject was found in the message
	SendResultMESSAGEVALIDATIONEMPTYSUBJECT

	//SendResultMESSAGEVALIDATIONEMPTYMESSAGE - No message body was found in the message
	SendResultMESSAGEVALIDATIONEMPTYMESSAGE

	//SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS -  Invalid custom headers found
	SendResultMESSAGEVALIDATIONINVALIDCUSTOMHEADERS

	//SendResultMESSAGEVALIDATIONINVALIDMETADATA -  Invalid meta data found
	SendResultMESSAGEVALIDATIONINVALIDMETADATA
)
