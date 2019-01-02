package model

type Path struct {
	Address string `json:"address" xorm:"'ADDRESS'"`
	Name    string `json:"name" xorm:"'NAME'"`
}

type Create struct {
	Address string `json:"address" xorm:"'ADDRESS'"`
	Name    string `json:"name" xorm:"'NAME'"`
}

type Delete struct {
	Address string `json:"address" xorm:"'ADDRESS'"`
	Name    string `json:"name" xorm:"'NAME'"`
}

type Response struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
	Error   string   `json:"error"`
}

type ResponseList struct {
	Success bool       `json:"success"`
	Data    [20]string `json:"data"`
	Error   string     `json:"error"`
}

type List struct {
	Address string
}

type Import struct {
	From string
	To   string
	Name string
}
