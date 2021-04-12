package injectionapi

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

var ErrorCodes = []int{500, 502, 503, 504}

type RetryHandler interface {
	Send(serializedRequest []byte) (*http.Response, error)
}

type retryHandler struct {
	HttpClient *http.Client
	EndpointUrl string
	Settings RetrySettings
}

func CreateRetryHandler(client *http.Client, endpointUrl string, settings RetrySettings) RetryHandler {
	return &retryHandler{
		HttpClient: client,
		EndpointUrl: endpointUrl,
		Settings: settings,
	}
}

func (retryHandler *retryHandler) Send(serializedRequest []byte) (*http.Response, error) {

	if retryHandler.Settings.GetMaximumNumberOfRetries() == 0 {
		request, err := createRequest(retryHandler.EndpointUrl, serializedRequest)
		if err != nil {
			return nil, err
		}
		response, err := retryHandler.HttpClient.Do(request)
		return response, err
	}
	attempts := 0
	for true {

		request, err := createRequest(retryHandler.EndpointUrl, serializedRequest)
		if err != nil {
			return nil, err
		}

		waitInterval := retryHandler.Settings.GetNextWaitInterval(attempts)

		response, err := retryHandler.HttpClient.Do(request)

		if err == nil {

			if elementInArray(ErrorCodes, response.StatusCode) {
				attempts++
				fmt.Println("Retry : ", attempts, " Status Code : ", response.StatusCode)
				if attempts > retryHandler.Settings.GetMaximumNumberOfRetries() {
					return response, errors.New(fmt.Sprintf("Received Http Status Code : %d", response.StatusCode))
				}

				time.Sleep(waitInterval)
			} else {

				return response, err
			}
		} else {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				attempts++
				fmt.Println("Retry : ", attempts, " Timeout Error : ", err)
				if attempts > retryHandler.Settings.GetMaximumNumberOfRetries() {
					return response, err
				}
				time.Sleep(waitInterval)
			} else {
				return response, err
			}
		}

	}

	return nil, nil
}

func elementInArray(array []int, element int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}
	return false
}

func createRequest(endpointUrl string, serializedRequest []byte) (*http.Request, error) {

	//create request
	request, err := http.NewRequest("POST", endpointUrl, bytes.NewBuffer(serializedRequest))
	if err != nil {
		return nil, err
	}

	//add headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "socketlabs-go/1.0.1 ("+runtime.Version()+")")

	return request, nil

}