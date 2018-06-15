package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of apollet",
	Long:  `All software has versions. This is apollet's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Apollet v0.0.1")
	},
}
