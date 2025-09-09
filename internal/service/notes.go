package service

import (
	"cli_notes/internal/entity"
	"cli_notes/internal/postgres"
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
	id, err := postgres.Insert(*note)
	if err != nil {
		return 0, fmt.Errorf("unable to add note: %v\n", err)
	}
	return id, nil
}

// Delete ...
func Delete(arg string, key string) (int64, error) {
	res, err := postgres.DeleteFromBD(arg, key)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// Edit ...
func Edit(flag, arg, change, newValue string) (*entity.Note, error) {
	note, err := postgres.Get(flag, arg)
	if err != nil {
		return nil, err
	}

	switch change {
	case "title":
		if len(newValue) > 30 {
			return nil, fmt.Errorf("too long title (must be less than 30 chars)")
		}
		note.Title = newValue
	case "body":
		note.Body = newValue
	case "tag":
		if len(newValue) > 30 {
			return nil, fmt.Errorf("too long tag (must be less than 30 chars)")
		}
		note.Tag = newValue
	case "done":
		fmt.Println("in the case done")
		done, err := strconv.ParseBool(newValue)
		if err != nil {
			return nil, fmt.Errorf("not a bool given")
		}
		note.Done = done
	}
	_, err = postgres.Insert(*note)
	if err != nil {
		return nil, err
	}
	return note, nil

}

// List ...
func List() error {
	notes, err := postgres.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', 1)
	fmt.Fprintln(w, "ID\t|Title\t|Body\t|Tag\t|Done\t")
	for _, note := range *notes {
		fmt.Fprintln(w, strconv.Itoa(note.ID)+"\t|"+note.Title+"\t|"+note.Body+"\t|"+note.Tag+"\t|"+strconv.FormatBool(note.Done)+"\t")
	}
	w.Flush()
	return nil
}

// Done ...
func Done(title string, done bool) error {
	_, err := Edit("title", title, "done", strconv.FormatBool(done))
	if err != nil {
		return err
	}
	return nil
}
