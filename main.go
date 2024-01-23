package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/miscord-dev/palog/pkg/palrcon"
)

var (
	rconEndpoint = os.Getenv("RCON_ENDPOINT")
	rconPassword = os.Getenv("RCON_PASSWORD")

	intervalRaw = os.Getenv("INTERVAL")
	interval    time.Duration

	timeoutRaw = os.Getenv("TIMEOUT")
	timeout    time.Duration
)

func init() {
	var err error

	if timeoutRaw == "" {
		timeoutRaw = "1s"
	}

	timeout, err = time.ParseDuration(timeoutRaw)
	if err != nil {
		slog.Error("failed to parse timeout", "error", err)
		os.Exit(1)
	}

	if intervalRaw == "" {
		intervalRaw = "5s"
	}

	interval, err = time.ParseDuration(intervalRaw)

	if err != nil {
		slog.Error("failed to parse interval", "error", err)
		os.Exit(1)
	}
}

func main() {
	palRCON := palrcon.NewPalRCON(rconEndpoint, rconPassword)
	palRCON.SetTimeout(timeout)

	var prev map[string]palrcon.Player

	makeMap := func(players []palrcon.Player) map[string]palrcon.Player {
		m := make(map[string]palrcon.Player)

		for _, player := range players {
			m[player.SteamID] = player
		}

		return m
	}

	for {
		{
			players, err := palRCON.GetPlayers()

			if err != nil {
				slog.Error("failed to get players", "error", err)
				goto NEXT
			}

			slog.Debug("Current players", "players", players)

			playersMap := makeMap(players)

			if prev == nil {
				prev = playersMap
				goto NEXT
			}

			for _, player := range players {
				if _, ok := prev[player.SteamID]; !ok {
					err := palRCON.Broadcast(fmt.Sprintf("joined:%s", player.Name))
					if err != nil {
						slog.Error("failed to broadcast", "error", err)
					}

					slog.Info("Player joined", "player", player)
				}
			}
			for _, player := range prev {
				if _, ok := playersMap[player.SteamID]; !ok {
					err := palRCON.Broadcast(fmt.Sprintf("left:%s", player.Name))
					if err != nil {
						slog.Error("failed to broadcast", "error", err)
					}

					slog.Info("Player left", "player", player)
				}
			}

			prev = playersMap
		}
	NEXT:
		time.Sleep(interval)
	}
}
