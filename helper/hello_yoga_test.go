package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// run unit test in go : go test

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before unit test")

	m.Run()

	// after
	fmt.Println("After unit test")
}

func TestHelloYogaAssert(t *testing.T) {
	result := HelloYoga("Yoga Saputra")
	assert.Equal(t, "Hello Yoga Saputra", result, "Result must be 'Hi Yoga Saputra'")
	fmt.Println("Test TestHelloYoga with assert done")
}

func TestHelloYogaRequire(t *testing.T) {
	result := HelloYoga("Yoga Saputra")
	require.Equal(t, "Hello Yoga Saputra", result, "Result must be 'Hi Yoga Saputra'")
	fmt.Println("Test TestHelloYoga with assert done")
}

// skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloYoga("Yoga Saputra")
	assert.Equal(t, "Hello Yoga Saputra", result, "Result must be 'Hi Yoga Saputra'")
}

// Test table
func TestTableHelloYoga(t *testing.T) {
	// struct slice
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Yoga",
			request:  "Yoga ",
			expected: "Hello Yoga",
		},
		{
			name:     "Saputra",
			request:  "Saputra",
			expected: "Hello Saputra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloYoga(test.name)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkTableHelloYoga(b *testing.B) {
	// struct slice
	benchs := []struct {
		name, request, expected string
	}{
		{
			name:     "Yoga",
			request:  "Yoga ",
			expected: "Hello Yoga",
		},
		{
			name:     "Saputra",
			request:  "Saputra",
			expected: "Hello Saputra",
		},
	}

	for _, bench := range benchs {
		b.Run(bench.name, func(b *testing.B) {
			result := HelloYoga(bench.name)
			require.Equal(b, bench.expected, result)
		})
	}
}
