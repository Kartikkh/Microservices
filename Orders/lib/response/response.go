package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message     string      `json:"message,omitempty"`
	ErrorString string      `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	StatusCode  int         `json:"-"`
	Error       error       `json:"-"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{StatusCode: http.StatusOK}
}

func (r *JSONResponse) SetStatusCode(code int) *JSONResponse {
	r.StatusCode = code
	return r
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func (r *JSONResponse) SetError(err error, a ...string) *JSONResponse {
	r.Error = err
	if err != nil {
		r.ErrorString = err.Error()
	}
	return r
}

func (r *JSONResponse) Send(w http.ResponseWriter) {
	b, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	w.Write(b)
}
