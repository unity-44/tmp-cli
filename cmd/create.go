/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/tmp-cli/helpers"
)

// nameValue varialbe for setting file name
var nameValue string

const defaultFile = "tmp.txt"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates new file",
	Long:  `creates a new file with which is stored in your default location. With additional flags you can open it or set a different location:`,

	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("name") {
			helpers.CreateFile(nameValue)
			helpers.OpenFileInCode(nameValue)
		} else {
			helpers.CreateFile(defaultFile)
			helpers.OpenFileInCode(defaultFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&nameValue, "name", "n", "", "Filename")

}
