package main

import (
	"testing"
)

func TestBuy2Get1Free(t *testing.T) {
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "123",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 30, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 20, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "D", Price: 15, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: true},
			},
			Promotions: []Promotion{
				{PromName: "Buy 2Get1Free Item", PromID: "B2G1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// Since SKU B, SKU C has 30 Baht with 2 Amount, the discount is 30 Baht.
		if order.Discount != 30 {
			t.Errorf("Expected discount to be 30, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID: "123",
			Items: []Item{
				{SKU: "A", Price: 15000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy 2Get1Free Item", PromID: "B2G1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The Promotion is not applicable for this order because there is only 1 item.
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})
}
func Test50percentoff(t *testing.T) {
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 3000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "50% Off", PromID: "HOFF"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The total should be 14000 Baht and the discount should be 7000 Baht.
		if order.Discount != 7000 {
			t.Errorf("Expected discount to be 7000, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID:    "1",
			Items: []Item{},
			Promotions: []Promotion{
				{PromName: "50% Off", PromID: "HOFF"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// This test case has no items so there should be no discount.
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})
}

func TestBuy1GetNext1Baht(t *testing.T) {
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 1000, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy1 Get Next 1 Baht", PromID: "B1N1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 999 Baht because SKU "B" has one unit of 1000 baht and second becomes 1 baht leading to a discount of 999 Baht
		if order.Discount != 999 {
			t.Errorf("Expected discount to be 7000, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy1 Get Next 1 Baht", PromID: "B1N1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should is invalid because there is only 1 item which is not applicable to the promotion.
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})
}

func TestHundredBahtDiscount(t *testing.T) {
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 1000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 1000, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 2000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Hundred Bhatt Discount", PromID: "D100"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 100 Baht because total is 5000 bhatt and the discount is 100 bhatt
		if order.Discount != 100 {
			t.Errorf("Expected discount to be 999, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 400, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 100, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 200, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Hundred Bhatt Discount", PromID: "D100"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 0 Baht because the total of the order is 800 and does not exceed 1000 baht
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})

}

func TestBuyABGetC(t *testing.T) {
	// TODO: Edge case where order is Valid selected item as well as free item.
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: true, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 3000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: true, ValidFiftyOff: false},
				{SKU: "C", Price: 2000, Amount: 2, ValidSelectedItem: true, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy A, B Get C Free", PromID: "B2I1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 3000 because SKU "B" which is a valid free item is present in the order.
		if order.Discount != 3000 {
			t.Errorf("Expected discount to be 3000, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: true, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 3000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: true, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy A, B Get C Free", PromID: "B2I1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 0 because there aren't two Valid Selected items(item A and B) in the order
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})
	t.Run("Edge Case: A, B items are in order but not C", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: true, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 3000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 2000, Amount: 2, ValidSelectedItem: true, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy A, B Get C Free", PromID: "B2I1"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 0 because there aren't isn't a item(C) in this order which can be used for the promotion, item C needs (ValidFreeItem set to true)
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})

}
func TestBuy1Get1Half(t *testing.T) {
	// For this prmotion to work, there needs to be one item and another item that has the field ValidFiftyOff
	t.Run("Applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 5000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 3000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: true},
			},
			Promotions: []Promotion{
				{PromName: "Buy 1 Get 1 Half Price", PromID: "B1NH"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 1500 because "B is a valid 50% off item and there is one other item which is A, or other instances of B itself"."
		if order.Discount != 1500 {
			t.Errorf("Expected discount to be 1500, got %f", order.Discount)
		}
	})
	t.Run("Not applicable", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 3000, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 1000, Amount: 3, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy 1 Get 1 Half Price", PromID: "B1NH"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The discount should be 0 because there isn't any item that has the validFiftyoff field set to true
		if order.Discount != 0 {
			t.Errorf("Expected discount to be 0, got %f", order.Discount)
		}
	})

}

func TestIncreasingDiscount(t *testing.T) {
	t.Run("Applicable(30% Discount)", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{
					SKU:               "A",
					Price:             5000,
					Amount:            50,
					ValidSelectedItem: false,
					ValidFreeItem:     false,
					ValidFiftyOff:     false,
				},
			},
			Promotions: []Promotion{
				{PromName: "1 15%, 2 20%, 3 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The Discount here should be 75000 because the total is 250000 and the discount is 30% but since its limited to 1000, the discount is 1000
		if order.Discount != 1000 {
			t.Errorf("Expected 1000, got %f", order.Discount)
		}
	})
	t.Run("Applicable(20% Discount)", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 500, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 300, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "1 15%, 2 20%, 3 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The total items are two so theres a 20% discount which is 160 when the total is 800
		if order.Discount != 160 {
			t.Errorf("Expected 160, got %f", order.Discount)
		}
	})
	t.Run("Applicable(15% Discount)", func(t *testing.T) {
		order := Order{
			ID: "1",
			Items: []Item{
				{SKU: "A", Price: 500, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "1 15%, 2 20%, 3 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// The total items are one so there's a 15% discount which is 75 when the total is 500
		if order.Discount != 75 {
			t.Errorf("Expected 75, got %f", order.Discount)
		}
	})
	t.Run("Not Applicable", func(t *testing.T) {
		order := Order{
			ID:    "1",
			Items: []Item{},
			Promotions: []Promotion{
				{PromName: "1 15%, 2 20%, 3 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// There are no items so the discount should be 0
		if order.Discount != 0 {
			t.Errorf("Expected 0, got %f", order.Discount)
		}
	})
}

func TestMultiplePromotions(t *testing.T) {
	// Testing for multiple promotions in a single order
	t.Run("Half Discount is highest promotion among 3", func(t *testing.T) {
		order := Order{
			ID: "5",
			Items: []Item{
				{SKU: "A", Price: 500, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 30.5, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 20.78, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "D", Price: 15.12, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy2Get1Free", PromID: "B2G1"},
				{PromName: "Discount 50% off", PromID: "HOFF"},
				{PromName: "Buy 1 get 15%, Buy 2 get 20%, Buy 3 get 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// Since the order has multiple discounts , the second discount which is 50% is the maximum discount hence it is chosen
		if order.Discount != 283.2 {
			t.Errorf("Expected discount to be 283.2, got %f", order.Discount)
		}
	})
	t.Run("Increasing Discount is highest promotion among 3", func(t *testing.T) {
		order := Order{
			ID: "5",
			Items: []Item{
				{SKU: "A", Price: 500, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "B", Price: 30.5, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "C", Price: 20.78, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
				{SKU: "D", Price: 15.12, Amount: 1, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy2Get1Free", PromID: "B2G1"},
				{PromName: "Buy 1 get 15%, Buy 2 get 20%, Buy 3 get 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// Since the order has multiple discounts , the increasing discount is chosen. the item has more than 3 times so the discount is at a maximum of 30%
		if order.Discount != 169.92 {
			t.Errorf("Expected discount to be 169.92, got %f", order.Discount)
		}
	})
	t.Run("Edge case: Two promotions have the same amount of discount", func(t *testing.T) {
		order := Order{
			ID: "5",
			Items: []Item{
				{SKU: "A", Price: 600, Amount: 2, ValidSelectedItem: false, ValidFreeItem: false, ValidFiftyOff: false},
			},
			Promotions: []Promotion{
				{PromName: "Buy2Get1Free", PromID: "B2G1"},
				{PromName: "Discount 50% off", PromID: "HOFF"},
				{PromName: "Buy 1 get 15%, Buy 2 get 20%, Buy 3 get 30%", PromID: "INCD"},
			},
		}
		order.CalcTotal()
		order.CalcDiscount()
		// Since the order has multiple discounts with same amount of discount, the left will be chosen. Discount will always remain the same in this case
		if order.Discount != 600 {
			t.Errorf("Expected discount to be 600, got %f", order.Discount)
		}
	})
}
