package main

import (
	"fmt"

	exampleconfig "github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

// For more info on AMP Html, visit: https://amp.dev/
func main() {

	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	bulk := message.BulkMessage{}

	bulk.Subject = "Sending A Bulk AMP Message"
	bulk.HtmlBody = "<html>This HTML will show if AMP is not supported on the receiving end of the email.</html>"
	bulk.AmpBody = "<!doctype html>" +
		"<html amp4email>" +
		"    <head>" +
		"    <meta charset=\"utf-8\">" +
		"     <script async src=\"https://cdn.ampproject.org/v0.js\"></script>" +
		"    <style amp4email-boilerplate>body{visibility:hidden}</style>" +
		"     <style amp-custom>" +
		"         h1 {" +
		"              margin: 1rem;" +
		"            }" +
		"      </style>" +
		"   </head>" +
		"        <body>" +
		"         <h1>This is the AMP Html Body of my message</h1>" +
		"        </body>" +
		"</html>"
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
