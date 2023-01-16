package main

import (
	"fmt"

	"github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	basic := message.BasicMessage{}
	basic.Subject = "Sending An Email With Tags"
	basic.HtmlBody = "<body><p><strong>Lorem Ipsum</strong></p></body>"
	basic.PlainTextBody = "Lorem Ipsum"

	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}
	basic.AddToEmailAddress("recipient@example.com")

	tag1 := "tag1"
	tag2 := "tag2"
	basic.Tags = append(basic.Tags, tag1, tag2)

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
