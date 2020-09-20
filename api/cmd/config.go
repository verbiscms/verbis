package cmd

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Used to cache and clear Verbis's configuration.",
	}
)

func init() {
	configCmd.AddCommand(configStoreCmd)
	configCmd.AddCommand(configClearCmd)
}

// cacheClearCmd represents the down command
var configStoreCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache all of the configuration files.",
	Long: `Cache will cache all of the configuration files based
on configuration path ./config. It will cache the yaml files
defined for faster reading. Recommended in production`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Cache()
		color.Green.Println("Successfully cached the configuration files.")
	},
}

// cacheClearCmd represents the down command
var configClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear will delete the configuration cache.",
	Long: `Cache will delete the cache all of the configuration 
files based on configuration path ./config.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.CacheClear()
		color.Green.Println("Successfully cleared the configuration files cache.")
	},
}