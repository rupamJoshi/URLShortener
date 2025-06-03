package service

type ShortURL struct {
	ID           string `json:"id"`
	ShortURL     string `json:"short_url"`
	OrignalURL   string `json:"orignal_url"`
	ResolveCount int    `json:"resolve_count"`
}
