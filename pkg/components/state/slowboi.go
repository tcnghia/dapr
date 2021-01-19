// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package state

import (
	"strconv"
	"time"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/dapr/pkg/logger"
)

// SlowStateStore is a slow state store
type SlowStateStore struct {
	state.DefaultBulkStore
	Logger logger.Logger
}

// NewSlowStateStore creates a new slow state store
func NewSlowStateStore(logger logger.Logger) *SlowStateStore {
	return &SlowStateStore{Logger: logger}
}

func (r *SlowStateStore) Init(metadata state.Metadata) error {
	text, ok := metadata.Properties["delay"]

	delay := 10
	if ok {
		i, err := strconv.ParseInt(text, 10, 32)
		if err == nil {
			delay = int(i)
		}
	}

	for i := 0; i < delay; i++ {
		r.Logger.Info("Sorry, can't contact the server. Just give me a minute to get my coffee.")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (r *SlowStateStore) Delete(req *state.DeleteRequest) error {
	return nil
}

func (r *SlowStateStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	return nil, nil
}

func (r *SlowStateStore) Set(req *state.SetRequest) error {
	return nil
}

func (r *SlowStateStore) Multi(request *state.TransactionalStateRequest) error {
	return nil
}
