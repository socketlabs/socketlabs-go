package injectionapi

type messageResultDto struct {
	Index          int
	ErrorCode      string
	AddressResults []AddressResult
}
