package main

import (
	"fmt"

	exampleconfig "github.com/socketlabs/socketlabs-go/examples"
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {

	client := injectionapi.CreateClient(exampleconfig.ServerId(), exampleconfig.APIKey())

	//Build the message
	basic := message.BasicMessage{}
	basic.Subject = "Sending Basic Complex Example"

	//A basic message supports up to 50 recipients per send and supports several different ways to add recipients
	basic.From = message.EmailAddress{EmailAddress: "from@example.com", FriendlyName: "from"}
	basic.AddToEmailAddress("recipient@example.com")             //Add a To address by passing the email address
	basic.AddCcFriendlyEmailAddress("cc@example.com", "cc name") //Add a CC address by passing the email address and a friendly name
	basic.AddBccEmailAddress("bcc@example.com")                  //Add a BCC address by passing the email address

	//Set reply to
	basic.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com", FriendlyName: "Reply Address"}

	//set html, amp, and text body parts
	basic.HtmlBody = "<body><p><strong>Html Body</strong></p><br /><img src=\"cid:Bus\" /></body>"
	basic.PlainTextBody = "Lorem Ipsum"
	basic.AmpBody = "<!doctype html>" +
		"<html amp4email>" +
		"    <head>" +
		"    <meta charset=\"utf-8\">" +
		"     <script async src=\"https://cdn.ampproject.org/v0.js\"></script>" +
		"    <style amp4email-boilerplate>body{visibility:hidden}</style>" +
		"     <style amp-custom>" +
		"         h1 {" +
		"              margin: 1rem;" +
		"            }" +
		"      </style>" +
		"   </head>" +
		"        <body>" +
		"         <h1>This is the AMP Html Body of my message</h1>" +
		"        </body>" +
		"</html>"

	basic.CharSet = "utf-8"

	//Tag message with mailing and message ids.
	//See https://support.socketlabs.com/index.php/Knowledgebase/Article/View/48/2/using-mailingcampaign-ids-and-message-ids-with-socketlabs-email-on-demand
	basic.MailingId = "My Mailing Id" //Identifier for groups of messages such as campaigns, jobs, or batches of messages.
	basic.MessageId = "My Message Id" //Typically used to identify a specific message or a recipient of a message.

	//Configure custom message headers
	basic.CustomHeaders = append(basic.CustomHeaders, message.NewCustomHeader("MyMessageHeader", "I am a message header"))

	//Configure message meta data
	basic.Metadata = append(basic.Metadata, message.NewMetadata("MyMetaDatra", "I am meta data"))

	//Add message tags
	basic.Tags = append(basic.Tags, "I am a Tag")
	basic.Tags = append(basic.Tags, "go-Example")
	basic.Tags = append(basic.Tags, "Basic-Example")

	//Send the message
	sendResponse, _ := client.SendBasic(&basic)

	//Output the results
	fmt.Println(sendResponse.Result.ToString())
	fmt.Println(sendResponse.ResponseMessage)
}
