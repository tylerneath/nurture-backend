/*
Copyright © 2023 Tyler Neath tylerneath24@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerneath/nuture-backend/internal/api"
	"github.com/tylerneath/nuture-backend/internal/config"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  "Long description for the server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolP("toggle", "t", false, "toggle the whatever")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {
	cfg := config.Config{}
	logger := zap.Must(zap.NewDevelopment())

	defer logger.Sync() // flushes buffer, if any

	if err := viper.UnmarshalKey("database", &cfg); err != nil {
		fmt.Errorf("viper not working")
	}

	print(cfg.String())

	api.Run(cmd.Context(), cfg, logger)
}
