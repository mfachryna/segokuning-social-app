package response

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

type Response struct {
	HttpStatus int         `json:"-"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       Meta        `json:"meta"`
}

var (
	ValidationAndParseBodyError = "Required fields are missing or invalid"
)

func (res *Response) GenerateResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.HttpStatus)
	json.NewEncoder(w).Encode(res)
}
