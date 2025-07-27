package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"valid header": {
			input: http.Header{"Authorization": []string{"ApiKey apikey123"}},
			want:  "apikey123",
			err:   nil,
		},
		"empty header": {
			input: http.Header{"Authorization": []string{}},
			want:  "",
			err:   errors.New("no authorization header included"),
		},
		"malformed header": {
			input: http.Header{"Authorization": []string{"Bearer bearerkey"}},
			want:  "",
			err:   errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			got, err := GetAPIKey(tc.input)

			if got != tc.want {
				t.Errorf("%s: got %v, want %v", name, got, tc.want)
			}

			if err == nil && tc.err != nil ||
				err != nil && tc.err == nil ||
				err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("%s: got %v, want %v", name, err, tc.err)
			}
		})
	}
}
