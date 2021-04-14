[![SocketLabs](https://static.socketlabs.com/logos/logo-dark-326x64.png)](https://www.socketlabs.com/developers)
# [![Twitter Follow](https://img.shields.io/twitter/follow/socketlabs.svg?style=social&label=Follow)](https://twitter.com/socketlabs) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/socketlabs/socketlabs-nodejs/blob/master/CONTRIBUTING.md)

The SocketLabs Email Delivery Go package allows you to easily send email messages via the [SocketLabs Injection API](https://www.socketlabs.com/docs/inject/).  The library makes it easy to build and send any type of message supported by the API, from a simple message to a single recipient all the way to a complex bulk message sent to a group of recipients with unique merge data per recipient.

# Table of Contents
* [Prerequisites and Installation](#prerequisites-and-installation)
* [Getting Started](#getting-started)
* [Managing API Keys](#managing-api-keys)
* [Examples and Use Cases](#examples-and-use-cases)
* [License](#license)


<a name="prerequisites-and-installation"></a>
# Prerequisites and Installation
## Prerequisites
* Go 1.2.2 or higher
* A SocketLabs account. If you don't have one yet, you can [sign up for a free account](https://signup.socketlabs.com/step-1?plan=free) to get started.

## Installation
```
go get github.com/socketlabs/socketlabs-go/injectionapi
```


<a name="getting-started"></a>
# Getting Started
## Obtaining your API Key and SocketLabs ServerId number
In order to get started, you'll need to enable the Injection API feature in the [SocketLabs Control Panel](https://cp.socketlabs.com).
Once logged in, navigate to your SocketLabs server's dashboard (if you only have one server on your account you'll be taken here immediately after logging in).
Make note of your 4 or 5 digit ServerId number, as you'll need this along with
your API key in order to use the Injection API.

To enable the Injection API, click on the "For Developers" dropdown on the top-level navigation, then choose the "Configure HTTP Injection API" option.
Once here, you can enable the feature by choosing the "Enabled" option in the
dropdown. Enabling the feature will also generate your API key, which you'll
need (along with your ServerId) to start using the API. Be sure to click the
"Update" button to save your changes once you are finished.

## Basic Message
A basic message is an email message like you'd send from a personal email client such as Outlook.
A basic message can have many recipients, including multiple To addresses, CC addresses, and even BCC addresses.
You can also send a file attachment in a basic message.

```GO

import (
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {
	client := injectionapi.CreateClient(000001, "YOUR-API-KEY")

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
	_, _ = client.SendBasic(&basic)
}
```

## Bulk Message
A bulk message usually contains a single recipient per message
and is generally used to send the same content to many recipients,
optionally customizing the message via the use of MergeData.
For more information about using Merge data, please see the [Injection API documentation](https://www.socketlabs.com/docs/inject/).
```GO
import (
	"github.com/socketlabs/socketlabs-go/injectionapi"
	"github.com/socketlabs/socketlabs-go/injectionapi/message"
)

func main() {

	client := injectionapi.CreateClient(000001, "YOUR-API-KEY")

	//Build the message
	bulk := message.BulkMessage{}

	bulk.Subject = "Sending A Bulk Message"
	bulk.HtmlBody = "<html>This is the Html Body of my message sent to the recipient with %%EyeColor%% eyes.</html>";
	bulk.PlainTextBody = "This is the Plain Text Body of my message sent to the recipient with %%EyeColor%% eyes.";
	bulk.From = message.EmailAddress{EmailAddress: "from@example.com"}
	bulk.ReplyTo = message.EmailAddress{EmailAddress: "replyto@example.com"}

	//Add Bulk Recipients
	recipient1 := bulk.AddToFriendlyBulkRecipient("recipient1@example.com", "Recipient #1")
	recipient1.AddMergeData("EyeColor", "Green")

	recipient2 := bulk.AddToFriendlyBulkRecipient("recipient2@example.com", "Recipient #2")
	recipient2.AddMergeData("EyeColor", "Blue")

	//Send the message
	_, _ = client.SendBulk(&bulk)
}
```

<a name="managing-api-keys"></a>
## Managing API Keys
For ease of demonstration, some of our examples may include the ServerId and API key directly in our code sample. Generally it is not considered a good practice to store sensitive information like this directly in your code. In most cases we recommend the use of [Environment Variables](https://flaviocopes.com/golang-environment-variables/).

<a name="examples-and-use-cases"></a>
# Examples and Use Cases
In order to demonstrate the many possible use cases for the SDK, we've provided
an assortment of code examples. These examples demonstrate many different
features available to the Injection API and SDK, including using templates
created in the [SocketLabs Email Designer](https://www.socketlabs.com/blog/introducing-new-email-designer/), custom email headers, sending
attachments, sending content that is stored in an HTML file, advanced bulk
merging, and even pulling recipients from a datasource.


### [Basic send example](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/Basic/BasicSend.go)
This example demonstrates a Basic Send.

### [Basic send complex example](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicComplex/BasicSendComplex.go)
This example demonstrates many features of the Basic Send, including adding multiple recipients, adding message and mailing id's, and adding an embedded image.

### [Basic send from HTML file](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicFromHtmlFile/BasicFromHtmlFile.go)
This example demonstrates how to read in your HTML content from an HTML file
rather than passing in a string directly.

### [Basic send from SocketLabs API Template](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicWithApiTemplate/BasicWithApiTemplate.go)
This example demonstrates the sending of a piece of content that was created in the
SocketLabs Email Designer. This is also known as the [API Templates](https://www.socketlabs.com/blog/introducing-api-templates/) feature.

### [Basic send with specified character set](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicWithAsciiCharset/BasicWithAsciiCharset.go)
This example demonstrates sending with a specific character set.

### [Basic send with file attachment](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicWithAttachment/BasicWithAttachment.go)
This example demonstrates how to add a file attachment to your message.

### [Basic send with custom email headers](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicWithCustomHeaders/BasicWithCustomHeaders.go)
This example demonstrates how to add custom headers to your email message.

### [Basic send with embedded image](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicWithEmbeddedImage/BasicWithEmbeddedImage.go)
This example demonstrates how to embed an image in your message.

### [Basic send with a web proxy](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicSendWithProxy/BasicSendWithProxy.go)
This example demonstrates how to use a proxy with your HTTP client.
### [Basic send with retry enabled](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/BasicSendWithRetry/BasicSendWithRetry.go)
This example demonstrates how to use the retry logic with your HTTP client.

### [Basic send with Amp ](https://github.com/socketlabs/socketlabs-go/blob/main/examples/basic/BasicWithAmpBody/BasicWithAmpBody.go)
This example demonstrates how to send a basic message with an AMP Html body.
For more information about AMP please see [AMP Project](https://amp.dev/documentation/)

### [Basic send with invalid file attachment](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/Invalid/BasicWithInvalidAttachment/BasicWithInvalidAttachment.go)
This example demonstrates the results of attempting to do a send with an invalid attachment.

### [Basic send with invalid from address](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/Invalid/BasicWithInvalidFrom/BasicWithInvalidFrom.go)
This example demonstrates the results of attempting to do a send with an invalid from address.

### [Basic send with invalid recipients](https://github.com/socketlabs/socketlabs-go/blob/master/examples/basic/Invalid/BasicWithInvalidRecipient/BasicWithInvalidRecipient.go)
This example demonstrates the results of attempting to do a send with invalid recipients.

### [Bulk send with multiple recipients](https://github.com/socketlabs/socketlabs-go/blob/master/examples/bulk/bulk/bulk.go)
This example demonstrates how to send a bulk message to multiple recipients.

### [Bulk send with complex merge including attachments](https://github.com/socketlabs/socketlabs-go/blob/master/examples/bulk/BulkComplex/BulkComplex.go)
This example demonstrates many features of the `BulkMessage()`, including
adding multiple recipients, merge data, and adding an attachment.

### [Bulk send with recipients pulled from a datasource](https://github.com/socketlabs/socketlabs-go/blob/master/examples/bulk/BulkFromDataSourceWithMerge/BulkFromDataSourceWithMerge.go)
This example uses a mock repository class to demonstrate how you would pull
your recipients from a database and create a bulk mailing with merge data.

### [Bulk send with Ascii charset and special characters](https://github.com/socketlabs/socketlabs-go/blob/master/examples/bulk/BulkWithAsciiCharsetMergeData/BulkWithAsciiCharsetMergeData.go)
This example demonstrates how to send a bulk message with a specified character
set and special characters.

### [Bulk send with merge data](https://github.com/socketlabs/socketlabs-go/blob/master/examples/bulk/BulkWithMergeData/BulkWithMergeData.go)
This example demonstrates how to send a bulk message to multiple recipients with
unique merge data per recipient.

### [Bulk send with Amp ](https://github.com/socketlabs/socketlabs-go/blob/main/examples/bulk/BulkWithAmpBody/BulkWithAmpBody.go)
This example demonstrates how to send a bulk message with an AMP Html body.
For more information about AMP please see [AMP Project](https://amp.dev/documentation/)


<a name="version"></a>
# Version
* 1.2.1 - Adding optional retry logic for Http requests. If configured, the request will retry when certain 500 errors occur (500, 502, 503, 504)
* 1.1.1 - Adding request timeout value on the client for Http requests
* 1.1.0 - Adds Amp Html Support
* 1.0.0 - Initial Release


<a name="license"></a>
# License
The SocketLabs.EmailDelivery library and all associated code, including any code samples, are [MIT Licensed](https://github.com/socketlabs/socketlabs-go/blob/master/LICENSE.MD).
