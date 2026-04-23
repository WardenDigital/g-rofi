/*
Copyright © 2026 WardenDigital <wardendigital@proton.me>
*/
package cmd

import (
	"github.com/WardenDigital/g-rofi/internal/handlers/show"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Alias for rofi show",
	Long:  `Rofi show command with one param`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		show.Show(args[0])
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
