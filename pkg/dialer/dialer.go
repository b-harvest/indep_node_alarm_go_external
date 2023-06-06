package dialer

import (
	"indep_node_alarm_go_external/pkg/alert"
	"indep_node_alarm_go_external/pkg/config"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func DialServers(config config.Config) {
	for {
		var wg sync.WaitGroup
		for name, address := range config.Servers {
			wg.Add(1)
			go func(name, address string) {
				defer wg.Done()
				conn, err := net.Dial("tcp", address)
				if err != nil {
					log.Error().Err(err).Str("server", name).Msg("Failed to connect to server")
					alert.SendAlerts(err, config, name, address)
					return
				}
				log.Info().Str("server", name).Msg("Successfully connected to server")
				defer conn.Close()
			}(name, address)
		}
		wg.Wait()
		time.Sleep(10 * time.Second) // adjust the sleep duration as needed
	}
}
