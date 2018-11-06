package injectionapi

type MessageResultDto struct {
	Index          int
	ErrorCode      string
	AddressResults []AddressResult
}
