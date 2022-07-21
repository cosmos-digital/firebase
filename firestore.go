package firebase

import (
	"context"
	"fmt"

	fb "firebase.google.com/go"
	firebase "github.com/cosmos-digital/firebase/internal"
	"google.golang.org/api/option"
)

type Config = fb.Config
type Options = option.ClientOption

func NewFirestore(ctx context.Context, opts Options, config *Config) (*firebase.Firestore, error) {
	connection := firebase.New(opts, config)
	instance, err := connection.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %v", err)
	}
	client, err := instance.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get firestore client: %v", err)
	}
	return client, nil
}
