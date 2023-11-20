/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command

// init will initialize the standard tmp folder where later all files and folders will be stored
// based on creation date
// will be set eighter in config or in shell config
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize standard tmp-cli folder",
	Long: `initialises standard tmp-cli folder and sets it in config or bash config. This folder will be
	used to store all files if not set by flag`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
