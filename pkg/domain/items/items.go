package items

// Item is a book representation
type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"picturs"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description contains description for an item
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

// Picture of an item
type Picture struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
