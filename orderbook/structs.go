package orderbook

type Order struct {
	Size      float64
	Limit     *Limit
	Timestamp int64
	Bid       bool
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map[float64]*Limit
	BuyLimits map[float64]*Limit
}

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}