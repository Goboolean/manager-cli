package entity

type ProductStatus struct {
	Relayable   bool //represents product data is being fetched and reliable to other ser
	Stored      bool //represents product data is being stored to trade data repository
	Transmitted bool //represents product data is being transmitted to other services realtime
}
