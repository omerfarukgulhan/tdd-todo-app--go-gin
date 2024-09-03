package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE todos RESTART IDENTITY CASCADE")
	if truncateResultErr != nil {
		log.Printf("Error truncating todos table: %v", truncateResultErr)
	} else {
		log.Printf("Todos table truncated")
	}

	_, truncateResultErr = dbPool.Exec(ctx, "TRUNCATE users RESTART IDENTITY CASCADE")
	if truncateResultErr != nil {
		log.Printf("Error truncating users table: %v", truncateResultErr)
	} else {
		log.Printf("Users table truncated")
	}
}
