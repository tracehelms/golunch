package yelp

import "testing"

func TestSearch(t *testing.T) {
	want := "yay"
	got := New().Search("hamburger")
	if got != want {
		t.Errorf("Expected: %q ; got: %q", want, got)
	}
}
