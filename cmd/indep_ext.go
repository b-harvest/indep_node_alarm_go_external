// main.go
// 54.95.40.202
package main

import (
	"indep_node_alarm_go_external/pkg/config"
	"indep_node_alarm_go_external/pkg/dialer"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "indep_go_ext",
	Short: "indep_go_ext checks the TCP dial status of servers",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig(viper.GetString("config"))
		if err != nil {
			//log.Fatal(err)
			log.Fatal().Err(err).Msg("Failed to load config file")
		}

		dialer.DialServers(config)
	},
}

func init() {
	pflag.String("config", "config.toml", "path to the config file")
	viper.BindPFlags(pflag.CommandLine)
	rootCmd.AddCommand()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Failed to execute root command")
	}
}
