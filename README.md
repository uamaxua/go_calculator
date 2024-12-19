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

5. Test HTTP server:
   #### - Send a GET request:
        curl "http://localhost:8080/calc?expression=(2%2B68)*10"

   #### - Response:
        Result: 700.00