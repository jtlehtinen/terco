package main

import (
	"testing"
)

func TestGetFieldTag(t *testing.T) {
	theme := new(TerminalTheme)

	want := "brightBlack"
	got := getFieldTag(theme, &theme.BrightBlack, "json")
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
