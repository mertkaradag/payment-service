package bank

// Bank is the interface that all banks should implement
type Bank interface {
	FetchCommission() float64
	FetchName() string
}
