package xml

// Response XML representation
type Response struct {
	Failure *IntegrityError
	Success *Success
}

// NewResponse with success, failuer attributes
func NewResponse(succsess *Success, failure *IntegrityError) *Response {
	response := new(Response)
	if failure.Body != nil || failure.Description != "" {
		response.Failure = failure
	} else {
		response.Success = succsess
	}

	return response
}

// Success response
type Success struct {
	UmmsID           string `xml:"success>ummsId"`
	ExternalSystemID string `xml:"success>externalSystemId"`
	ExternalCaseID   string `xml:"success>externalCaseId"`
	SchemaVersion    string `xml:"schemaVersion,attr"`
	Xmlns            string `xml:"xmlns,attr"`
	Ns2              string `xml:"ns2,attr"`
	EntityType       string `xml:"entityType"`
}

// IntegrityError representation for Message
type IntegrityError struct {
	Xmlns         string `xml:"xmlns,attr"`
	Description   string `xml:"description"`
	Body          []byte `xml:"body"`
	SchemaVersion string `xml:"schemaVersion,attr"`
}

// IFailure interface for error-response
// type IFailure interface {
// 	GetBody() []byte
// 	GetDescription() string
// }

// GetBody for IntegrityError : IFailure
// func (err IntegrityError) GetBody() []byte {
// 	return err.Body
// }

// GetDescription for IntegrityError : IFailure
// func (err IntegrityError) GetDescription() string {
// 	return err.Description
// }
