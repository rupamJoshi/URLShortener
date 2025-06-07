package service

import (
	"crypto/rand"
	"errors"
	"fmt"

	"math/big"

	"short_url.com/config"
)

var shortURLMap = make(map[string]*ShortURL)

type URLShortenerService struct {
	config *config.Config
}

type URLShortener interface {
	Shorten(orignalURL string) (string, error)
	ResolveOrignalURL(shortURL string) (string, error)
	Analytics(shortURL string) (int, error)
}

func generateRandomString(cfg *config.Config) (string, error) {
	length := cfg.ShortURLConf.Length
	charset := cfg.ShortURLConf.Charset

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

	short, err := generateRandomString(u.config)
	if err != nil {
		return "", err
	}
	fmt.Println(short, u.config.ShortURLConf.Length)
	shortURL := fmt.Sprintf("%s://%s:%s/%s", u.config.Schema, u.config.Host, u.config.Port, short)
	fmt.Println(shortURL)
	shortURLMap[short] = &ShortURL{
		ID:           short,
		ShortURL:     shortURL,
		OrignalURL:   orignalURL,
		ResolveCount: 1,
	}

	return shortURL, nil
}

func (u *URLShortenerService) ResolveOrignalURL(shortURL string) (string, error) {

	s := shortURLMap[shortURL]
	if s == nil {
		return "", errors.New("URL not found")
	}

	orignalURL := s.OrignalURL

	s.ResolveCount = s.ResolveCount + 1

	shortURLMap[shortURL] = s

	return orignalURL, nil
}

func (u *URLShortenerService) Analytics(shortURL string) (int, error) {

	s := shortURLMap[shortURL]
	if s == nil {
		return 0, nil
	}

	return s.ResolveCount, nil

}

func NewURLService(config *config.Config) URLShortener {

	return &URLShortenerService{
		config: config,
	}
}
