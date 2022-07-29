package main

import "testing"

func TestMain(t *testing.T) {
	// Covers edge case of two items in the order with promotion buy2get1free. The item with bigger price is chosen for the voucher.
	t.Run("Test Buy2Get1Free Item", func(t *testing.T) {
		order := &Order{
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
	t.Run("Test 50% Off", func(t *testing.T) {
		order := &Order{
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
	t.Run("Test Buy1 Get Next 1 Baht", func(t *testing.T) {
		order := &Order{
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
	t.Run("Test Hundred Bhatt Discount", func(t *testing.T) {
		order := &Order{
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
		// The discount should be 999 Baht because SKU "B" has one unit of 1000 baht and second becomes 1 baht leading to a discount of 999 Baht
		if order.Discount != 100 {
			t.Errorf("Expected discount to be 999, got %f", order.Discount)
		}
	})
	// TODO: Edge case where order is Valid selected item as well as free item.
	t.Run("Test Buy 2 Get 1 Item Free", func(t *testing.T) {
		order := &Order{
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
	// For this prmotion to work, there needs to be one item and another item that has the field ValidFiftyOff
	t.Run("Test Buy 1 Get 1 Half Price", func(t *testing.T) {
		order := &Order{
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
	// Edge case of Increasing discount where the discount should be bigger than 1000 but due to the limit, it is capped at 1000
	t.Run("Test Increasing Discount(Total >5000)", func(t *testing.T) {
		order := &Order{
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
		if order.Discount != 1000 {
			t.Errorf("Expected 1000, got %f", order.Discount)
		}
	})

}
