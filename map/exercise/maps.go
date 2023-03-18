package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

func ServerInfo(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")
	stats := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1

		case Offline:
			stats[Offline] += 1

		case Maintenance:
			stats[Maintenance] += 1

		case Retired:
			stats[Retired] += 1

		default:
			panic("unhandled server status")
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{
		"darkstar",
		"aiur",
		"omicron",
		"w359",
		"baseline",
	}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	ServerInfo(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	ServerInfo(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}

	ServerInfo(serverStatus)
}
