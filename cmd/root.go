package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gips",
	Short: "GIPS is a cross-platform IPS patcher",
	Long:  "GIPS is a cross-platform IPS patcher built with Go, built by Aria Taylor. Check out their website at https://aricodes.net",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("We here.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
