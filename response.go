package response

import (
	"net/http"
	"os"
)

type ApiResponseStatus struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

type ApiResponse struct {
	ProcessId string            `json:"process_id"`
	Path      string            `json:"path"`
	Status    ApiResponseStatus `json:"status"`
	Request   interface{}       `json:"request"`
	Errors    []error           `json:"errors"`
	Data      interface{}       `json:"data"`
}

func NewApiResponse(path string) *ApiResponse {
	resp_status := ApiResponseStatus{
		Code: http.StatusOK,
		Type: http.StatusText(http.StatusOK),
	}
	resp_error := []error{}
	resp_data := make([]interface{}, 0)

	return &ApiResponse{
		ProcessId: os.Getenv("PROCESS_ID"),
		Status:    resp_status,
		Errors:    resp_error,
		Data:      resp_data,
		Path:      path,
	}
}
