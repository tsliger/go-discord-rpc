// Apple Darwin specific socket implementation
package discordrpc

import (
	"fmt"
	"net"
	"path/filepath"
)

// Get socket from var directory.
func getDiscordSocket() (string, error) {
	pattern := "/var/folders/*/*/*/discord-ipc-0"
	matches, err := filepath.Glob(pattern)

	if err != nil {
		return "", err
	}

	if len(matches) > 0 {
		return matches[0], nil
	}

	return "", nil // Return empty string if no match is found
}

// Form connection to the socket.
func connectToSocket(sock string) (net.Conn, error) {
	conn, err := net.Dial("unix", sock)

	if err != nil {
		return nil, fmt.Errorf("Issue connecting to the socket: %w", err) // Wrap the error for better debugging
	}

	return conn, nil
}
