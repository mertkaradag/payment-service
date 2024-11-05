package bank

// Akbank is a bank
type Akbank struct {
	Commission float64
}

// FetchCommission returns the commission of Akbank
func (a *Akbank) FetchCommission() float64 {
	return a.Commission
}

// FetchName returns the name of Akbank
func (a *Akbank) FetchName() string {
	return "AKBANK"
}
