package worker

import (
	"context"
	"time"
)

type Worker struct {
	config  *config
	ctx     context.Context
	ticker  *time.Ticker
	running bool
}

type config struct {
	workingInterval time.Duration
}
