# Go Calculator

This is a simple expression calculator implemented in Go. <br>
Supported operators: '+', '-', '*', '/', '(', ')'.

It provides two modes of operation:

1. A basic HTTP server to evaluate mathematical expressions.
2. The same functionality implemented via NATS messaging.

---

## Prerequisites

### Software Required

1. **Go**: Install from [https://go.dev/dl/](https://go.dev/dl/).
2. **Docker**: Install from [https://www.docker.com/get-started/](https://www.docker.com/get-started/).
3. **NATS-CLI**: Install from [https://github.com/nats-io/natscli/](https://github.com/nats-io/natscli/).

---

## How to Run

### HTTP Calculator

1. Run NATS server with docker:
   ```bash
    docker run --rm -d -p 4222:4222 --name nats-server nats:latest

2. Clone this repository:
   ```bash
   git clone https://github.com/your_username/go_calculator.git
   cd go_calculator

3. Download go dependencies:
    ```bash
    go mod download

4. Start main application:
    ```bash
    go run main.go

5. Test HTTP request:
   #### - Send a GET request:
        curl "http://localhost:8080/calc?expression=(2%2B68)*10"

   #### - Response:
        Result: 700.00
6. Test NATS request:
   #### - Subscribe to reply subject (do not close terminal):
        nats sub calc_result

   #### - Open terminal in another window and send request:
    ```bash
      nats pub calc '{\"expression\": \"2+2\"}' --reply calc_result
   
  #### - Open terminal with subscription and check response:
   ```bash
   Received on "calc_result"
   {"result":4}
