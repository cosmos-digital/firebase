package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Config = firebase.Config
type Options = option.ClientOption

type connection struct {
	opts   Options
	config *Config
}

func New(opts Options, config *Config) *connection {
	return &connection{
		opts:   opts,
		config: config,
	}
}

func (c *connection) Connect(ctx context.Context) (*Instance, error) {
	client, err := firebase.NewApp(ctx, c.config, c.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %v", err)
	}

	return NewInstance(client), nil
}
