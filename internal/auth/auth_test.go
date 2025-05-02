package auth

import (
	"net/http"
	"testing"
)

func CreateHeader(key string, value string) (http.Header, error) {
	req, err := http.NewRequest("POST", "XD.com", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(key, value)
	return req.Header, nil
}

func TestAPIKey(t *testing.T) {
	tests := map[string]struct {
		headerKey   string
		headerValue string
		wantKey     string
		wantErr     error
	}{
		"no auth header": {
			headerKey:   "Content-Type",
			headerValue: "application/json",
			wantKey:     "",
			wantErr:     ErrNoAuthHeaderIncluded,
		},
		"valid auth header": {
			headerKey:   "Authorization",
			headerValue: "ApiKey GardnerMartin",
			wantKey:     "GardnerMartin",
			wantErr:     nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header, err := CreateHeader(tc.headerKey, tc.headerValue)
			if err != nil {
				t.Fatalf("error in creating the request: %v", err)
			}

			gotKey, err := GetAPIKey(header)
			if err != tc.wantErr {
				t.Fatalf("expected error: %v, got: %v", tc.wantErr, err)
			}

			if gotKey != tc.wantKey {
				t.Fatalf("expected API key: %v, got: %v", tc.wantKey, gotKey)
			}
		})
	}
}
