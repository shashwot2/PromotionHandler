package main

import "fmt"

type Order struct {
	ID         string      // Order ID
	Items      []Item      // List of items in the order
	Promotions []Promotion // List of available promotions
	Total      float64     // Total price of the order
	Discount   float64     // Total discount of the order
}

type Item struct {
	SKU           string
	Price         float64
	Amount        int64
	ValidFreeItem bool
	ValidFiftyOff bool
}

type Promotion struct {
	PromName string
	PromID   string
}

func main() {
	order := &Order{
		ID: "123",
		Items: []Item{
			{SKU: "A", Price: 50, Amount: 1},
			{SKU: "B", Price: 30, Amount: 1},
			{SKU: "C", Price: 20, Amount: 1},
			{SKU: "D", Price: 15, Amount: 1},
		},
		Promotions: []Promotion{
			{PromName: "Free Item", PromID: "F"},
			{PromName: "50% Off", PromID: "O"},
		},
	}
	order.CalcTotal()
	order.CalcDiscount()
	order.Print()
}

func (order *Order) CalcTotal() {
	total := float64(0)
	for _, item := range order.Items {
		total += item.Price * float64(item.Amount)
	}
	order.Total = total
}

func (order *Order) CalcDiscount() {
	discount := float64(0)
	for _, prom := range order.Promotions {
		discount += prom.CalcDiscount(order)
	}
	order.Discount = discount
}

func (order *Order) Print() {
	fmt.Printf("Order ID: %s\n", order.ID)
	fmt.Printf("Total: %.2f\n", order.Total)
	fmt.Printf("Discount: %.2f\n", order.Discount)
	fmt.Printf("Total Payable: %.2f\n", order.Total-order.Discount)
}
func (prom Promotion) CalcFreeItem(Order *Order) float64 {
	return 0
}

func (prom Promotion) Calc50Off(Order *Order) float64 {
	return Order.Total * 0.5
}

func (prom Promotion) CalcDiscount(order *Order) float64 {
	switch prom.PromID {
	case "F":
		return prom.CalcFreeItem(order)
	case "O":
		return prom.Calc50Off(order)
	}
}
