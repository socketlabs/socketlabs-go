package injectionapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"runtime"

	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/core"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/core/serialization"
	"github.com/PraneethChandraThota/socketlabs-go/injectionapi/message"
)

const endpointURL = "https://inject.socketlabs.com/api/v1/email"

// ISocketlabsClient is used to easily send messages through the Socketlabs Injection API
type ISocketlabsClient interface {

	// SendBasic sends a basic email message and returns the response from the Injection API.
	SendBasic(message *message.BasicMessage) (SendResponse, error)

	// SendBulk sends a bulk email message and returns the response from the Injection API.
	SendBulk(message *message.BulkMessage) (SendResponse, error)

	// SetEndpointURL sets the API endpoint. Under normal circumstances this should not be used.
	SetEndpointURL(endpointURL string)

	// SetProxyUrl sets the proxy url.
	SetProxyURL(proxyURL string)

	// GetEndpointURL retreives the API endpoint.
	GetEndpointURL() string
}

// socketlabsClient is the default ISocketlabsClient implementation
type socketlabsClient struct {
	ServerID    int
	APIKey      string
	EndpointURL string
	ProxyURL    string
}

// CreateClient instatiates new client using the specified credentials
func CreateClient(serverID int, apiKey string) ISocketlabsClient {
	return &socketlabsClient{
		ServerID:    serverID,
		APIKey:      apiKey,
		EndpointURL: endpointURL,
	}
}

// CreateClientWithProxy instatiates new client using the specified credentials
func CreateClientWithProxy(serverID int, apiKey string, proxyURL string) ISocketlabsClient {
	return &socketlabsClient{
		ServerID:    serverID,
		APIKey:      apiKey,
		EndpointURL: endpointURL,
		ProxyURL:    proxyURL,
	}
}

// SetEndpointURL sets the API endpoint. Under normal circumstances this should not be used.
func (socketlabsClient *socketlabsClient) SetEndpointURL(endpointURL string) {
	socketlabsClient.EndpointURL = endpointURL
}

// SetProxyUrl sets the proxy url.
func (socketlabsClient *socketlabsClient) SetProxyURL(proxyURL string) {
	socketlabsClient.ProxyURL = proxyURL
}

// GetEndpointURL retreives the API endpoint.
func (socketlabsClient *socketlabsClient) GetEndpointURL() string {
	return socketlabsClient.EndpointURL
}

// SendBasic sends a basic email message and returns the response from the Injection API.
func (socketlabsClient *socketlabsClient) SendBasic(message *message.BasicMessage) (SendResponse, error) {

	sendResponse := SendResponse{
		Result: SendResultUNKNOWNERROR,
	}

	validator := sendValidator{}

	sendResponse = validator.ValidateCredentials(socketlabsClient.ServerID, socketlabsClient.APIKey)
	if sendResponse.Result != SendResultSUCCESS {
		sendResponse.ResponseMessage = sendResponse.Result.ToResponseMessage()
		return sendResponse, nil
	}

	sendResponse = validator.ValidateBasicMessage(*message)
	if sendResponse.Result != SendResultSUCCESS {
		sendResponse.ResponseMessage = sendResponse.Result.ToResponseMessage()
		return sendResponse, nil
	}

	//create injection request from factory
	factory := core.InjectionRequestFactory{
		ServerID: socketlabsClient.ServerID,
		APIKey:   socketlabsClient.APIKey,
	}
	request := factory.GenerateBasicRequest(message)

	return socketlabsClient.sendInjectionRequest(&request)
}

// SendBulk sends a bulk email message and returns the response from the Injection API.
func (socketlabsClient *socketlabsClient) SendBulk(message *message.BulkMessage) (SendResponse, error) {

	sendResponse := SendResponse{
		Result: SendResultUNKNOWNERROR,
	}

	validator := sendValidator{}

	sendResponse = validator.ValidateCredentials(socketlabsClient.ServerID, socketlabsClient.APIKey)
	if sendResponse.Result != SendResultSUCCESS {
		sendResponse.ResponseMessage = sendResponse.Result.ToResponseMessage()
		return sendResponse, nil
	}

	sendResponse = validator.ValidateBulkMessage(*message)
	if sendResponse.Result != SendResultSUCCESS {
		sendResponse.ResponseMessage = sendResponse.Result.ToResponseMessage()
		return sendResponse, nil
	}

	//create injection request from factory
	factory := core.InjectionRequestFactory{
		ServerID: socketlabsClient.ServerID,
		APIKey:   socketlabsClient.APIKey,
	}
	request := factory.GenerateBulkRequest(message)

	return socketlabsClient.sendInjectionRequest(&request)
}

func (socketlabsClient socketlabsClient) sendInjectionRequest(injectionRequest *serialization.InjectionRequestJson) (sendResponse SendResponse, err error) {

	//attempt to serialize injection request
	serializedRequest, err := json.Marshal(injectionRequest)
	if err != nil {
		return SendResponse{}, err
	}

	//create request
	req, err := http.NewRequest("POST", socketlabsClient.GetEndpointURL(), bytes.NewBuffer(serializedRequest))
	if err != nil {
		return SendResponse{}, err
	}

	//add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "socketlabs-go/1.0.1 ("+runtime.Version()+")")

	//create http client
	client, err := socketlabsClient.createHttpClient(socketlabsClient.ProxyURL)
	if err != nil {
		return SendResponse{}, err
	}

	//issue http request
	resp, err := client.Do(req)
	if err != nil {
		return SendResponse{}, err
	}

	//map to response and return
	return injectionResponseParser{}.Parse(resp)
}

func (socketlabsClient *socketlabsClient) createHttpClient(proxyUrl string) (*http.Client, error) {

	//create http client
	if socketlabsClient.ProxyURL == "" {
		return &http.Client{}, nil
	}

	//attempt to parse proxy url
	var proxyURL *url.URL
	proxyURL, err := url.Parse(socketlabsClient.ProxyURL)
	if err != nil {
		return nil, err
	}

	//create client with proxy url
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	return client, nil
}
