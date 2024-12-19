package listeners

import (
	"github.com/nats-io/nats.go"
	"log"
)

func ListenNatsCalcRequests() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, subErr := nc.Subscribe("hello", func(m *nats.Msg) {
		log.Printf("Received a request: %s\n", string(m.Data))
		err := m.Respond([]byte("Hello, World!"))
		if err != nil {
			log.Printf("Error responding to message: %v", err)
		}
	})
	if subErr != nil {
		log.Printf("Error subscribing to subject 'hello': %v", err)
		return
	}

	log.Println("Listening for requests on 'hello'...")
	select {}
}
