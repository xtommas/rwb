/*
Copyright © 2024 Tomas
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rwb",
	Short: "Get a random anime girl from Wallhaven and set it as the desktop background",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for getting a random image
		// and setting it as background
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
