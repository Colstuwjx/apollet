package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of apollctl",
	Long:  `All software has versions. This is apollctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apollctl v0.0.1")
	},
}
