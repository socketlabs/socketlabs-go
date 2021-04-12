package core

import (
	"net/http"

	"github.com/PraneethChandraThota/socketlabs-go/injectionapi"
)

type RetryHandler interface {
	Send(client *http.Client, request *http.Request) (injectionapi.SendResponse, error)
}

