package reading_writing_json_5

import (
	"encoding/json"
	"net/http"
)

type helloWorldRequest struct {
	name string `json:"name"`
}

type helloWorldResponse struct {
	Message string
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
