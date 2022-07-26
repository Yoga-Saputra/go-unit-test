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
