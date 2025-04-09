package mymemory_test

import (
	"testing"

	"translate/providers/mymemory"
)

func TestTranslation_Correct(t *testing.T) {
	text, err := mymemory.New().Translate(t.Context(), "en", "ru", "Hello World!")
	if err != nil {
		t.Fatal(err)
	}

	if text != "Привет мир!" {
		t.Fatalf("expected 'Привет', got %s", text)
	}
}
