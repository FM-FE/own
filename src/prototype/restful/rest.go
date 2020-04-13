package restful

import "net/http"

type CommonResponse struct {
	Result string `json:"result"`
	Error  string `json:"errors"`
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	
}

func Hello(w http.ResponseWriter, r *http.Request) {

}
