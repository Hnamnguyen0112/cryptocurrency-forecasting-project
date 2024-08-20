package coinbaseworker

type Request struct {
  Type string `json:"type"`
  ProductIds []string `json:"product_ids"`
  Channels []string `json:"channels"`
}
