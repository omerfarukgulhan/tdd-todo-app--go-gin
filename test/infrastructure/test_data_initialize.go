package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var INSERT_TODOS = `INSERT INTO todos (id, user_id, title, description, is_completed, created_at, updated_at) 
VALUES
  ('1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p', '2c3d4e5f-6g7h-8i9j-0k1l-2m3n4o5p6q7r', 'Buy groceries', 'Purchase fruits, vegetables, and bread', false, '2024-09-01 10:00:00', '2024-09-01 10:00:00'),
  ('2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7q', '3d4e5f6g-7h8i-9j0k-1l2m-3n4o5p6q7r8s', 'Complete assignment', 'Finish the report for the upcoming meeting', true, '2024-09-02 09:30:00', '2024-09-02 09:30:00'),
  ('3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8r', '2c3d4e5f-6g7h-8i9j-0k1l-2m3n4o5p6q7r', 'Workout session', 'Attend the gym for a cardio session', false, '2024-09-03 18:00:00', '2024-09-03 18:00:00'),
  ('4d5e6f7g-8h9i-0j1k-2l3m-4n5o6p7q8r9s', '4e5f6g7h-8i9j-0k1l-2m3n-4o5p6q7r8s9t', 'Read a book', 'Start reading a new novel', true, '2024-09-04 20:00:00', '2024-09-04 20:00:00');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertTodosResult, insertTodosErr := dbPool.Exec(ctx, INSERT_TODOS)
	if insertTodosErr != nil {
		log.Printf("Error inserting todos data: %v", insertTodosErr)
	} else {
		log.Printf("Todos data created with %d rows", insertTodosResult.RowsAffected())
	}
}
