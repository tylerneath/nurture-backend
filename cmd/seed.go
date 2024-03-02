/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerneath/nuture-backend/internal/config"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"github.com/tylerneath/nuture-backend/internal/store"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: seed,
}

func seed(cmd *cobra.Command, args []string) {
	fmt.Println("seeding DB")
	cfg := config.Config{}
	viper.UnmarshalKey("database", &cfg)
	// steps -> open a new connection for the db
	// check if there are users in the containers
	// if there are immediately exist
	// else, run queries into the new db
	db := store.MustCreateNewDB(cfg)
	db.AutoMigrate(&models.User{}, &models.Message{})

	// create two new users
	for i := 0; i < 1; i++ {
		if db.Create(&models.User{
			Email: "ragarig",
		}).Error != nil {
			fmt.Println("error creating user")
		}
	}

}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
