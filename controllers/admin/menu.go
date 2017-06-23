package admin

//Menu 菜单结构体
type Menu struct {
	Name    string
	Icon    string
	Module  string
	Action  string
	Submenu []*Submenu
	Role    string
}

//Submenu 子菜单结构体
type Submenu struct {
	Name   string
	Icon   string
	Action string
	Parent string
}
