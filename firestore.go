package firebase

import (
	"context"
	"fmt"

	firebase "github.com/cosmos-digital/firebase/internal"
)

type Config = firebase.Config

func NewFirestore(ctx context.Context, opts firebase.Options, config *firebase.Config) (*firebase.Firestore, error) {
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
