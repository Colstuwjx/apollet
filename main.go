package main

import (
	"github.com/Colstuwjx/apollet/cmd"
)

func main() {
	/*
	   1. start the agent;
	   2. shared memory;
	   3. cli and polling, persistent;
	   4. offer IPC communication via unixsocket;
	*/

	cmd.Execute()
}
