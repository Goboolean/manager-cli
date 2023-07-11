package entity

type ProductStatus struct {
	Reliable    bool // Value representing product data is being fetched and reliable to other services
	Stored      bool // Value representing product data is being stored to trade data repository
	Transmitted bool // Value representing product data is being transmitted to other services realtime
}
