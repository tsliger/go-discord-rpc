package discordrpc

import (
	"fmt"
	"net"
)

type Client struct {
	id   string
	conn net.Conn
}

func NewClient(appId string) (*Client, error) {
	socket, err := getDiscordSocket()

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve socket information: %v", err)
	}

	conn, err := connectToSocket(socket)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to socket: %v", err)
	}

	err = sendHandshake(conn, appId)
	if err != nil {
		return nil, fmt.Errorf("Failed to send initial handshake: %v", err)
	}

	op, payload, err := receiveResponse(conn)
	if err != nil {
		return nil, fmt.Errorf("Response not receieved from initial handshake: %v", err)
	}
	fmt.Println(op, payload)

	return &Client{id: appId, conn: conn}, nil
}

func (c *Client) CloseClient() error {
	if c.conn != nil {
		err := c.conn.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
