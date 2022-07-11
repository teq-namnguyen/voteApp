package main

import (
	"context"
	"github.com/namnguyen/backend/client/postgresql"
	"github.com/namnguyen/backend/migration"
)

func main() {
	ctx := context.Background()
	client := postgresql.GetClient
	migration.Up(client(ctx))
}
