package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var shortURLMap = make(map[string]ShortURL)

type URLShortenerService struct {
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

	shortURL, err := generateRandomString()
	if err != nil {
		return "", err
	}

	fmt.Println("Shorten", orignalURL)

	shortURLMap[shortURL] = ShortURL{
		ID:           shortURL,
		ShortURL:     shortURL,
		OrignalURL:   orignalURL,
		ResolveCount: 0,
	}

	return shortURL, nil
}

func (u *URLShortenerService) ResolveOrignalURL(shortURL string) (string, error) {

	s := shortURLMap[shortURL]

	orignalURL := s.OrignalURL

	fmt.Printf("Resolve %s, %+v %+v", orignalURL, s, shortURLMap)

	s.ResolveCount = s.ResolveCount + 1

	shortURLMap[shortURL] = s

	return orignalURL, nil
}

func NewURLService() URLShortener {

	return &URLShortenerService{}
}
