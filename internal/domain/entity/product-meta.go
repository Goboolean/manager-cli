package entity

// All metadata that this domain needs
type ProductMeta struct {
	Name     string // Human readable product name
	Code     string // Product code used in one's market ex) AAPL(apple inc), 005930(samsung)
	Exchange string // Human readable Product's exchange name  ex) nasdaq, kospi etc...

}
