package service

import (
	"cli_notes/internal/entity"
	"cli_notes/internal/postgresql"
	"fmt"
)

// Add ...
func Add(title string, body string, tag string) (int, error) {
	if len(title) > 30 {
		return 0, fmt.Errorf("title length too long")
	}
	note := entity.NewNote(title, body, tag)
	id, err := postgresql.Insert(*note)
	if err != nil {
		return 0, fmt.Errorf("unable to add note: %v\n", err)
	}
	return id, nil
}

// Delete ...
func Delete() {

}

// Edit ...
func Edit() {

}
