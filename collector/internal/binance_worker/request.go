package binanceworker

type Request struct {
  Method string `json:"method"`
  Params []string `json:"params"`
  ID int `json:"id"`
}
