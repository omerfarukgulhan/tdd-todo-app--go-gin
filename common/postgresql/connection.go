package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime)

	connConfig, parseConfigErr := pgxpool.ParseConfig(connString)
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}

	conn, err := pgxpool.ConnectConfig(context, connConfig)
	if err != nil {
		log.Fatal("Unable to connect to database: %v\n", err)
	}

	createTables(context, conn)

	return conn
}

func createTables(ctx context.Context, dbPool *pgxpool.Pool) {
	createTodoTableQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL,                 
		title VARCHAR(255) NOT NULL,
		description TEXT,
		is_completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)  
	);
	`

	_, err := dbPool.Exec(ctx, createTodoTableQuery)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	log.Println("Tables created or already exist.")
}
