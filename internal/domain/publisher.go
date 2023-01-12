package domain

type Publisher struct {
	ID        string `json:"id"`
	Domain    string `json:"domain"`
	IPAddress string `json:"ip_address"`
}
