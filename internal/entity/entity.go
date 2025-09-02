package entity

type Note struct {
	id    int
	title string
	body  string
	tag   string
}

// NewNote ...
func NewNote(title string, body string, tag string) *Note {
	return &Note{
		id:    -1,
		title: title,
		body:  body,
		tag:   tag,
	}
}

func (n *Note) GetID() int {
	return n.id
}

func (n *Note) GetTitle() string {
	return n.title
}

func (n *Note) GetBody() string {
	return n.body
}

func (n *Note) GetTag() string {
	return n.tag
}
