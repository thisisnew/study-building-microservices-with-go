package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "HelloWorld",
	}

	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Add("Content-Type", "application/javascript")
		result := fmt.Sprintf("%s(%s)", callback, string(data))
		fmt.Fprintf(w, result)
	} else {
		fmt.Fprintf(w, string(data))
	}
}
