package orderbook

type Limits []*Limit
type ByBestAsk struct{ Limits }
type Orders []*Order

func (ask ByBestAsk) Len() int           { return len(ask.Limits) }
func (ask ByBestAsk) Swap(i, j int)      { ask.Limits[i], ask.Limits[j] = ask.Limits[j], ask.Limits[i] }
func (ask ByBestAsk) Less(i, j int) bool { return ask.Limits[i].Price < ask.Limits[j].Price }

type ByBestBid struct{ Limits }

func (bid ByBestBid) Len() int           { return len(bid.Limits) }
func (bid ByBestBid) Swap(i, j int)      { bid.Limits[i], bid.Limits[j] = bid.Limits[j], bid.Limits[i] }
func (bid ByBestBid) Less(i, j int) bool { return bid.Limits[i].Price > bid.Limits[j].Price }

func (orders Orders) Len() int           { return len(orders) }
func (orders Orders) Swap(i, j int)      { orders[i], orders[j] = orders[j], orders[i] }
func (orders Orders) Less(i, j int) bool { return orders[i].Timestamp > orders[j].Timestamp }