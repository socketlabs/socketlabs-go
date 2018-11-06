package exampleconfig

import (
	"os"
	"strconv"
)

func ServerId() int {

	envVar := os.Getenv("SOCKETLABS_SERVER_ID")

	//if environment variable does not exist, crash
	if envVar == "" {
		panic("Missing SOCKETLABS_SERVER_ID environment variable. This should be the integer value of your Socketlabs server id. " +
			"See https://github.com/socketlabs/socketlabs-go/blob/master/README.md#obtaining-your-api-key-and-socketlabs-serverid-number")
	}

	//serverId
	serverid, err := strconv.Atoi(envVar)
	if err != nil {
		panic(err)
	}

	return serverid
}

func APIKey() string {

	envVar := os.Getenv("SOCKETLABS_INJECTION_API_KEY")
	//if environment variable does not exist, crash
	if envVar == "" {
		panic("Missing SOCKETLABS_INJECTION_API_KEY environment variable. This should be the string value of your Socketlabs injection api key. " +
			"See https://github.com/socketlabs/socketlabs-go/blob/master/README.md#obtaining-your-api-key-and-socketlabs-serverid-number")
	}
	return envVar
}

func EndpointURL() string {
	return "https://inject.socketlabs.com/api/v1/email"
}

func ProxyURL() string {
	return "http://localhost:8888"
}
