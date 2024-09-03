package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var INSERT_USERS = `INSERT INTO users (id, username, email, password) 
VALUES
  (1, 'user1', 'user1@mail.com', '12345'),
  (2, 'user2', 'user2@mail.com', '12345'),
  (3, 'user3', 'user3@mail.com', '12345'),
  (4, 'user4', 'user4@mail.com', '12345');
`

var INSERT_TODOS = `INSERT INTO todos (id, user_id, title, description, is_completed, created_at, updated_at) 
VALUES
  (1, 1, 'Buy groceries', 'Purchase fruits, vegetables, and bread', false, '2024-09-01 10:00:00', '2024-09-01 10:00:00'),
  (2, 1, 'Complete assignment', 'Finish the report for the upcoming meeting', true, '2024-09-02 09:30:00', '2024-09-02 09:30:00'),
  (3, 1, 'Workout session', 'Attend the gym for a cardio session', false, '2024-09-03 18:00:00', '2024-09-03 18:00:00'),
  (4, 2, 'Read a book', 'Start reading a new novel', true, '2024-09-04 20:00:00', '2024-09-04 20:00:00');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertUsersResult, insertUsersErr := dbPool.Exec(ctx, INSERT_USERS)
	if insertUsersErr != nil {
		log.Printf("Error inserting todos data: %v", insertUsersErr)
	} else {
		log.Printf("Todos data created with %d rows", insertUsersResult.RowsAffected())
	}

	insertTodosResult, insertTodosErr := dbPool.Exec(ctx, INSERT_TODOS)
	if insertTodosErr != nil {
		log.Printf("Error inserting todos data: %v", insertTodosErr)
	} else {
		log.Printf("Todos data created with %d rows", insertTodosResult.RowsAffected())
	}

}
