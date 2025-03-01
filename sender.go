package discordrpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

// Sends data to Discord in proper format.
func sendOperation(conn net.Conn, op uint32, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	var header bytes.Buffer

	if err := binary.Write(&header, binary.LittleEndian, op); err != nil {
		return err
	}
	if err := binary.Write(&header, binary.LittleEndian, uint32(len(data))); err != nil {
		return err
	}

	if _, err := conn.Write(append(header.Bytes(), data...)); err != nil {
		return err
	}

	return nil
}

// Function to match arg type on official Discord docs.
func (c *Client) SetActivity(activityData *ActivityData) error {
	// Check for connection being nil
	if !isDiscordRunning() {
		if c.conn != nil {
			c.conn.Close()
		}
		c.conn = nil
		return errors.New("Discord closed.")
	} else {
		err := c.reconnect()

		if err != nil {
			return err
		}
	}

	c.sendActivity(activityData)

	return nil
}

// Send updated activity data.
func (c *Client) sendActivity(activityData *ActivityData) error {
	// Pad fields below with a space character, when setting these
	// to a string length of one it fails to display in discord.
	if len(activityData.Details) == 1 {
		activityData.Details += " "
	}

	if len(activityData.State) == 1 {
		activityData.State += " "
	}

	if len(activityData.Assets.LargeText) == 1 {
		activityData.Assets.LargeText += " "
	}

	if len(activityData.Assets.SmallText) == 1 {
		activityData.Assets.SmallText += " "
	}

	activity := internalActivity{
		Cmd: "SET_ACTIVITY",
		Args: internalArgs{
			Pid:      os.Getpid(),
			Activity: activityData,
		},
		Nonce: fmt.Sprintf("clear_activity_%d", time.Now().UnixNano()),
	}

	err := sendOperation(c.conn, 1, activity)

	if err != nil {
		return err
	}

	return nil
}

// Receive response back from handshake.
func receiveResponse(conn net.Conn) (uint32, map[string]interface{}, error) {
	header := make([]byte, 8)
	_, err := conn.Read(header)
	if err != nil {
		return 0, nil, err
	}

	op := binary.LittleEndian.Uint32(header[:4])
	length := binary.LittleEndian.Uint32(header[4:])

	data := make([]byte, length)
	_, err = conn.Read(data)
	if err != nil {
		return 0, nil, err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return 0, nil, err
	}

	return op, payload, nil
}

// Initialize the connection.
func sendHandshake(conn net.Conn, appId string) error {
	msg := handshake{
		V:        1,
		ClientID: appId,
	}

	err := sendOperation(conn, 0, msg)

	if err != nil {
		return err
	}

	return nil
}
