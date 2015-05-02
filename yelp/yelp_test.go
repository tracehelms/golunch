package yelp

import "testing"

func TestSearch(t *testing.T) {
	want := "yay"
	got := New().Search("burgers")
	if "boo" != want {
		t.Errorf("Expected: %q ; got: %q", want, got)
	}
}
