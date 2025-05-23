package mymemory_test

import (
	"testing"

	"github.com/maxwolf8852/translate/providers/mymemory"

	"github.com/stretchr/testify/require"
)

func TestTranslation_Correct(t *testing.T) {
	text, err := mymemory.New().Translate(t.Context(), "en", "ru", "Hello World!")

	require.NoError(t, err)
	require.Equal(t, "Привет мир!", text)
}
