package admin

type Menu struct {
	Name    string
	Icon    string
	Module  string
	Action  string
	Submenu []*Submenu
	Role    string
}

type Submenu struct {
	Name   string
	Icon   string
	Action string
	Parent string
}
