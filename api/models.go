package api

type GetShortenedURLRequest struct {
	OrignalURL string `json:"orignal_url"`
}

type GetShortenedURLResponse struct {
	ID       string
	ShortURL string
}

type GetOrignalURLRequest struct {
	ShortURL string `json:"short_url"`
}

type GetOrignalURLResponse struct {
	OrignalURL string
}

type ShortURLAnalytics struct {
	OrignalURL string
	Count      int
}
