package agollo

import (
	"time"
)

const (
	defaultConfName  = "app.properties"
	defaultDumpFile  = ".agollo"
	defaultNamespace = "application"

	longPoolInterval      = time.Second * 2
	longPoolTimeout       = time.Second * 90
	queryTimeout          = time.Second * 2
	defaultNotificationID = -1
)
