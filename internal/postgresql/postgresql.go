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
	pool, err := NewPool()
	if err != nil {
		log.Println("failed connection to db", err)
		pool = Reconnect()
		return 0, fmt.Errorf("unable to insert: %v\n", err)
	}

	defer pool.Close()
	row := pool.QueryRow(context.Background(), "insert into notes (Title, Body, Tag) values ($1, $2, $3) returning ID",
		note.GetTitle(), note.GetBody(), note.GetTag())
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to insert: %v\n", err)
	}

	return id, nil
}

// Reconnect trying to сonnect 3 times with 2 seconds interval to database if connection was lost
func Reconnect() *pgxpool.Pool {
	var err error
	var pool *pgxpool.Pool
	for i := 0; i <= 3; i++ {
		pool, err = NewPool()

		if err == nil {
			err = pool.Ping(context.Background())
			if err == nil {
				log.Println("reconnected to db successfully!")
				return pool
			}
		}
		time.Sleep(2 * time.Second)
	}

	log.Fatal("unable to connect to db")
	return nil
}
