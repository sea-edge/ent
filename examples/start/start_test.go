package main

import (
	"context"
	"testing"

	"entgo.io/ent/examples/start/ent"
	"entgo.io/ent/examples/start/ent/enttest"
	"entgo.io/ent/examples/start/ent/migrate"
)

func Test_main(t *testing.T) {
	initializeEntClient(t)
}

func initializeEntClient(t *testing.T) (*ent.Client, error) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	return client, client.Schema.Create(context.Background())
}
