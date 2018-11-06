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
	basic.Subject = "Sending An Email With Custom Headers"
	basic.HtmlBody = "<body><p><strong>Lorem Ipsum</strong></p></body>"
	basic.PlainTextBody = "Lorem Ipsum"

	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}
	basic.AddToEmailAddress("recipient@example.com")

	header1 := message.NewCustomHeader("My-Header", "1...2...3...")
	header2 := message.NewCustomHeader("Example-Type", "BasicSendWithCustomHeaders")
	basic.CustomHeaders = append(basic.CustomHeaders, header1, header2)

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
