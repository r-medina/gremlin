package gremlin

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	RequestId string          `json:"requestId"`
	Status    *ResponseStatus `json:"status"`
	Result    *ResponseResult `json:"result"`
}

type ResponseStatus struct {
	Code       int                    `json:"code"`
	Attributes map[string]interface{} `json:"attributes"`
	Message    string                 `json:"message"`
}

type ResponseResult struct {
	Data json.RawMessage        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

// Implementation of the stringer interface. Useful for exploration
func (r Response) String() string {
	return fmt.Sprintf("Response \nRequestId: %v, \nStatus: {%#v}, \nResult: {%#v}\n", r.RequestId, r.Status, r.Result)
}

// ResponseError represents the information returned from a Gremlin server in
// the event of an error.
type ResponseError struct {
	Code    int
	Name    string
	Message string
}

func newResponseError(code int, msg string) error {
	name, ok := ErrorMsg[code]
	if !ok {
		name = "Unkown Error"
	}

	return &ResponseError{
		Code:    code,
		Name:    name,
		Message: msg,
	}
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("%s: %d: %s", err.Name, err.Code, err.Message)
}
