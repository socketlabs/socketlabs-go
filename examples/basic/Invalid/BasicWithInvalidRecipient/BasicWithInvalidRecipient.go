package main

import (
	"fmt"
	"strconv"

	"github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {

	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	basic := message.BasicMessage{}

	basic.Subject = "Sending a Test Message"
	basic.HtmlBody = "<html>This is the Html Body of my message.</html>"
	basic.PlainTextBody = "This is the Plain Text Body of my message."

	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}

	//these will fail validation
	basic.AddToEmailAddress("!@#$!@#$!@#$@#!$")
	basic.AddToEmailAddress("failure.com")
	basic.AddToEmailAddress("ImMissingSomething")
	basic.AddToEmailAddress("Fail@@!.Me")

	//some good addresses
	basic.AddToEmailAddress("this@works")
	basic.AddToEmailAddress("recipient@example.com")

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)

	//The address results will provide the list of specific addresses that failed validation
	for _, address := range sendResponse.AddressResults {
		fmt.Println("Accepted: " + strconv.FormatBool(address.Accepted) + ", " + address.ErrorCode + ", " + address.EmailAddress)
	}
}
