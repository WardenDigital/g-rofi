/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/WardenDigital/rofi-wrapper/internal/handlers/search"
	"github.com/spf13/cobra"
)

// browserSearchCmd represents the browserSearch command
var browserSearchCmd = &cobra.Command{
	Use:   "browserSearch",
	Short: "Search using the browser",
	Long:  `This command allows you to perform a search using your default web browser.`,
	Run: func(cmd *cobra.Command, args []string) {
		search.Search()
	},
}

func init() {
	rootCmd.AddCommand(browserSearchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// browserSearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// browserSearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
