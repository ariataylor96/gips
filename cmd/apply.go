package cmd

import (
	"errors"
	"fmt"
	"gips/records"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(applyCmd)
}

func exitWithError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply an IPS patch to a ROM",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("This command only takes 2 arguments, a ROM file and an IPS file")
		}

		var hasIpsFile bool = false
		for _, val := range args {
			if strings.ToLower(filepath.Ext(val)) == ".ips" {
				hasIpsFile = true
				break
			}
		}

		if !hasIpsFile {
			return errors.New("Please provide an IPS file")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		var ipsFileName, romFileName string

		// Let the user supply the file names in any order
		for _, fileName := range args {
			if strings.ToLower(filepath.Ext(fileName)) != ".ips" {
				romFileName = fileName
			} else {
				ipsFileName = fileName
			}
		}

		// We probably shouldn't apply an IPS patch to another IPS patch
		if (romFileName) == "" {
			exitWithError("Please provide a ROM file!")
		}

		ipsRecords := records.FromFile(ipsFileName)
		romData, err := os.ReadFile(romFileName)
		if err != nil {
			panic(err)
		}

		for _, record := range ipsRecords {
			record.Apply(&romData)
		}

		os.WriteFile("out.rom", romData, 0644)
	},
}
