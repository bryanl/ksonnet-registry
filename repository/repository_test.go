package repository

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDependencies_ToMap(t *testing.T) {
	deps := Dependencies{
		{Name: "one", Constraint: "1.0"},
		{Name: "two", Constraint: "2.0"},
		{Name: "three", Constraint: "3.0"},
	}

	expected := map[string]string{
		"one":   "1.0",
		"two":   "2.0",
		"three": "3.0",
	}

	require.Equal(t, expected, deps.ToMap())
}
