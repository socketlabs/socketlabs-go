package injectionapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type injectionResponseParser struct {
}

func (parser injectionResponseParser) Parse(resp *http.Response) (SendResponse, error) {
	//read body into byte array
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return SendResponse{}, err
	}

	//deserialize
	var injectionResponse injectionResponseDto
	err = json.Unmarshal(body, &injectionResponse)
	if err != nil {
		return SendResponse{}, err
	}

	var sendResult = determineSendResult(injectionResponse.ErrorCode, resp.StatusCode)
	newResponse := SendResponse{
		Result:             sendResult,
		TransactionReceipt: injectionResponse.TransactionReceipt,
		ResponseMessage:    sendResult.ToResponseMessage(),
	}

	if sendResult == SendResultWARNING {
		if injectionResponse.MessageResults != nil && len(injectionResponse.MessageResults) > 0 {
			errorCode := injectionResponse.MessageResults[0].ErrorCode
			newResponse.Result = SendResultSUCCESS.Parse(errorCode)
		}
	}

	if injectionResponse.MessageResults != nil && len(injectionResponse.MessageResults) > 0 {
		newResponse.AddressResults = injectionResponse.MessageResults[0].AddressResults
	}

	return newResponse, nil
}

func determineSendResult(errorCode string, statusCode int) SendResult {
	switch statusCode {
	case 200:
		sendResult := SendResultSUCCESS.Parse(errorCode)
		if sendResult == -1 {
			return SendResultUNKNOWNERROR
		}
		return sendResult
	case 500:
		return SendResultINTERNALERROR
	case 408:
		return SendResultTIMEOUT
	case 401:
		return SendResultINVALIDAUTHENTICATION
	default:
		return SendResultUNKNOWNERROR
	}

	return SendResultSUCCESS
}
