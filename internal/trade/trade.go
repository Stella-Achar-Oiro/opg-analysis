package trade

import (
	"opg/internal/news"
	"opg/internal/pos"
)

type Selection struct {
	Ticker string
	pos.Position
	Articles []news.Article
}

type Deliverer interface {
	Deliver(selections []Selection) error
}