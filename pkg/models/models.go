package models

type Host struct {
	Id    uint16
	Name  string
	Ip    string
	Mac   string
	Hw    string
	Date  string
	Known uint16
	Now   uint16
}

type Conf struct {
	Iface    string
	DbPath   string
	GuiIP    string
	GuiPort  string
	Timeout  int
	ShoutUrl string
	Theme    string
}
