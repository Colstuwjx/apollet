package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "apollctl",
		Short: "apollctl is an CLI tool for apollet",
		Long:  `A simple to use CLI tool for apollet, refers: https://github.com/Colstuwjx/apollet`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: empty cmd should be forbidden.
		},
	}
)

func init() {
	// add sub cmds
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
