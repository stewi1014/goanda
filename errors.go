package goanda

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func newAPIError(responceCode int, body io.ReadCloser) APIError {
	defer body.Close()

	msg := struct {
		errorMessage string `json:"errorMessage,string"`
		rejectReason string `json:"rejectReason,string"`
	}{}

	b, _ := ioutil.ReadAll(body)

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return APIError{
			HTTPResponceCode: responceCode,
			Message:          string(b),
		}
	}
	return APIError{
		HTTPResponceCode: responceCode,
		Message:          msg.errorMessage + msg.rejectReason,
	}
}

// APIError is returned when the Oanda server responds with an error
//
// Message is the returned error message from the server if possible to unmarshal,
// otherwise it is simply the entire body of the responce
type APIError struct {
	HTTPResponceCode int
	Message          string
}

// APIError implements error
func (a APIError) Error() string {
	return fmt.Sprintf("Oanda API Error %v: %v",
		http.StatusText(a.HTTPResponceCode),
		a.Message,
	)
}
