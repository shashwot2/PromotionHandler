package main

import (
	"fmt"
)

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

// This design relies on PromID calling the methods of the Promotion struct.
// If Certain PromID's are included in the Object, the methods will be carried out when calculating The maximum discount
// Inorder to keep it efficient, Only PromID's applied to the Order struct should be included.
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
			{PromName: "Buy 2Get1Free Item", PromID: "B2G1"},
			{PromName: "50% Off", PromID: "HOFF"},
			{PromName: "Buy1 Next1 Baht", PromID: "B1N1"},
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

func (order *Order) Print() {
	fmt.Printf("Order ID: %s\n", order.ID)
	fmt.Printf("Total: %.2f\n", order.Total)
	fmt.Printf("Discount: %.2f\n", order.Discount)
	fmt.Printf("Total Payable: %.2f\n", order.Total-order.Discount)
}
func (prom Promotion) CalcFreeItem(Order Order) float64 {
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].Amount >= 2 {
			return Order.Items[i].Price
		}
	}
	return 0
}

func (prom Promotion) Calc50Off(Order Order) float64 {
	return Order.Total * 0.5
}
func (prom Promotion) CalcBuy1Next1Baht(Order Order) float64 {
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].Amount > 1 {
			return Order.Items[i].Price - 1
		}
	}
	return 0
}
func (prom Promotion) CalcHundredBhatDiscount(Order Order) float64 {
	if Order.Total >= 1000 {
		return 100
	} else {
		return 0
	}
}
func (prom Promotion) CalcBuy2Get1ItemFree(Order Order) float64 {
	if len(Order.Items) > 2 {
		for i := 0; i < len(Order.Items); i++ {
			if Order.Items[i].ValidFreeItem {
				return Order.Items[i].Price
			}
		}
		return 0
	}
	return 0
}
func (prom Promotion) CalcBuy1GetHalfPrice(Order Order) float64 {
	var Fiftyoff []float64
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].ValidFiftyOff {
			Fiftyoff = append(Fiftyoff, Order.Items[i].Price)
		}
	}
	if len(Fiftyoff) == 0 {
		return 0
	} else if len(Fiftyoff) == 1 {
		return Fiftyoff[0] / 2
	} else {
		return maxarr(Fiftyoff) / 2
	}
}

func (prom Promotion) CalcIncreaseDiscount(Order Order) float64 {
	var totalItems int64 = 0
	for i := 0; i < len(Order.Items); i++ {
		totalItems += Order.Items[i].Amount
	}
	var discount float64
	switch {
	case totalItems > 1:
		{
			discount = Order.Total * 0.15
			break
		}
	case totalItems > 2:
		{
			discount = Order.Total * 0.2
			break
		}
	case totalItems > 3:
		{
			discount = Order.Total * 0.3
			break
		}
	}
	if discount >= 1000 {
		return 1000
	} else {
		return discount
	}
	return 0
}

func max(leftN, RightN float64) float64 {
	if leftN > RightN {
		return leftN
	}
	return RightN
}
func maxarr(arr []float64) float64 {
	if len(arr) <= 0 {
		return 0
	}
	var max float64
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func (order *Order) CalcDiscount() {
	if len(order.Promotions) == 0 {
		order.Discount = 0
		return
	}
	for i := 0; i < len(order.Promotions); i++ {
		switch order.Promotions[i].PromID {
		case "B2G1":
			order.Discount = max(order.Discount, order.Promotions[i].CalcFreeItem(*order))
		case "HOFF":
			order.Discount = max(order.Discount, order.Promotions[i].Calc50Off(*order))
		case "B1N1":
			order.Discount = max(order.Discount, order.Promotions[i].CalcBuy1Next1Baht(*order))
		case "D100":
			order.Discount = max(order.Discount, order.Promotions[i].CalcHundredBhatDiscount(*order))
		case "B2I1":
			order.Discount = max(order.Discount, order.Promotions[i].CalcBuy2Get1ItemFree(*order))
		case "B1G1":
			order.Discount = max(order.Discount, order.Promotions[i].CalcBuy1GetHalfPrice(*order))
		case "INCD":
			order.Discount = max(order.Discount, order.Promotions[i].CalcIncreaseDiscount(*order))
		}

	}
}
