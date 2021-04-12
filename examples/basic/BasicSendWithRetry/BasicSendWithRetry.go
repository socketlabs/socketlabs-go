package main

import (
	"fmt"
	"net"

	"github.com/PraneethChandraThota/socketlabs-go/examples"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/message"
)

func main() {

	//Create a client with proxy
	//client := injectionapi.CreateClientWithProxy(exampleconfig.ServerId(), exampleconfig.APIKey(), exampleconfig.ProxyURL())
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())
	client.SetRequestTimeout(10)
	client.SetNumberOfRetries(2)
	client.SetEndpointURL("http://127.0.0.1:5000/")
	fmt.Println("Number of Retries set : ", client.GetNumberOfRetries())


	//Or Create the client and then set the proxy on the client
	// client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())
	// client.SetProxyURL(exampleconfig.ProxyURL())

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
	sendResponse, err := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)

	if err != nil {
		fmt.Println("Error : ", err)
		err, ok := err.(net.Error)
		if ok && err.Timeout() {
			fmt.Println("Time out error occured : ", err)
		}
	}
}
