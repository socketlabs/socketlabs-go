package message

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Attachment  info
type Attachment struct {
	Content       []byte
	MimeType      string
	Name          string
	ContentID     string
	CustomHeaders []CustomHeader
}

// NewAttachment Factory Constructor that creates attachment using content from specified file and path.
func NewAttachment(filePath string) (Attachment, error) {
	var a Attachment

	// validate passed in filePath here
	if len(filePath) == 0 {
		return a, errors.New("filePath can not be blank")
	}

	// figure out the file extension and mimetype
	_, file := filepath.Split(filePath)
	fileExtension := filepath.Ext(filePath)

	mimeType, err := getMimeTypeForExtension(fileExtension)
	if err != nil {
		return a, err
	}

	// read the file into memory
	contentAsBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return a, err
	}

	// populate the attachment and return
	a.Name = file
	a.MimeType = mimeType
	a.Content = contentAsBytes

	return a, nil
}

// NewAttachmentFromFile Factory Constructor that creates attachment using content from specified file and path.
func NewAttachmentFromFile(name string, mimeType string, filePath string) (Attachment, error) {
	a, err := NewAttachment(filePath)
	if err != nil {
		return a, err
	}

	a.Name = name
	a.MimeType = mimeType

	return a, nil
}

// NewAttachmentFromByteArray  Factory Constructor that creates attachment using content from a byte array
func NewAttachmentFromBytes(name string, mimeType string, content []byte) (Attachment, error) {
	var a Attachment

	// validate passed in filePath here
	if len(mimeType) == 0 {
		return a, errors.New("mimeType can not be blank")
	}

	// populate the attachment and return
	a.Name = name
	a.MimeType = mimeType
	a.Content = content

	return a, nil
}

// Gets the appropriate mime type for a given file extension
func getMimeTypeForExtension(extension string) (string, error) {

	result := ""
	var err error

	// validate passed in extension here
	if len(extension) == 0 {
		err = errors.New("Extension can not be blank")
		return result, err
	}

	extension = strings.Replace(extension, ".", "", 1)

	// return proper mime type for the passed in file extension
	switch strings.ToLower(extension) {
	case "txt", "ini", "sln", "cs", "js", "config", "vb":
		result = "text/plain"
	case "jpeg", "jpg":
		result = "image/jpeg"
	case "xml":
		result = "application/xml"
	case "html":
		result = "text/html"
	case "wav":
		result = "audio/wav"
	case "eml":
		result = "message/rfc822"
	case "mp3":
		result = "audio/mpeg"
	case "mp4":
		result = "video/mp4"
	case "bmp":
		result = "image/bmp"
	case "gif":
		result = "image/gif"
	case "png":
		result = "image/png"
	case "zip":
		result = "application/x-zip-compressed"
	case "doc":
		result = "application/msword"
	case "docx":
		result = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case "xls":
		result = "application/vnd.ms-excel"
	case "xlsx":
		result = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "ppt":
		result = "application/vnd.ms-powerpoint"
	case "pptx":
		result = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case "csv":
		result = "text/csv"
	case "pdf":
		result = "application/pdf"
	case "mov":
		result = "video/quicktime"
	default:
		err = errors.New("Unknown file extension") // handle error here
	}

	return result, err
}

// AddCustomHeader adds a custom header to the attachment
func (attachment *Attachment) AddCustomHeader(name string, value string) {
	customHeader := NewCustomHeader(name, value)
	attachment.CustomHeaders = append(attachment.CustomHeaders, customHeader)
}
