package apollet

import (
	"log"

	"github.com/Colstuwjx/apollet/pkg/config"
)

type Agent struct {
	config *config.Config
}

func NewAgent(conf *config.Config) *Agent {
	return &Agent{
		config: conf,
	}
}

func (this *Agent) Start() {
	// TODO: implement agent and startup, listen signals.
	// init client and offers IPC communication.
	log.Println("Conf: ", this.config)
	log.Println("Start agent...")
}
