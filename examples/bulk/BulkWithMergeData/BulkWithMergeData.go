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
	bulk.Subject = "Sending A BulkSend with MergeData"
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com"}

	//Build the Content (Note the %% symbols used to denote the data to be merged)
	bulk.HtmlBody = "<html>" +
		"<head><title>Merge Data Example</title></head>" +
		"<body>" +
		"<h1>Global Merge Data</h1>" +
		"<p>CompanyMotto = %%Motto%%</p>" +
		"<h1>Per Recipient Merge Data</h1>" +
		"<p>EyeColor = %%EyeColor%%</p>" +
		"<p>HairColor = %%HairColor%%</p>" +
		"</body>" +
		"</html>"

	bulk.PlainTextBody = "Global Merge Data " +
		"CompanyMotto = %%Motto%%" +
		"     " +
		"Per Recipient Merge Data" +
		"     EyeColor = %%EyeColor%%" +
		"     HairColor = %%HairColor%%"

	//Add some global merge data
	bulk.AddGlobalMergeData("Motto", "When hitting the Inbox matters!")

	//Add recipients with merge data
	var recipient1 = message.NewFriendlyBulkRecipient("recipient1@example.com", "Recipient #1")
	recipient1.MergeData["EyeColor"] = "Blue"
	recipient1.MergeData["HairColor"] = "Blond"

	var recipient2 = message.NewFriendlyBulkRecipient("recipient2@example.com", "Recipient #2")
	recipient2.MergeData["EyeColor"] = "Green"
	recipient2.MergeData["HairColor"] = "Brown"

	var recipient3 = message.NewFriendlyBulkRecipient("recipient3@example.com", "Recipient #3")
	recipient2.MergeData["EyeColor"] = "Hazel"
	recipient2.MergeData["HairColor"] = "Black"

	//Add the recipients to the message
	bulk.To = append(bulk.To, recipient1, recipient2, recipient3)

	//Send the message
	sendResponse, _ := client.SendBulk(&bulk)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
