package models

type Customer struct {
	ID       string `form:"id,omitempty"`
	Name     string `form:"name"`
	LastName string `form:"lastName"`
	Phone    int    `form:"phone"`
	Address  string `form:"address"`
	Fav      string `form:"fav"`
}

type Order struct {
	ID       string   `form:"id,omitempty"`
	Number   int      `form:"number"`
	Customer Customer `form:"customer"`
	Product  string   `form:"product"`
}
