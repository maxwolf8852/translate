package linguee_test

import (
	"testing"

	"translate"

	"translate/providers/linguee"

	"github.com/stretchr/testify/require"
)

func TestTranslation_Correct(t *testing.T) {
	text, err := linguee.New().Translate(t.Context(), translate.EN, translate.DE, "Hello World!")

	require.NoError(t, err)
	require.Equal(t, "Hallo Welt!", text)
}

func TestTranslation_NoLang(t *testing.T) {
	text, err := linguee.New().Translate(t.Context(), "auto", translate.DE, "Hello World!")

	require.ErrorIs(t, err, translate.ErrUnsupportedLang)
	require.Empty(t, text)
}
