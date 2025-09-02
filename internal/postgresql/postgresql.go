package postgresql

import (
	"cli_notes/internal/entity"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

func NewPool() (*pgxpool.Pool, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST_IP"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v\n", err)
	}
	if err := dbPool.Ping(context.Background()); err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("unable to ping database: %v\n", err)
	}
	log.Println("db connection pool created successfully")
	return dbPool, nil
}

// Insert ...
func Insert(note entity.Note) (int, error) {
	pool, err := ConnectWithRetry(3)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()
	row := pool.QueryRow(context.Background(), "insert into notes (Title, Body, Tag) values ($1, $2, $3) returning ID",
		note.Title, note.Body, note.Tag)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to insert: %v\n", err)
	}

	return id, nil
}

// GetAll ...
func GetAll() (*[]entity.Note, error) {
	pool, err := ConnectWithRetry(3)
	if err != nil {
		log.Fatal(err)
	}

	var notes []entity.Note
	defer pool.Close()
	rows, err := pool.Query(context.Background(), "select * from notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		note := entity.Note{}
		err = rows.Scan(&note.ID, &note.Title, &note.Body, &note.Tag)
		if err != nil {
			log.Println(err)
			continue
		}
		notes = append(notes, note)
	}
	if len(notes) == 0 {
		return nil, fmt.Errorf("empty table db")
	}
	return &notes, nil
}

// ConnectWithRetry trying to сonnect 3 times with 2 seconds interval to database
func ConnectWithRetry(maxAttempts int) (*pgxpool.Pool, error) {
	var err error
	var pool *pgxpool.Pool
	for i := 1; i <= maxAttempts; i++ {
		pool, err = NewPool()
		if err == nil {
			return pool, nil
		}
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("unable to connect to db: %v\n", err)
}
