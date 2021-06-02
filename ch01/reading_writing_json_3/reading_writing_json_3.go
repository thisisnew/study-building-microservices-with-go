package reading_writing_json_3

import (
	"encoding/json"
	"net/http"
	"study-building-microservices-with-go/ch01/reading_writing_json_2"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := reading_writing_json_2.HelloWorldResponse{
		Message: "HelloWorld",
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
