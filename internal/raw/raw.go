package raw

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

type Loader interface {
	LOad() ([]Stock, error)
}

type Filterer interface {
	Filter([]Stock) []Stock
}
