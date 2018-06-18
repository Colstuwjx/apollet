package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Colstuwjx/apollet/pkg/apollctl"
	"github.com/Colstuwjx/apollet/pkg/apollet"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get value via calling apollet",
	Long:  `Get value via calling apollet, should offer APOLLET_ADDR env or set in ~/.apollet file`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add config parser for apollctl.
		// client, err := apollctl.NewClient("unix", "/tmp/apollet.sock")
		client, err := apollctl.NewClient("http", "127.0.0.1:10810")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// TODO: make flags controlled.
		var (
			appId     = "SampleApp"
			cluster   = "default"
			namespace = "application"
			key       = args[0]
		)

		value := client.GetString(appId, cluster, namespace, key)
		if value != apollet.NotFoundDefaultValue {
			fmt.Println(value)
			os.Exit(0)
		}

		os.Exit(1)
	},
}
