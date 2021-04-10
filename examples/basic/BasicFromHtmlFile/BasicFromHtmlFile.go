package main

import (
	"fmt"
	"io/ioutil"

	"github.com/PraneethChandraThota/socketlabs-go/examples"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	basic := message.BasicMessage{}

	basic.Subject = "Simple Html file with text"
	basic.PlainTextBody = "This is the Plain Text Body of my message."
	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.AddToFriendlyEmailAddress("recipient1@example.com", "Recipient #1")

	//Get html content from a file
	pathToFile := "../../Html/SampleEmail.html"
	htmlAsBytes, _ := ioutil.ReadFile(pathToFile)
	basic.HtmlBody = string(htmlAsBytes)

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
