package commons

import "time"

const (
	ServiceName = "kubemq-consumer"

	HttpServerHostFormat          = "%s:%d"
	HttpServerWriteTimeoutDefault = time.Second * 15
	HttpServerReadTimeoutDefault  = time.Second * 15
	HttpServerIdelTimeoutDefault  = time.Second * 60
)
