package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	client *firestore.Client
}

func NewFirestore(client *firestore.Client) *Firestore {
	return &Firestore{
		client: client,
	}
}

func (s *Firestore) Save(ctx context.Context, collection string, data interface{}) (*firestore.DocumentRef, error) {
	ref, _, err := s.client.Collection(collection).Add(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to save data: %v", err)
	}
	return ref, nil
}

func (s *Firestore) Close() error {
	return s.client.Close()
}
