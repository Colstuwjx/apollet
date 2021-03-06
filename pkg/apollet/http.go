package apollet

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (this *Agent) ServeHTTP() {
	if !this.config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/ping", this.ping)
	r.GET("/get_string", this.getString)

	if this.config.Http == nil {
		fmt.Println("Http host and port must be specified!")
		os.Exit(1)
	}

	switch this.config.Http.Scheme {
	case "http":
		bind := net.JoinHostPort(this.config.Http.Host, strconv.Itoa(this.config.Http.Port))
		err := r.Run(bind)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "unix":
		bind := this.config.Http.Host
		err := r.RunUnix(bind)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Scheme ", this.config.Http.Scheme, " did not support yet.")
		os.Exit(1)
	}
}
