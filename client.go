package discordrpc

import (
	"fmt"
	"net"
)

// Creates new client
func NewClient(appId string) (*Client, error) {
	appClient := &Client{id: appId, conn: nil}
	conn, err := createConnection()

	if err != nil {
		return appClient, err
	}

	err = initializeClient(conn, appId)
	if err != nil {
		return appClient, err
	}

	appClient.conn = conn
	return appClient, nil
}

// Creates connection to socket.
func createConnection() (net.Conn, error) {
	socket, err := getDiscordSocket()

	if err != nil {
		return nil, err
	}

	conn, err := connectToSocket(socket)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to socket: %v", err)
	}

	return conn, nil
}

// Initializes client via handshake.
func initializeClient(conn net.Conn, appId string) error {
	err := sendHandshake(conn, appId)
	if err != nil {
		return fmt.Errorf("Failed to send initial handshake: %v", err)
	}

	_, _, err = receiveResponse(conn)
	if err != nil {
		return fmt.Errorf("Response not receieved from initial handshake: %v", err)
	}

	return nil
}

// Reconnects client to socket when Discord closes and reopens.
func (c *Client) reconnect() error {
	conn, err := createConnection()

	if err != nil {
		return err
	}

	err = initializeClient(conn, c.id)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

// Closes client connection.
func (c *Client) CloseClient() error {
	if c.conn != nil {
		err := c.conn.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
