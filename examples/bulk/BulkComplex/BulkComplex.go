package main

import (
	"bytes"
	"fmt"

	"github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	bulk := message.BulkMessage{}

	// Add some global merge-data (These will be applied to all Recipients unless specifically overridden by Recipient level merge)
	bulk.AddGlobalMergeData("Motto", "When hitting the Inbox Matters!")
	bulk.AddGlobalMergeData("Birthday", "Unknown")
	bulk.AddGlobalMergeData("Age", "An unknown number of")
	bulk.AddGlobalMergeData("Upsell", "BTW:  You are eligible for discount pricing when you upgrade your service!")

	//Add Bulk Recipients
	recipient1 := bulk.AddToBulkRecipient("recipient1@example.com")
	recipient1.AddMergeData("Birthday", "08/05/1991") // this will override the global merge data for Birthday for this specific Recipient
	recipient1.AddMergeData("Age", "27")              // this will override the global merge data for Age for this specific Recipient

	recipient2 := bulk.AddToBulkRecipient("recipient2@example.com")
	recipient2.AddMergeData("Birthday", "04/12/1984") // this will override the global merge data for Birthday for this specific Recipient
	recipient2.AddMergeData("Age", "34")              // this will override the global merge data for Age for this specific Recipient
	recipient2.AddMergeData("UpSell", "")             // This will override the Global Merge-Data for this specific Recipient

	recipient3 := bulk.AddToBulkRecipient("recipient3@example.com")
	recipient3.FriendlyName = "Recipient 3"

	bulk.AddToFriendlyBulkRecipient("recipient2@example.com", "Recipient #2")
	bulk.AddToFriendlyBulkRecipient("recipient3@example.com", "Recipient #3")
	bulk.AddToFriendlyBulkRecipient("recipient4@example.com", "Recipient #4")

	//Add recipients with merge data
	bulk.Subject = "Complex BulkSend Example"
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com", FriendlyName: "FromMe"}
	bulk.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	bulk.AddCustomHeader("testMessageHeader", "I am a message header")

	var htmlBody bytes.Buffer
	htmlBody.WriteString("<html>")
	htmlBody.WriteString("   <head><title>Complex</title></head>")
	htmlBody.WriteString("   <body>")
	htmlBody.WriteString("       <h1>Merged Data</h1>")
	htmlBody.WriteString("       <p>")
	htmlBody.WriteString("           Motto = <b>%%Motto%%</b> </br>")
	htmlBody.WriteString("           Birthday = <b>%%Birthday%%</b> </br>")
	htmlBody.WriteString("           Age = <b>%%Age%%</b> </br>")
	htmlBody.WriteString("           UpSell = <b>%%UpSell%%</b> </br>")
	htmlBody.WriteString("       </p>")
	htmlBody.WriteString("       </br>")
	htmlBody.WriteString("       <h1>Example of Merge Usage</h1>")
	htmlBody.WriteString("       <p>")
	htmlBody.WriteString("           Our company motto is '<b>%%Motto%%</b>'. </br>")
	htmlBody.WriteString("           Your birthday is <b>%%Birthday%%</b> and you are <b>%%Age%%</b> years old. </br>")
	htmlBody.WriteString("            </br>")
	htmlBody.WriteString("           <b>%%UpSell%%<b>")
	htmlBody.WriteString("       </p>")
	htmlBody.WriteString("   </body>")
	htmlBody.WriteString("</html>")
	bulk.HtmlBody = htmlBody.String()

	//skipping plain text for this example
	//bulk.PlainTextBody = "Some Text"

	attachment, _ := message.NewAttachment("../../Img/bus.png")
	attachment.AddCustomHeader("Atachment-Header", "I Am A Bus")
	bulk.Attachments = append(bulk.Attachments, attachment)

	//Send the message
	sendResponse, _ := client.SendBulk(&bulk)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
