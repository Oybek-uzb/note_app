package rest

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIResponse struct {
	IsOK     bool
	Response *http.Response
	Error    APIError
}

func (ar *APIResponse) Body() io.ReadCloser {
	return ar.Response.Body
}

func (ar *APIResponse) ReadBody() ([]byte, error) {
	defer ar.Response.Body.Close()
	return ioutil.ReadAll(ar.Response.Body)
}

func (ar *APIResponse) StatusCode() int {
	return ar.Response.StatusCode
}

func (ar *APIResponse) Location() (*url.URL, error) {
	return ar.Response.Location()
}

type APIError struct {
	Message          string `json:"message,omitempty"`
	ErrorCode        string `json:"error_code,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
}

func (aep *APIError) ToString() string {
	return fmt.Sprintf("Error Code: %s, Error: %s, Developer Error: %s", aep.ErrorCode, aep.Message, aep.DeveloperMessage)
}
