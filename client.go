package discordrpc

import (
	"fmt"
	"net"
)

var conn net.Conn

// Create new client with appID from Discord.
func NewClient(appId string) error {
	socket, err := getDiscordSocket()

	if err != nil {
		return fmt.Errorf("Failed to retrieve socket information: %v", err)
	}

	conn, err = connectToSocket(socket)
	if err != nil {
		return fmt.Errorf("Failed to connect to socket: %v", err)
	}

	err = sendHandshake(conn, appId)
	if err != nil {
		return fmt.Errorf("Failed to send initial handshake: %v", err)
	}

	op, payload, err := receiveResponse(conn)
	if err != nil {
		return fmt.Errorf("Response not receieved from initial handshake: %v", err)
	}
	fmt.Println(op, payload)

	return nil
}

// Close socket connection.
func CloseClient() {
	conn.Close()
}
