package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

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

func NewClient(appId string) (net.Conn, error) {
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

	return conn, nil
}

func CloseClient() {

}

func main() {
	conn, err := NewClient("1332158263432708146")

	data := ActivityData{
		State:      "t",
		Type:       2,
		Details:    "D",
		Timestamps: ActivityTimestamp{
			// Start: int(time.Now().Unix()),
			// Start: 1,
		},
		Assets: ActivityAssets{
			LargeText:  "M",
			LargeImage: "https://i.ytimg.com/vi/jlRmgQ7VXgA/sddefault.jpg",
			SmallText:  "M",
			SmallImage: "https://i.ytimg.com/vi/jlRmgQ7VXgA/sddefault.jpg",
		},
	}

	err = SendActivity(conn, data)
	if err != nil {
		panic(fmt.Errorf("Failed to send activity: %v", err))
	}

	// keep process idle
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	defer conn.Close()
}
