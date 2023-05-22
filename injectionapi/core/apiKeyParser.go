package core

import (
	"strings"
)

// Parses a provided api key and provides a result.
type ApiKeyParser struct{}

func (parser ApiKeyParser) Parse(wholeApiKey string) (response string) {

	if len(strings.TrimSpace(wholeApiKey)) == 0 {
		response = "InvalidEmptyOrWhitespace"
		return
	}

	if len(strings.TrimSpace(wholeApiKey)) != 61 {
		response = "InvalidKeyLength"
		return
	}

	splitIndex := strings.Index(wholeApiKey, ".")

	if splitIndex == -1 || splitIndex > 50 {
		response = "InvalidKeyFormat"
		return
	}

	apiKeyRunes := []rune(wholeApiKey)

	publicPart := string(apiKeyRunes[0:splitIndex])

	if len(strings.TrimSpace(publicPart)) != 20 {
		response = "InvalidPublicPartLength"
		return
	}

	splitIndex++

	secretPart := string(apiKeyRunes[splitIndex:len(wholeApiKey)])

	if len(strings.TrimSpace(secretPart)) != 40 {
		response = "InvalidSecretPartLength"
		return
	}

	response = "Success"
	return
}
