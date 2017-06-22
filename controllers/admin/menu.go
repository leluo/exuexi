package admin

type Menu struct {
	Icon    string
	Fa      string
	Href    string
	Name    string
	Disable bool
	subs    []*SubMenu
}

type SubMenu struct {
	Href    string
	Name    string
	Disable bool
}

func NewMenu() *Menu {
	return new(Menu)
}

func (m *Menu) GetHome(sub ...SubMenu) Menu {
	menu := Menu{
		Icon:    "color5",
		Fa:      "fa-home",
		Href:    "Index.Index",
		Name:    "首页",
		Disable: true,
		subs: []*SubMenu{
			&SubMenu{Href: "Index.Home", Name: "哈哈驾校", Disable: true},
			&SubMenu{Href: "Index.Ribe", Name: "成电驾校", Disable: true},
		},
	}
	return menu
}
