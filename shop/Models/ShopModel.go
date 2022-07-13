package Models

type Product struct {
	ProductId uint   `json:"ProductId" gorm:"primary_key"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	Quantity  uint   `json:"quantity"`
}
type Customer struct {
	CustomerId uint   `json:"CustomerId" gorm:"primary_key"`
	Name       string `json:"name"`
	Age        uint   `json:"age"`
	LastOrder  string `json:"LastOrder"`
}

type Order struct {
	OrderId    uint   `json:"OrderId" gorm:"primary_key"`
	CustomerId uint   `json:"CustomerId"`
	ProductId  uint   `json:"ProductId"`
	Quantity   uint   `json:"quantity"`
	Status     string `json:"status"`
}
