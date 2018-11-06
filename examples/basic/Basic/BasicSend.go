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

	basic.Subject = "Sending a Basic Message"
	basic.HtmlBody = "<html>This is the Html Body of my message.</html>"
	basic.PlainTextBody = "This is the Plain Text Body of my message."

	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	//A basic message supports up to 50 recipients and supports several different ways to add recipients
	basic.AddToEmailAddress("recipient@example.com")   //Add a To address by passing the email address
	basic.AddCcEmailAddress("recipient2@example.com")  //Add a CC address by passing the email address and a friendly name
	basic.AddBccEmailAddress("recipient3@example.com") //Add a BCC address by passing the email address

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
