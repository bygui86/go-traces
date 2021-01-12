package database

import (
	"github.com/bygui86/go-traces/http-server/logging"
)

// TODO replace with https://github.com/hashicorp/go-memdb ?

func New() InMemoryDbInt {
	logging.Log.Info("Create new in-memory DB")

	return &InMemoryDb{
		products: make(map[string]*Product, 1000),
	}
}
