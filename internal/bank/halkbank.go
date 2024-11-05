package bank

// Halkbank is a bank
type Halkbank struct {
	Commission float64
}

// FetchCommission returns the commission of Halkbank
func (h *Halkbank) FetchCommission() float64 {
	return h.Commission
}

// FetchName returns the name of Halkbank
func (h *Halkbank) FetchName() string {
	return "HALKBANK"
}
