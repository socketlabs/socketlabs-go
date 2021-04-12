package core

import (
	"net/http"

	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
)

type RetryHandler interface {
	Send(request *http.Request) (*http.Response, error)
}

type retryHandler struct {
	HttpClient *http.Client
	EndpointUrl string
	Settings injectionapi.RetrySettings
}

func CreateRetryHandler(client *http.Client, endpointUrl string, settings injectionapi.RetrySettings) RetryHandler {
	return &retryHandler{
		HttpClient: client,
		EndpointUrl: endpointUrl,
		Settings: settings,
	}
}

func (retryHandler *retryHandler) Send(request *http.Request) (*http.Response, error) {
	if retryHandler.Settings.GetMaximumNumberOfRetries() == 0 {
		response, err := retryHandler.HttpClient.Do(request)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
	//attempts := 0
	//for true {
	//	waitInterval := retryHandler.Settings.GetNextWaitInterval(attempts)
	//
	//}
	return nil, nil
}