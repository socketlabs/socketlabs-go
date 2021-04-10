package main

import (
	"fmt"

	"github.com/PraneethChandraThota/socketlabs-go/examples"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/message"
)

func main() {

	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	basic := message.BasicMessage{}

	basic.Subject = "Sending Using a Template"
	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.AddToFriendlyEmailAddress("recipient1@example.com", "Recipient #1")

	//Pass in the ID of the template to use as the htmlBody
	basic.ApiTemplate = "1"

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
