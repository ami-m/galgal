package apis

type Api interface {
	// SendRequest Resposnible for sending the request and return the response
	SendRequest() (interface{}, error)
	// GenerateResponse Responsible to generate the response to the caller with the correct convention
	GenerateResponse() (interface{}, error)
}
