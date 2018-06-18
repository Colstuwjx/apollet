package apollet

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (this *Agent) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (this *Agent) getString(c *gin.Context) {
	var gsr GetStringRequest

	value := NotFoundDefaultValue
	err := c.Bind(&gsr)
	if err != nil {
		fmt.Println("Err: ", err)
		c.JSON(400, gin.H{"data": ""})
		return
	}

	value = this.client.GetString(gsr.AppId, gsr.Cluster, gsr.Namespace, gsr.Key, NotFoundDefaultValue)
	if value == NotFoundDefaultValue {
		c.JSON(400, gin.H{"data": ""})
		return
	}

	c.JSON(200, gin.H{"data": value})
	return
}
