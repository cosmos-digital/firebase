package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos-digital/firebase/firestore"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(wd)
	opts := option.WithCredentialsFile(
		fmt.Sprintf("%s/sample.json", dir),
	)
	config := &firestore.Config{
		ProjectID: "theta-trees-339100",
	}
	client, err := firestore.New(ctx, opts, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ref, err := client.Save(ctx, "collection", map[string]interface{}{
		"name": "John Doe",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(ref.ID)
	os.Exit(0)
}
