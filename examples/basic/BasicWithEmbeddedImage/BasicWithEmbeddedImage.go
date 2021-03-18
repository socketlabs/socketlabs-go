package main

import (
	"fmt"

	"github.com/PraneethChandraThota/socketlabs-go/examples"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	basic := message.BasicMessage{}
	basic.Subject = "Sending an Email With An Embedded Image"
	basic.PlainTextBody = "Lorem Ipsum"
	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.AddToFriendlyEmailAddress("recipient@example.com", "Recipient")

	//Add an image attachment
	attachment, _ := message.NewAttachment("../../Img/bus.png")
	attachment.ContentID = "bus" // specify the cid as 'Bus' so it can be embedded
	basic.Attachments = append(basic.Attachments, attachment)

	//Display the attachment in the html body
	basic.HtmlBody = "<body><p><strong>Lorem Ipsum</strong></p><br /><img src=\"cid:Bus\" /></body>"

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
