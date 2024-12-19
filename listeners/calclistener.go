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
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	_, err = nc.Subscribe(CalcSubject, func(m *nats.Msg) {
		handleCalcRequest(m)
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to subject '%s': %v", CalcSubject, err)
	}
	log.Printf("Listening for requests on subject: '%s'...", CalcSubject)
	select {}
}

func handleCalcRequest(m *nats.Msg) {
	log.Printf("Received request: %s", string(m.Data))
	var request Request
	if err := json.Unmarshal(m.Data, &request); err != nil {
		sendErrorResponse(m, "Invalid JSON format")
		log.Printf("Error parsing JSON: %v", err)
		return
	}
	result, err := calculator.Calculate(request.Expression)
	if err != nil {
		sendErrorResponse(m, err.Error())
		log.Printf("Error calculating expression: %v", err)
		return
	}
	sendSuccessResponse(m, result)
}

func sendErrorResponse(m *nats.Msg, errorMsg string) {
	response := Response{Error: errorMsg}
	sendResponse(m, response)
}

func sendSuccessResponse(m *nats.Msg, result float64) {
	response := Response{Result: result}
	sendResponse(m, response)
}

func sendResponse(m *nats.Msg, response Response) {
	resBytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to serialize response: %v", err)
		return
	}
	if err := m.Respond(resBytes); err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}
