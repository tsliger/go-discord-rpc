package discordrpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func sendOperation(conn net.Conn, op uint32, payload interface{}) error {
	if payload == nil {
		payload = map[string]interface{}{}
	}

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

// func (c *Client) SendActivity(data ActivityData) error {
func (c *Client) SendActivity(data interface{}) error {
	// Pad fields below with a space character, when setting these
	// to a string length of one it fails to display in discord.
	if data == nil {
		err := sendOperation(c.conn, 1, nil) // `nil` will send an empty payload
		if err != nil {
			return err
		}

		return nil
	}

	if activityData, ok := data.(ActivityData); ok {
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
			Nonce: "1234",
		}

		err := sendOperation(c.conn, 1, activity)

		if err != nil {
			fmt.Println("Failed to send operation")
			return err
		}
	}

	return nil
}

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
