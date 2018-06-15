package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apollet",
	Short: "Apollet is an agent for apollo config center",
	Long:  `A reliable agent for apollo, refers: https://github.com/Colstuwjx/apollo-agent`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("CMD: ", cmd)
		log.Println("Args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	/*
	   1. start the agent
	   2. offers CLI
	   3. IPC communication, persistent configs etc
	*/
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
