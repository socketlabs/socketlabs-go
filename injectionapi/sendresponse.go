package injectionapi

//SendResponse The response of an SocketLabsClient send request.
type SendResponse struct {
	Result             SendResult
	TransactionReceipt string
	ResponseMessage    string
	AddressResults     []AddressResult
}
