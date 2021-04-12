package main

import (
	"fmt"

	"github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {

	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	bulk := message.BulkMessage{}

	bulk.Subject = "Sending A Test Message"
	bulk.HtmlBody = "<html>This is the Html Body of my message.</html>"
	bulk.PlainTextBody = "This is the Plain Text Body of my message."
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com"}
	bulk.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	//Add Bulk Recipients
	bulk.AddToFriendlyBulkRecipient("recipient1@example.com", "Recipient #1")
	bulk.AddToFriendlyBulkRecipient("recipient2@example.com", "Recipient #2")
	bulk.AddToFriendlyBulkRecipient("recipient3@example.com", "Recipient #3")
	bulk.AddToFriendlyBulkRecipient("recipient4@example.com", "Recipient #4")

	//Send the message
	sendResponse, _ := client.SendBulk(&bulk)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
