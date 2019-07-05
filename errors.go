package goanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func newAPIError(request *http.Request, responce *http.Response) APIError {
	defer responce.Body.Close()

	msg := struct {
		ErrorMessage string
		RejectReason string
	}{}

	apiErr := APIError{
		Responce: responce,
		Request:  request,
	}

	b, _ := ioutil.ReadAll(responce.Body)
	err := json.Unmarshal(b, &msg)
	if err != nil {
		apiErr.Message = string(b)
	} else {
		apiErr.Message = msg.RejectReason + msg.ErrorMessage
	}
	return apiErr
}

// APIError is returned when the Oanda server responds with an error
//
// Message is the returned error message from the server if possible to unmarshal,
// otherwise it is simply the entire body of the responce
type APIError struct {
	Request  *http.Request
	Responce *http.Response
	Message  string
}

// APIError implements error
func (a APIError) Error() string {
	return fmt.Sprintf("Oanda API Error [Url: %v, Responce: %v]: %v",
		a.Request.URL.String(),
		a.Responce.Status,
		a.Message,
	)
}
