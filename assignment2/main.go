package main

import (
	"fmt"
)

var itemSlice = []Item{}

type Checkout interface {
	Scan(item Item)
	CalculateMyOrderTotal()
}

func Scan(item Item) []Item {
	itemSlice = append(itemSlice, item)
	return itemSlice
}

func CalculateMyOrderTotal() float64 {
	totalSum := 0.0
	for _, v := range itemSlice {
		fmt.Println(itemSlice)
		totalSum += v.itemPrice
	}
	return totalSum
}

func PricingService(listItem []Item) {
	for _, v := range listItem {
		if v.itemName == "A" && v.itemQuantity > 0 {
			var nonDisPrice float64
			quan := int64(float64(v.itemQuantity) / float64(3))
			offPrice := float64(quan * 130)
			nonDiscountItem := v.itemQuantity % 3
			if nonDiscountItem != 0 {
				nonDisPrice = float64(nonDiscountItem * 50)

			}
			v.itemPrice = offPrice + nonDisPrice

			fmt.Println(v.itemQuantity, v.itemPrice)

		} else if v.itemName == "B" {

		} else if v.itemName == "C" {

		} else if v.itemName == "D" {

		}
	}
	fmt.Println(listItem)

}

type Item struct {
	itemName     string
	itemPrice    float64
	itemQuantity int
}

func main() {
	item1 := &Item{
		"A",
		50,
		5,
	}
	Scan(*item1)
	item2 := &Item{
		"B",
		30,
		2,
	}
	Scan(*item2)
	item3 := &Item{
		"C",
		20,
		5,
	}
	Scan(*item3)

	PricingService(itemSlice)
	fmt.Println(CalculateMyOrderTotal())

}
