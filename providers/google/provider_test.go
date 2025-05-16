package google_test

import (
	"testing"

	"translate"

	"translate/providers/google"

	"github.com/stretchr/testify/require"
)

func TestTranslation_Correct(t *testing.T) {
	text, err := google.New().Translate(t.Context(), "en", "ru", "Hello World!")

	require.NoError(t, err)
	require.Equal(t, "Привет, мир!", text)
}

func TestTranslation_Auto(t *testing.T) {
	text, err := google.New().Translate(t.Context(), "auto", "ru", "Hello World!")

	require.NoError(t, err)
	require.Equal(t, "Привет, мир!", text)
}

func TestTranslation_NoLang(t *testing.T) {
	text, err := google.New().Translate(t.Context(), "ki", "ru", "Hello World!")

	require.ErrorIs(t, err, translate.ErrUnsupportedLang)
	require.Empty(t, text)
}
