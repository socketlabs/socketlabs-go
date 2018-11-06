package injectionapi

type InjectionResponseDto struct {
	ErrorCode          string
	TransactionReceipt string
	MessageResults     []MessageResultDto
}
