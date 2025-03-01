package discordrpc

import (
	"log"

	"github.com/shirou/gopsutil/v4/process"
)

// Checks if Discord is actively running.
func isDiscordRunning() bool {
	processes, err := process.Processes()
	if err != nil {
		log.Println("Error fetching processes:", err)
		return false
	}

	for _, p := range processes {
		name, err := p.Name()
		if err == nil && (name == "Discord" || name == "Discord.exe") {
			return true
		}
	}

	return false
}
