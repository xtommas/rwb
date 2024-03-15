/*
Copyright © 2024 Tomas
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/reujab/wallpaper"
	"github.com/spf13/cobra"
	"github.com/xtommas/rwb/api"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rwb",
	Short: "Get a random anime girl from Wallhaven and set it as the desktop background",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wallURL, err := api.GetRandomWall()
		if err != nil {
			panic(err)
		}

		err = wallpaper.SetFromURL(wallURL)
		if err != nil {
			panic(err)
		}
		fmt.Println("Waifu set!")
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
