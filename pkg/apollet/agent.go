package apollet

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Colstuwjx/apollet/pkg/config"
)

type Agent struct {
	config *config.Config
	client *Client
}

func NewAgent(conf *config.Config) *Agent {
	if conf.Apollo == nil {
		return nil
	}

	// Init the agollo clients
	client := NewClient(
		conf.Apollo.Server,
	)

	if client == nil {
		return nil
	}

	return &Agent{
		config: conf,
		client: client,
	}
}

func (this *Agent) storePid() error {
	// Quit fast if no pidfile
	pidPath := this.config.Pid
	if pidPath == "" {
		return nil
	}

	// Open the PID file
	pidFile, err := os.OpenFile(pidPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Could not open pid file: %v", err)
	}
	defer pidFile.Close()

	// Write out the PID
	pid := os.Getpid()
	_, err = pidFile.WriteString(fmt.Sprintf("%d", pid))
	if err != nil {
		return fmt.Errorf("Could not write to pid file: %s", err)
	}
	return nil
}

func (this *Agent) Start() {
	// Write out the PID file if necessary
	if err := this.storePid(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Startup the http server
	go this.ServeHTTP()

	// Handle exit
	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-term:
			fmt.Println("Exiting...")
			os.Exit(2)
		}
	}
}
