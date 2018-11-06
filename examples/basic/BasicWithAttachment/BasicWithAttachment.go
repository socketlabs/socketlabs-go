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
	basic := message.BasicMessage{}
	basic.Subject = "Sending an Attachment"
	basic.HtmlBody = "<html>This is the Html Body of my message.</html>"
	basic.PlainTextBody = "This is the Plain Text Body of my message."
	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.AddToEmailAddress("recipient@socketlabs.com")

	//Add attachment (with optional headers if desired)
	attachment, _ := message.NewAttachment("../../Img/bus.png")
	attachment.CustomHeaders = append(attachment.CustomHeaders, message.CustomHeader{Name: "header1", Value: "value1"})
	attachment.CustomHeaders = append(attachment.CustomHeaders, message.CustomHeader{Name: "header2", Value: "value2"})
	basic.Attachments = append(basic.Attachments, attachment)

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
