package service

type URLShortenerService struct {
}

type URLShortener interface {
	Shorten(orignalURL string) (string, error)
	ResolveOrignalURL(shortURL string) (string, error)
}

func (u *URLShortenerService) Shorten(orignalURL string) (string, error) {

	return "", nil
}

func (u *URLShortenerService) ResolveOrignalURL(shortURL string) (string, error) {

	return "", nil
}

func NewURLService() URLShortener {

	return &URLShortenerService{}
}
