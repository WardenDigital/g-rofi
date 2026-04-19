/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/WardenDigital/g-rofi/internal/handlers/search"
	"github.com/spf13/cobra"
)

// browserSearchCmd represents the browserSearch command
var browserSearchCmd = &cobra.Command{
	Use:   "browserSearch",
	Short: "Search using the browser",
	Long:  `This command allows you to perform a search using your default web browser.`,
	Run: func(cmd *cobra.Command, args []string) {
		params := cmd.Flags()
		search.Search(params.Lookup("browser").Value.String(), params.Lookup("engine").Value.String())
	},
}

func init() {
	rootCmd.AddCommand(browserSearchCmd)

	browserSearchCmd.Flags().StringP("browser", "b", "brave-browser", "Specify the browser to use for searching")
	browserSearchCmd.Flags().StringP("engine", "e", "google", "Specify the search engine to use")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// browserSearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// browserSearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
