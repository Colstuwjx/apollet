package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Colstuwjx/apollet/pkg/apollet"
	"github.com/Colstuwjx/apollet/pkg/config"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "apollet",
		Short: "apollet is an agent for apollo config center",
		Long:  `A reliable, multiple language supported agent for apollo, refers: https://github.com/Colstuwjx/apollet`,
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := config.NewConf(cfgFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			agent := apollet.NewAgent(conf)
			if agent == nil {
				fmt.Println("Init agent failed!")
				os.Exit(1)
			}

			agent.Start()
		},
	}
)

func init() {
	rootCmd.Flags().StringVar(&cfgFile, "conf", "", "config file path (default is testdata/conf.toml)")
	rootCmd.MarkFlagRequired("conf")

	// add sub cmds
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
