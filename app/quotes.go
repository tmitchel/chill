package app

import (
	"math/rand"
	"time"
)

// Quote represents a single quote including the author.
type Quote struct {
	ID      int
	Author  string
	Content string
}

// Quotes provides a type for a slice of Quote's.
type Quotes []*Quote

// BuildQuotes converts the map of quotes to a slice of Quotes
func BuildQuotes() *Quotes {
	rand.Seed(time.Now().UnixNano())
	qts := make(Quotes, len(quotes))
	i := 0
	for quote, author := range quotes {
		qts[i] = &Quote{
			ID:      i,
			Author:  author,
			Content: quote,
		}
		i++
	}
	return &qts
}

// RandomQuote returns a random quote.
func (q *Quotes) RandomQuote() *Quote {
	return (*q)[rand.Intn(len(*q))]
}

var quotes map[string]string = map[string]string{
	"Be yourself; everyone else is already taken.":                                                                         "Oscar Wilde",
	"In three words I can sum up everything I've learned about life: it goes on.":                                          "Robert Frost",
	"Don’t walk in front of me… I may not follow. Don’t walk behind me… I may not lead. Walk beside me… just be my friend": "Albert Camus",
}
