package entity

type Note struct {
	ID    int
	Title string
	Body  string
	Tag   string
	Done  bool
}

// NewNote ...
func NewNote(title string, body string, tag string) *Note {
	return &Note{
		ID:    -1,
		Title: title,
		Body:  body,
		Tag:   tag,
		Done:  true,
	}
}
