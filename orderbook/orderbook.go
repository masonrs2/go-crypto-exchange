package orderbook

import (
	"fmt"
	"time"
)

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

func (limit *Limit) AddOrder(order *Order) {
	order.Limit = limit
	limit.Orders = append(limit.Orders, order)
	limit.TotalVolume += order.Size
}

func (limit *Limit) DeleteOrder(order *Order) {
	for i := 0; i < len(limit.Orders); i++ {
		if limit.Orders[i] == order {
			limit.Orders[i] = limit.Orders[len(limit.Orders)-1]
			limit.Orders = limit.Orders[:len(limit.Orders)-1]
		}
	}

	order.Limit = nil
	limit.TotalVolume -= order.Size
}

func (order *Order) String() string {
	return fmt.Sprintf("[size: %.2f]", order.Size)
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[float64]*Limit),
		BuyLimits: make(map[float64]*Limit),
	}
}

func (ob *Orderbook) PlaceOrder(price float64, order *Order) []Match {
	// Match the orders

	// Adds the rest of the order to the orderbook
	if order.Size > 0.0 {
		ob.Add(price, order)
	}

	return []Match{}
}

func (ob *Orderbook) Add(price float64, order *Order) {
	var limit *Limit

	if order.Bid {
		limit = ob.BuyLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}

	if limit == nil {
		limit = NewLimit(price)
		limit.AddOrder(order)

		if order.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BuyLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
}
