// Apple Darwin specific socket implementation
package discordrpc

import (
	"fmt"
	"net"
	"github.com/Microsoft/go-winio"
)

const pipeName = `\\.\pipe\discord-ipc-0`

// Get Discord socket.
func getDiscordSocket() (string, error) {
	return pipeName, nil
}

// Connects to Discord socket.
func connectToSocket(sock string) (net.Conn, error) {
	conn, err := winio.DialPipe(sock, nil)
	if err != nil {
		return nil, fmt.Errorf("Issue connecting to the pipe: %w", err)
	}
	return conn, nil
}