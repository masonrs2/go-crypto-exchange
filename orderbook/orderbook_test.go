package orderbook

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	limit := NewLimit(10_000)
	buyOrder1 := NewOrder(true, 5)
	buyOrder2 := NewOrder(true, 8)
	buyOrder3 := NewOrder(true, 10)
	
	limit.AddOrder(buyOrder1)
	limit.AddOrder(buyOrder2)
	limit.AddOrder(buyOrder3)

	limit.DeleteOrder(buyOrder2)

	fmt.Println(limit)
}

func TestOrderbook(t *testing.T) {
	orderbook := NewOrderbook() 

	buyOrder1 := NewOrder(true, 10) 
	buyOrder2 := NewOrder(true, 2000) 

	orderbook.PlaceOrder(18_000, buyOrder1)
	orderbook.PlaceOrder(19_000, buyOrder2)
	
	fmt.Printf("%+v\n", orderbook)
}