package entity

type ProductStatus struct {
	Relayable   bool //represents product data is being fetched and reliable to other ser
	Stored      bool //represents product data is being stored to trade data repository
	Transmitted bool //represents product data is being transmitted to other services realtime
}

func (s *ProductStatus) IsOk() bool {

	var StatusInInt int8

	if s.Relayable {
		StatusInInt = StatusInInt | 1<<2
	} else if s.Stored {
		StatusInInt = StatusInInt | 1<<1
	} else if s.Transmitted {
		StatusInInt = StatusInInt | 1<<0
	}

	if 0b001 <= StatusInInt && StatusInInt <= 0b011 {
		return false
	} else {
		return true
	}
}
