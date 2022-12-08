package entity

type MyTeam struct {
	ID      string
	Name    string
	Players []Player
	Score   float64
}

func NewMyTeam(id, name string) MyTeam {
	return MyTeam{
		ID:   id,
		Name: name,
	}
}
