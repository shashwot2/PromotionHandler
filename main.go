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

//Please note that item C isn't "Added" but discount is included for item C. The promotion isn't valid if item C isn't present.
// There should also be two seperate items of A and B. Two of A doesn't satisfy the condition of this promotion.
type Item struct {
	SKU               string
	Price             float64
	Amount            int64
	ValidSelectedItem bool // For determining if the particular item is applicable for Buy A,B get C added for free,
	ValidFreeItem     bool // For determining if this particular item can be added to order for free in the promotion
	ValidFiftyOff     bool // For determining if this particular SKU is selected for 50% off
}

// This design relies on PromID calling the methods of the Promotion struct. These are like the voucher codes used in the store.
// If Certain PromID's are included in the Object, the methods will be carried out when calculating The maximum discount
// Inorder to keep it efficient, Only PromID's applied to the Order struct should be included.
type Promotion struct {
	PromName string
	PromID   string
}

func (order *Order) CalcTotal() {
	var total float64 = 0
	for _, item := range order.Items {
		total += item.Price * float64(item.Amount)
	}
	order.Total = total
}

// Total needs to be calculated before calling this function because it needs order.Total to compute the discount
// In the Edge case of two promotions having the same dicscount, the left will be chosen which means order.Discount will not be changed
func (order *Order) CalcDiscount() {
	// Guard cases where there are 0 items in which case there is always no discount
	if len(order.Promotions) <= 0 || len(order.Items) == 0 {
		order.Discount = 0
		return
	}
	for i := 0; i < len(order.Promotions); i++ {
		switch order.Promotions[i].PromID {
		case "B2G1":
			order.Discount = Max(order.Discount, order.Promotions[i].Buy2Get1Free(*order))
		case "HOFF":
			order.Discount = Max(order.Discount, order.Promotions[i].C50Off(*order))
		case "B1N1":
			order.Discount = Max(order.Discount, order.Promotions[i].Buy1N1B(*order))
		case "D100":
			order.Discount = Max(order.Discount, order.Promotions[i].C100Baht(*order))
		case "B2I1":
			order.Discount = Max(order.Discount, order.Promotions[i].BuyABFreeC(*order))
		case "B1NH":
			order.Discount = Max(order.Discount, order.Promotions[i].Buy1NextHalf(*order))
		case "INCD":
			order.Discount = Max(order.Discount, order.Promotions[i].DInc30(*order))
		}

	}
}

// Basic Max comparison function for making the code easier to read
func Max(leftN, rightN float64) float64 {
	if leftN == rightN {
		return leftN
	}
	if leftN > rightN {
		return leftN
	}
	return rightN
}
func (order *Order) Print() {
	fmt.Printf("Order ID: %s\n", order.ID)
	fmt.Printf("Total: %.2f\n", order.Total)
	fmt.Printf("Discount: %.2f\n", order.Discount)
	fmt.Printf("Total Payable: %.2f\n", order.Total-order.Discount)
}

// This function addresses edge case of two items in the order with buy2get1free. The higher item with bigger price is chosen for buy2get1free
func (prom Promotion) Buy2Get1Free(Order Order) float64 {
	var maxamount float64 = 0
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].Amount >= 2 && maxamount < Order.Items[i].Price {
			maxamount = Order.Items[i].Price
		}
	}
	return maxamount
}

// Multiplication is faster than division
func (prom Promotion) C50Off(Order Order) float64 {
	return Order.Total * 0.5
}
func (prom Promotion) Buy1N1B(Order Order) float64 {
	var HighestDiscount float64 = 0
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].Amount > 1 && Order.Items[i].Price > HighestDiscount {
			HighestDiscount = Order.Items[i].Price - 1
		}
	}
	return HighestDiscount
}
func (prom Promotion) C100Baht(Order Order) float64 {
	if Order.Total >= 1000 {
		return 100
	} else {
		return 0
	}
}

// The Highest Discounted FreeItem is added as Discount
// The first Loop checks if the items that are selected for this particular promotion is greater than two
func (prom Promotion) BuyABFreeC(Order Order) float64 {
	if len(Order.Items) < 2 {
		return 0
	}
	var SelectedItems float64 = 0
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].ValidSelectedItem {
			SelectedItems++
		}
	}
	if SelectedItems < 2 {
		return 0
	}
	var MaxFreeitem float64 = 0
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].ValidFreeItem && Order.Items[i].Price > MaxFreeitem {
			MaxFreeitem = Order.Items[i].Price
		}
	}
	return MaxFreeitem
}

// The temporary array is for checking all the items that are applicable to being 50% off and only applying the Half price on the greatest item prioritizing high discount
func (prom Promotion) Buy1NextHalf(Order Order) float64 {
	if len(Order.Items) < 1 {
		return 0
	}
	var MaxFiftyoff float64 = 0
	for i := 0; i < len(Order.Items); i++ {
		if Order.Items[i].ValidFiftyOff && Order.Items[i].Price*0.5 > MaxFiftyoff {
			MaxFiftyoff = Order.Items[i].Price * 0.5
		}
	}
	return MaxFiftyoff
}

func (prom Promotion) DInc30(Order Order) float64 {
	var totalItems int64 = 0
	for i := 0; i < len(Order.Items); i++ {
		totalItems += Order.Items[i].Amount
	}
	var discount float64
	switch {
	case totalItems == 1:
		{
			discount = Order.Total * 0.15
			break
		}
	case totalItems == 2:
		{
			discount = Order.Total * 0.2
			break
		}
	case totalItems >= 3:
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
}
