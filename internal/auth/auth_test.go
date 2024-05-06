package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "Valid API key",
			headers: http.Header{"Authorization": []string{"ApiKey ABC123"}},
			want:    "ABC123",
			wantErr: nil,
		},
		{
			name:    "No auth headers",
			headers: http.Header{},
			want:    "",
			wantErr: errors.New("no authorization header included"),
		},
		{
			name:    "malformed auth headers",
			headers: http.Header{"Authorization": []string{"Bearer Token"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if got != tc.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
			if (err != nil && tc.wantErr == nil) || (err == nil && tc.wantErr != nil) || (err != nil && tc.wantErr != nil && err.Error() != tc.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
