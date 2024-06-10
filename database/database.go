package database

import (
	"context"
	"fmt"
	"log"

	"api/ent"

	"api/ent/migrate"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func NewEntClient(database string) (*ent.Client, error) {
	ctx := context.Background()
	// database := "employee"
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"welcome@123",
		"postgres",
		database,
	)

	client, err := ent.Open(dialect.Postgres, dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
	}))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	errSchema := client.Schema.Create(ctx, migrate.WithDropIndex(true))
	if errSchema != nil {
		log.Fatalf("failed creating schema resources for database %v : %v", database, err)
	}

	return client, nil
}
