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
	basic.Subject = "Sending An ASCII Charset Email"
	basic.From = message.EmailAddress{EmailAddress: "from@example.com"}
	basic.AddToEmailAddress("recipient@example.com")

	//Set the character set
	basic.CharSet = "ASCII"
	basic.HtmlBody = "<body><strong>Lorem Ipsum</strong>Unicode: ✔ - Check</body>"

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
