package service

import (
	"testing"

	"short_url.com/config"
)

func Test_generateRandomString(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "Test: Should Generate a Random string",
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomString(&config.Config{ShortURLConf: config.ShortURLConf{
				Length:  6,
				Charset: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			}})
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("generateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLShortenerService_Shorten(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		orignalURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test: Should generate a short URL",
			fields: fields{
				config: &config.Config{
					Host:   "",
					Port:   "",
					Schema: "",
				},
			},
			args: args{
				orignalURL: "foo",
			},
			want:    11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLShortenerService{
				config: tt.fields.config,
			}
			got, err := u.Shorten(tt.args.orignalURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortenerService.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("URLShortenerService.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLShortenerService_ResolveOrignalURL(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		shortURL string
	}
	shortURLMap["foo"] = &ShortURL{
		OrignalURL: "bar",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test: should resolve and return a orignal URL",
			fields:  fields{},
			args:    args{shortURL: "foo"},
			want:    "bar",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLShortenerService{
				config: tt.fields.config,
			}
			got, err := u.ResolveOrignalURL(tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortenerService.ResolveOrignalURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLShortenerService.ResolveOrignalURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
