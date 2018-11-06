package message

import (
	"strings"
)

// EmailAddress  info
type EmailAddress struct {
	FriendlyName string
	EmailAddress string
}

// NewEmailAddress  Factory Constructor
func NewEmailAddress(email string) EmailAddress {
	var e EmailAddress
	e.EmailAddress = email
	return e
}

// NewFriendlyEmailAddress  Factory Constructor
func NewFriendlyEmailAddress(email string, friendlyName string) EmailAddress {
	var e EmailAddress
	e.EmailAddress = email
	e.FriendlyName = friendlyName
	return e
}

// IsValid  Determines if the Email Address is valid
func (e EmailAddress) IsValid() bool {
	if len(e.EmailAddress) == 0 {
		return false
	}

	parts := strings.Split(e.EmailAddress, "@")

	if len(parts) != 2 {
		return false
	}

	//320 used over 256 to be more lenient
	if len(parts) > 320 {
		return false
	}

	if len(parts[0]) < 1 {
		return false
	}

	if len(parts[1]) < 1 {
		return false
	}

	// chech for bad email characters
	if strings.Contains(e.EmailAddress, " ") || strings.ContainsAny(e.EmailAddress, ",;") {
		return false
	}

	return true
}
