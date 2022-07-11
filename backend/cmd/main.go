package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/namnguyen/backend/client/postgresql"
	"github.com/namnguyen/backend/delivery/http"
	"github.com/namnguyen/backend/migration"
	"github.com/namnguyen/backend/repository"
	"github.com/namnguyen/backend/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	taskPtr := flag.String("task", "server", "server")

	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Local = loc
	}

	ctx := context.Background()
	client := postgresql.GetClient
	repo := repository.New(client)
	useCase := usecase.New(repo)

	migration.Up(client(ctx))

	switch *taskPtr {
	case "server":
		executeServer(useCase)
	default:
		executeServer(useCase)
	}

	defer postgresql.Disconnect()
}

func executeServer(useCase *usecase.UseCase) {
	h := http.NewHTTPHandler(useCase)
	h.Logger.Fatal(h.Start(":9000"))
}
