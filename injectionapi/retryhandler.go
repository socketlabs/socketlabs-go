package injectionapi

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

var ErrorCodes = [4]int{500, 502, 503, 504}

type RetryHandler interface {
	Send(request *http.Request) (*http.Response, error)
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

func (retryHandler *retryHandler) Send(request *http.Request) (*http.Response, error) {
	if retryHandler.Settings.GetMaximumNumberOfRetries() == 0 {
		response, err := retryHandler.HttpClient.Do(request)
		return response, err
	}
	attempts := 0
	for true {
		waitInterval := retryHandler.Settings.GetNextWaitInterval(attempts)
		response, err := retryHandler.HttpClient.Do(request)
		if err == nil {

			if elementInArray(ErrorCodes, response.StatusCode) {
				attempts++
				fmt.Println("Retry : ", attempts)
				if attempts > retryHandler.Settings.GetMaximumNumberOfRetries() {
					return response, errors.New("Received Http Status Code : " + string(response.StatusCode))
				}
				time.Sleep(waitInterval)
			} else {
				return response, err
			}

		} else {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				attempts++
				fmt.Println("Retry : ", attempts)
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

func elementInArray(array [4]int, element int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}
	return false
}