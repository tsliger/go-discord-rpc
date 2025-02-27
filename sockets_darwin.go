package main

import (
	"fmt"
	"net"
	"path/filepath"
)

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

func connectToSocket(sock string) (net.Conn, error) {
	conn, err := net.Dial("unix", sock)

	if err != nil {
		return nil, fmt.Errorf("Issue connecting to the socket: %w", err) // Wrap the error for better debugging
	}

	return conn, nil
}
