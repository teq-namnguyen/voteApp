package migration

import (
	"context"
	"github.com/namnguyen/backend/client/postgresql"
)

func main() {
	ctx := context.Background()
	client := postgresql.GetClient
	Up(client(ctx))
}
