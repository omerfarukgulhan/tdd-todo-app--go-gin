package infrastructure

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
	"todo-app--go-gin/common/postgresql"
	"todo-app--go-gin/persistence"
)

var userRepository persistence.IUserRepository
var todoRepository persistence.ITodoRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "5432",
		UserName:              "postgres",
		Password:              "153515",
		DbName:                "workshops",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})

	todoRepository = persistence.NewTodoRepository(dbPool)
	userRepository = persistence.NewUserRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func SetupData(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func ClearData(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05", timeStr)
}

func MustParseTime(timeStr string) time.Time {
	t, err := ParseTime(timeStr)
	if err != nil {
		panic(err)
	}

	return t
}
