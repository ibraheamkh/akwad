package http

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success  bool              `json:"success"`
	Result   interface{}       `json:"result"`
	Errors   []ResponseMessage `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
}

func NewSuccessResponseWithMessage(result interface{}, message string, code int) Response {
	return Response{
		Success:  true,
		Result:   result,
		Errors:   []ResponseMessage{},
		Messages: []ResponseMessage{{code, message}},
	}
}

func NewSuccessResponse(result interface{}) Response {
	return Response{
		Success:  true,
		Result:   result,
		Errors:   []ResponseMessage{},
		Messages: []ResponseMessage{},
	}
}

func NewErrorResponse(message string, code int) Response {
	return Response{
		Success:  false,
		Result:   nil,
		Errors:   []ResponseMessage{{code, message}},
		Messages: []ResponseMessage{},
	}
}

func SendResponse(w http.ResponseWriter, result interface{}) error {
	response := NewSuccessResponse(result)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(response)
	return err
}

func SendResponseWithMessage(w http.ResponseWriter, result interface{}, message string, code int) error {
	response := NewSuccessResponseWithMessage(result, message, code)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(response)
	return err
}

func SendErrorResponse(w http.ResponseWriter, message string, code int) error {
	response := NewErrorResponse(message, code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	enc := json.NewEncoder(w)
	err := enc.Encode(response)
	return err
}
