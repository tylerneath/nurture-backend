package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerneath/nuture-backend/internal/config"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"github.com/tylerneath/nuture-backend/internal/store"
	"gorm.io/gorm/logger"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief descirption of you command",
	Long:  "Long description of your command",
	Run:   migrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	cfg := config.Config{}
	viper.UnmarshalKey("database", &cfg)
	db := store.MustCreateNewDB(cfg)
	db.Logger.LogMode(logger.Info)

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		panic(fmt.Sprintf("error creating extension: %v", err))
	}

	println("applying migrations")
	db.AutoMigrate(&models.User{}, &models.Message{})

}
