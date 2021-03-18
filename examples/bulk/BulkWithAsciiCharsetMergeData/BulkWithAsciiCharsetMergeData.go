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
	bulk := message.BulkMessage{}
	bulk.Subject = "Sending A Test Message"
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com"}

	//Set the content to use MergeData
	bulk.HtmlBody = "<html>" +
		"<head><title>ASCII Merge Data Example</title></head>" +
		"<body>" +
		"<h1>Merge Data</h1>" +
		"<p>Complete? = %%Complete%%</p>" +
		"</body>" +
		"</html>"

	bulk.PlainTextBody = "Merge Data     Complete? = %%Complete%%"

	//Set the character set
	bulk.CharSet = "ASCII"

	//Create recipients with MergeData that uses the character set (✔, ✘)
	var recipient1 = message.NewFriendlyBulkRecipient("recipient1@example.com", "Recipient #1")
	recipient1.MergeData["Complete"] = "✔"

	var recipient2 = message.NewFriendlyBulkRecipient("recipient2@example.com", "Recipient #2")
	recipient2.MergeData["Complete"] = "✔"

	var recipient3 = message.NewFriendlyBulkRecipient("recipient3@example.com", "Recipient #3")
	recipient3.MergeData["Complete"] = "✘"

	//Add the recipients to the message
	bulk.To = append(bulk.To, recipient1, recipient2, recipient3)

	//Send the message
	sendResponse, _ := client.SendBulk(&bulk)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
