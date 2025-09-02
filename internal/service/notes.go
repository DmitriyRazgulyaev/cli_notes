package service

import (
	"cli_notes/internal/entity"
	"cli_notes/internal/postgresql"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
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

// List ...
func List() error {
	notes, err := postgresql.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', 1)
	fmt.Fprintln(w, "ID\t|Title\t|Body\t|Tag\t")
	for _, note := range *notes {
		fmt.Fprintln(w, strconv.Itoa(note.ID)+"\t|"+note.Title+"\t|"+note.Body+"\t|"+note.Tag+"\t")
	}
	w.Flush()
	return nil
}
