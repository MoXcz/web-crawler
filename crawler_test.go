package main

import "testing"

func Test_comparURLs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rawBasedURL   string
		rawCurrentURL string
		want          bool
		wantErr       bool
	}{
		{
			name:          "different URLs",
			rawBasedURL:   "https/wikipedia.com",
			rawCurrentURL: "https://developer.mozilla.org/en-US/docs/Glossary/State_machine",
			want:          false,
			wantErr:       false,
		},
		{
			name:          "equal URLs",
			rawBasedURL:   "https://en.wikipedia.org/wiki/Linus_Torvalds",
			rawCurrentURL: "https://en.wikipedia.org/wiki/Ken_Thompson",
			want:          true,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := compareHostURLs(tt.rawBasedURL, tt.rawCurrentURL)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("comparURLs() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("comparURLs() succeeded unexpectedly")
			}

			if tt.want != got {
				t.Errorf("comparURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
