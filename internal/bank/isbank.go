package bank

// Isbank is a bank
type Isbank struct {
	Commission float64
}

// FetchCommission returns the commission of Isbank
func (i *Isbank) FetchCommission() float64 {
	return i.Commission
}

// FetchName returns the name of Isbank
func (i *Isbank) FetchName() string {
	return "ISBANK"
}
