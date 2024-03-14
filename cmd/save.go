/*
Copyright © 2024 Tomas
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	savePath string
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "saves the current background picture to a folder (Pictures/waifus by default)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Logic here
		fmt.Println("save called")
	},
}

func init() {
	saveCmd.Flags().StringVarP(&savePath, "path", "p", "", "The path where the image should be saved")
	rootCmd.AddCommand(saveCmd)
}
