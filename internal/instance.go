package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

type Instance struct {
	client *firebase.App
}

func NewInstance(client *firebase.App) *Instance {
	return &Instance{
		client: client,
	}
}

func (i *Instance) Firestore(ctx context.Context) (*Firestore, error) {
	client, err := i.client.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get firestore client: %v", err)
	}
	return NewFirestore(client), nil
}
