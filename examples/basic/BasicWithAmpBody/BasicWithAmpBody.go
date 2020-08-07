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

	basic := message.BasicMessage{}

	basic.Subject = "Sending a Basic AMP Message"
	basic.HtmlBody = "<html>This HTML will show if AMP is not supported on the receiving end of the email.</html>"
	basic.AmpBody = "<!doctype html>" +
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

	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	//A basic message supports up to 50 recipients and supports several different ways to add recipients
	basic.AddToEmailAddress("recipient1@example.com")  //Add a To address by passing the email address
	basic.AddCcEmailAddress("recipient2@example.com")  //Add a CC address by passing the email address and a friendly name
	basic.AddBccEmailAddress("recipient3@example.com") //Add a BCC address by passing the email address

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
