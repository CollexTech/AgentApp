package cmd

import (
	"backend/constants"
	"backend/repository/datastore"
	"backend/server"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCommand = &cobra.Command{
		Use:   ``,
		Short: "nodeops oms",
		PreRun: func(cmd *cobra.Command, args []string) {
			loadConfigFiles()
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("nodeops oms")
		},
	}
	serveHTTP = &cobra.Command{
		Use:   `serve-http`,
		Short: "Start HTTP server",
		PreRun: func(cmd *cobra.Command, args []string) {
			loadConfigFiles()
			datastore.Get()
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initiating HTTP server")
			server.Initialize()
		},
	}
)

func RegisterCommands() {
	rootCommand.AddCommand(serveHTTP)
	rootCommand.Execute()
}

func loadConfigFiles() {

	viper.AutomaticEnv()
	currentEnv := viper.GetString("ENVIRONMENT")
	fmt.Printf("Current env: %s\n", currentEnv)

	viper.SetConfigType("yaml")
	switch currentEnv {
	case constants.ENVIRONMENT_PRODUCTION:
		viper.AddConfigPath("./config")
		viper.SetConfigName("prod")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	case constants.ENVIRONMENT_STAGING:
		viper.AddConfigPath("./config")
		viper.SetConfigName("stage")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	default:
		viper.AddConfigPath("./config")
		viper.SetConfigName("dev")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	viper.AutomaticEnv()
}
