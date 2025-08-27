package main

import (
	"reflect"
	"testing"
)

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

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []Page
	}{
		{
			name: "order count descending",
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: []Page{
				{Link: "url4", CountTo: 10},
				{Link: "url5", CountTo: 7},
				{Link: "url1", CountTo: 5},
				{Link: "url3", CountTo: 3},
				{Link: "url2", CountTo: 1},
			},
		},
		{
			name: "alphabetize",
			input: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
			expected: []Page{
				{Link: "a", CountTo: 1},
				{Link: "b", CountTo: 1},
				{Link: "c", CountTo: 1},
				{Link: "d", CountTo: 1},
				{Link: "e", CountTo: 1},
			},
		},
		{
			name: "order count then alphabetize",
			input: map[string]int{
				"d": 2,
				"a": 1,
				"e": 3,
				"b": 1,
				"c": 2,
			},
			expected: []Page{
				{Link: "e", CountTo: 3},
				{Link: "c", CountTo: 2},
				{Link: "d", CountTo: 2},
				{Link: "a", CountTo: 1},
				{Link: "b", CountTo: 1},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []Page{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: []Page{},
		},
		{
			name: "one key",
			input: map[string]int{
				"url1": 1,
			},
			expected: []Page{
				{Link: "url1", CountTo: 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected Link: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
