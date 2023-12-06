# Chat Application

## Description

This project is a simple chat application that enables users to connect via WebSockets and communicate in real-time. The server is written in Go, leveraging the `gorilla/websocket` package. The client features a Text-based User Interface (TUI) built with the `charm/bubbletea` library, offering an interactive and user-friendly way to engage in chat.

## Installation

### Prerequisites

- Go (1.21.5 or later)

### Setting Up the Server

1. Clone the repository:

   ```bash
   git clone https://github.com/Askath/Go-ChatAndClient.git
   ```

   ```

   ```

2. Run the server:
   ```bash
   go run server.go
   ```

### Setting Up the Client

1. Navigate to the client directory:
   ```bash
   cd  client
   ```
2. Run the client with TUI:
   ```bash
   go run client.go
   ```

## Usage

After starting the server and client:

- Interact with the client's TUI to send and receive messages in real-time.
- To exit the client, press `q` or `Ctrl+C`.

## Features

- Real-time messaging via WebSockets
- Interactive Text-based User Interface (TUI) for clients

## License

This project is licensed under the [MIT License](LICENSE).
