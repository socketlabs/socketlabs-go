package injectionapi

type injectionResponseDto struct {
	ErrorCode          string
	TransactionReceipt string
	MessageResults     []messageResultDto
}
