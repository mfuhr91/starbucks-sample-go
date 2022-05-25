package models

type Customer struct {
	ID       string `form:"id,omitempty"`
	Name     string `form:"name"`
	LastName string `form:"lastName"`
	Doc      int    `form:"doc"`
	Phone    int    `form:"phone"`
	Address  string `form:"address"`
	Disabled bool
}

type Order struct {
	ID         string   `json:"id,omitempty"`
	Number     int      `json:"number,omitempty"`
	Customer   Customer `json:"customer"`
	Items      []Item   `json:"items"`
	FinalPrice float64  `json:"price,string"`
	Time       string   `json:"time" time_format:"2006-01-02T15:04" time_utc:"true"`
}

type Product struct {
	ID       string  `form:"id,omitempty"`
	Quantity int     `form:"quantity,string"`
	Name     string  `form:"name"`
	Price    float64 `form:"price" json:"price,string"`
	Disabled bool
}

type Item struct {
	ID        string
	OrderID   string
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity,string"`
	Product
}
