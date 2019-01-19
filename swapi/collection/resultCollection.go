package collection

type ResultCollection struct {
	count    int    `json:"count"`
	previous string `json:"previous"`
	next     string `json:"next"`
}
