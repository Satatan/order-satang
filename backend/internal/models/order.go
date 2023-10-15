package models

type Order struct {
	ID       uint64
	UserID   string
	Product  string
	Price    int
	Quantity int
}
