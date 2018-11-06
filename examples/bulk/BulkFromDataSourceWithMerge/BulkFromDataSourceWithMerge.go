package main

import (
	"fmt"

	"github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/examples/Integration/Repository"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Retrieve data from the datasource
	data := testdata.GetCustomers()

	//Build the message
	bulk := message.BulkMessage{}

	bulk.Subject = "Hello %%FirstName%%"
	bulk.PlainTextBody = "Hello %%FirstName%% %%LastName%%. Is your favorite color still %%FavoriteColor%%?"
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com"}
	bulk.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	//Merge in the customers from the datasource
	for _, customer := range data {
		var recipient = bulk.AddToBulkRecipient(customer.EmailAddress)
		recipient.MergeData["FirstName"] = customer.FirstName
		recipient.MergeData["LastName"] = customer.LastName
		recipient.MergeData["FavoriteColor"] = customer.FavoriteColor
	}

	//Send the message
	sendResponse, _ := client.SendBulk(&bulk)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
