package main

func main() {

}

type CheckoutService interface {
	scan()
	pricing()
	checkout()
}

type Item struct {
	itemName     string
	itemPrice    float32
	itemQuantity int
}

func (i *Item) scan() {

}
