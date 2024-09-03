package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE todos RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Printf(truncateResultErr.Error())
	} else {
		log.Printf("Todos table truncated")
	}

	_, truncateResultErr = dbPool.Exec(ctx, "TRUNCATE users RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Printf(truncateResultErr.Error())
	} else {
		log.Printf("Users table truncated")
	}
}
