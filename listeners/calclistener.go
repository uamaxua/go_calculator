package listeners

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go_calculator/calculator"
	"log"
)

const CalcSubject = "calc"

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func ListenNatsCalcRequests() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, subErr := nc.Subscribe(CalcSubject, func(m *nats.Msg) {
		log.Printf("Received a request: %s", string(m.Data))
		var request Request
		jsonErr := json.Unmarshal(m.Data, &request)
		if jsonErr != nil {
			log.Printf("Error decoding JSON: %v", err)
			_ = m.Respond([]byte(`{"error": "Invalid JSON format"}`))
			return
		}
		result, jsonErr := calculator.Calculate(request.Expression)
		res, _ := json.Marshal(Response{Result: result})
		err := m.Respond(res)
		if err != nil {
			log.Printf("Error responding to message: %v", err)
		}
	})
	if subErr != nil {
		log.Printf("Error subscribing to subject 'hello': %v", err)
		return
	}
	log.Printf("Listening for requests on subject: '%s'...", CalcSubject)
	select {}
}
