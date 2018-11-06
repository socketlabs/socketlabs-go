package message

import "strings"

//BulkRecipient struct
type BulkRecipient struct {
	FriendlyName string
	Email        string
	MergeData    map[string]string
}

// NewBulkRecipient  Factory Constructor
func NewBulkRecipient(email string) BulkRecipient {
	var e BulkRecipient
	e.Email = email
	e.MergeData = make(map[string]string)
	return e
}

// NewFriendlyBulkRecipient  Factory Constructor
func NewFriendlyBulkRecipient(email string, friendlyName string) BulkRecipient {
	var e BulkRecipient
	e.Email = email
	e.FriendlyName = friendlyName
	e.MergeData = make(map[string]string)
	return e
}

// IsValid  Determines if the Email Address is valid
func (e BulkRecipient) IsValid() bool {

	if len(e.Email) == 0 {
		return true
	}

	parts := strings.Split(e.Email, "@")

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

	// check for bad email characters
	if strings.Contains(e.Email, " ") || strings.ContainsAny(e.Email, ",;") {
		return false
	}

	return true
}

// AddGlobalMergeData adds global merge data
func (recipient *BulkRecipient) AddMergeData(key string, value string) *BulkRecipient {
	if recipient.MergeData == nil {
		recipient.MergeData = make(map[string]string)
	}
	recipient.MergeData[key] = value

	return recipient
}
