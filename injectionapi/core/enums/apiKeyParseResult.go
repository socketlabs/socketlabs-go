package enums

type ApiParseResult string

const (
	None                             ApiParseResult = "None"
	Invalid                          ApiParseResult = "Invalid"
	InvalidKeyLength                 ApiParseResult = "InvalidKeyLength"
	InvalidKeyFormat                 ApiParseResult = "InvalidKeyFormat"
	InvalidEmptyOrWhitespace         ApiParseResult = "InvalidEmptyOrWhitespace"
	InvalidUnableToExtractPublicPart ApiParseResult = "InvalidUnableToExtractPublicPart"
	InvalidUnableToExtractSecretPart ApiParseResult = "InvalidUnableToExtractSecretPart"
	InvalidPublicPartLength          ApiParseResult = "InvalidPublicPartLength"
	InvalidSecretPartLength          ApiParseResult = "InvalidSecretPartLength"
	Success                          ApiParseResult = "Success"
)
