package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GETAPIKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		header           http.Header
		expectedResponse string
		expectError      bool
	}{
		{
			name: "valid api key",
			header: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			expectedResponse: "1234567890",
			expectError:      false,
		},
		{
			name: "not valid api key",
			header: http.Header{
				"Authorization": []string{"ApiKey1234567890"},
			},
			expectedResponse: "malformed authorization header",
			expectError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			apiKey, err := GetAPIKey(tt.header)
			if tt.expectError {
				assert.Equal(t, apiKey, "")
				assert.Equal(t, tt.expectedResponse, err.Error())
			} else {
				assert.Equal(t, tt.expectedResponse, apiKey)
				assert.Equal(t, err, nil)
			}
		})
	}
}
