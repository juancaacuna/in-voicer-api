package invoicer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCreateInvoices(t *testing.T) {
	expected := []Invoice{Invoice{}, Invoice{}}
	result := []Invoice{Invoice{}, Invoice{}}
	CreateInvoices(&result)
	assert.True(t, reflect.DeepEqual(result, expected), "Should parse an array of Invoices.")
}

func TestGenerateAutoCalculations(t *testing.T) {
	expected := Invoice{
		Items: []Item{
			Item{ Quantity: 2, UnitCost: 20 },
			Item{ Quantity: 1, UnitCost: 30 },
		},
		AmountPaid: 20,
		Tax: 10,
		Discount: 5,
		Fields: Field{
			Tax: "%",
			Discount: "$",
		},
		Calculations: Calculation{
			Subtotal: 70,
			Tax: 7,
			Discount: -5,
			Total: 72,
			Balance: 52,
		},
	}
	result := Invoice{
		Items: []Item{
			Item{ Quantity: 2, UnitCost: 20 },
			Item{ Quantity: 1, UnitCost: 30 },
		},
		AmountPaid: 20,
		Tax: 10,
		Discount: 5,
		Fields: Field{
			Tax: "%",
			Discount: "$",
		},
	}
	generateAutoCalculations(&result)
	assert.True(t, reflect.DeepEqual(result, expected), "Should generate all calculations for invoice.")
}

func TestCalculateSubTotal(t *testing.T) {
	expected := Invoice{
		Tax: 10,
		Discount: 5,
		Fields: Field{
			Tax: "%",
			Discount: "$",
		},
		Calculations: Calculation{
			Subtotal: 100,
			Tax: 10,
			Discount: -5,
			Total: 105,
		},
	}
	result := Invoice{
		Tax: 10,
		Discount: 5,
		Fields: Field{
			Tax: "%",
			Discount: "$",
		},
		Calculations: Calculation{
			Subtotal: 100,
			Tax: 0,
			Discount: 0,
			Total: 100,
		},
	}
	calculateSubTotal(&result, result.Tax, result.Fields.Tax, &result.Calculations.Tax, false)
	calculateSubTotal(&result, result.Discount, result.Fields.Discount, &result.Calculations.Discount, true)
	assert.True(t, reflect.DeepEqual(result, expected), "Should calculate subtotals for invoice.")
}