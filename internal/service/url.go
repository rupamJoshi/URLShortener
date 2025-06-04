package service

import (
	"crypto/rand"
	"fmt"

	"math/big"

	"short_url.com/config"
)

var shortURLMap = make(map[string]ShortURL)

type URLShortenerService struct {
	config *config.Config
}

type URLShortener interface {
	Shorten(orignalURL string) (string, error)
	ResolveOrignalURL(shortURL string) (string, error)
}

func generateRandomString() (string, error) {
	length := 6
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randInt, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		result[i] = charset[randInt.Int64()]
	}
	return string(result), nil
}

func (u *URLShortenerService) Shorten(orignalURL string) (string, error) {

	short, err := generateRandomString()
	if err != nil {
		return "", err
	}

	shortURL := fmt.Sprintf("%s://%s:%s/%s", u.config.Schema, u.config.Host, u.config.Port, short)
	shortURLMap[short] = ShortURL{
		ID:           short,
		ShortURL:     shortURL,
		OrignalURL:   orignalURL,
		ResolveCount: 0,
	}

	return shortURL, nil
}

func (u *URLShortenerService) ResolveOrignalURL(shortURL string) (string, error) {

	s := shortURLMap[shortURL]

	orignalURL := s.OrignalURL

	s.ResolveCount = s.ResolveCount + 1

	shortURLMap[shortURL] = s

	return orignalURL, nil
}

func NewURLService(config *config.Config) URLShortener {

	return &URLShortenerService{
		config: config,
	}
}
