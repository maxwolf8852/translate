package translate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/maxwolf8852/translate"

	"github.com/maxwolf8852/translate/providers/google"

	"github.com/stretchr/testify/require"
)

const (
	mockCorrectInput  = "correct"
	mockCorrectOutput = "translated"
)

var mockError = errors.New("mock error")

type mockProvider struct{}

func (p *mockProvider) Translate(_ context.Context, _, _ translate.Lang, text string) (string, error) {
	if text == mockCorrectInput {
		return mockCorrectOutput, nil
	}
	return "", mockError
}

type errorProvider struct{}

func (p *errorProvider) Translate(_ context.Context, _, _ translate.Lang, text string) (string, error) {
	return "", mockError
}

func TestReal_Correct(t *testing.T) {
	client, err := translate.New(translate.WithProvider(google.New()))
	require.NoError(t, err)

	out, err := client.Translate(t.Context(), translate.EN, translate.FR, "Hello world!")
	require.NoError(t, err)
	require.Equal(t, "Bonjour le monde!", out)
}

func TestReal_NoProviders(t *testing.T) {
	client, err := translate.New()
	require.NoError(t, err)

	testError(t, client, "", translate.ErrNoProviders)
}

func TestMock_Correct(t *testing.T) {
	client, err := translate.New(translate.WithProvider(&mockProvider{}))
	require.NoError(t, err)

	testCorrect(t, client, mockCorrectInput)
}

func TestMock_Incorrect(t *testing.T) {
	client, err := translate.New(translate.WithProvider(&mockProvider{}))
	require.NoError(t, err)

	testError(t, client, "", mockError)
}

func TestMock_SeveralProvidersSkipErrors(t *testing.T) {
	t.Run("skip errors", func(t *testing.T) {
		client, err := translate.New(
			translate.WithSkipErrors(),
			translate.WithProvider(&errorProvider{}),
			translate.WithProvider(&mockProvider{}))
		require.NoError(t, err)
		testCorrect(t, client, mockCorrectInput)
	})

	t.Run("not skip errors", func(t *testing.T) {
		client, err := translate.New(translate.WithProvider(&errorProvider{}), translate.WithProvider(&mockProvider{}))
		require.NoError(t, err)
		testError(t, client, mockCorrectInput, mockError)
	})
}

func testError(t *testing.T, client *translate.Client, input string, targetErr error) {
	t.Helper()

	out, err := client.Translate(t.Context(), "", "", input)
	require.ErrorIs(t, err, targetErr)
	require.Empty(t, out)
}

func testCorrect(t *testing.T, client *translate.Client, input string) {
	t.Helper()

	out, err := client.Translate(t.Context(), "", "", input)
	require.NoError(t, err)
	require.Equal(t, mockCorrectOutput, out)
}
