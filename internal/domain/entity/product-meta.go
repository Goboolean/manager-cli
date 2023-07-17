package entity

// All metadata that this domain needs
type ProductMeta struct {
	Id       string // {type}.{name}.{location}
	Name     string // Human readable product name
	Code     string // Product code used in one's market ex) AAPL(apple inc), 005930(samsung)
	Exchange string // Human readable Product's exchange name  ex) nasdaq, kospi etc...
	Type     string // Type of product ex) stock or encrypt
	Location string // Standard 3 letter country Code fallowing ISO 3166 -1
}
