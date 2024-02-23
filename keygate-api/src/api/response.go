package api


type JSendResponse struct {
	Status string `json:"status"`
	Data interface{} `json:"data,omitempty"`
	Message string `json:"message"`
}

const (
	ResponseSuccess string = "success"
	ResponseFail string = "fail"
	ResponseError string = "error"
)