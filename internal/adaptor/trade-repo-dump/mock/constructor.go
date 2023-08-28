package tradeRepoDumpMock

type TradeDumpAdaptorMock struct {
	database string
}

func New() *TradeDumpAdaptorMock {
	return &TradeDumpAdaptorMock{
		database: "goboolean-stock",
	}
}
