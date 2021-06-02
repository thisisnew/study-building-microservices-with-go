package reading_writing_json_4

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type helloWorldRequest struct {
	name string `json:"name"`
}

type helloWorldResponse struct {
	Message string
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
